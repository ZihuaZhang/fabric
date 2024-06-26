/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package policies

import (
	"math"

	cb "github.com/ZihuaZhang/fabric-protos-go/common"
	"github.com/ZihuaZhang/fabric/common/policydsl"
	"github.com/ZihuaZhang/fabric/protoutil"
	mspa "github.com/hyperledger/fabric-protos-go/msp"
)

const (
	BlockValidationPolicyKey = "BlockValidation"
)

func EncodeBFTBlockVerificationPolicy(consenterProtos []*cb.Consenter, ordererGroup *cb.ConfigGroup) {
	n := len(consenterProtos)
	f := (n - 1) / 3

	var identities []*mspa.MSPPrincipal
	var pols []*cb.SignaturePolicy
	for i, consenter := range consenterProtos {
		pols = append(pols, &cb.SignaturePolicy{
			Type: &cb.SignaturePolicy_SignedBy{
				SignedBy: int32(i),
			},
		})
		identities = append(identities, &mspa.MSPPrincipal{
			PrincipalClassification: mspa.MSPPrincipal_IDENTITY,
			Principal:               protoutil.MarshalOrPanic(&mspa.SerializedIdentity{Mspid: consenter.MspId, IdBytes: consenter.Identity}),
		})
	}

	quorumSize := ComputeBFTQuorum(n, f)
	sp := &cb.SignaturePolicyEnvelope{
		Rule:       policydsl.NOutOf(int32(quorumSize), pols),
		Identities: identities,
	}
	ordererGroup.Policies[BlockValidationPolicyKey] = &cb.ConfigPolicy{
		// Inherit modification policy
		ModPolicy: ordererGroup.Policies[BlockValidationPolicyKey].ModPolicy,
		Policy: &cb.Policy{
			Type:  int32(cb.Policy_SIGNATURE),
			Value: protoutil.MarshalOrPanic(sp),
		},
	}
}

func ComputeBFTQuorum(totalNodes, faultyNodes int) int {
	return int(math.Ceil(float64(totalNodes+faultyNodes+1) / 2))
}
