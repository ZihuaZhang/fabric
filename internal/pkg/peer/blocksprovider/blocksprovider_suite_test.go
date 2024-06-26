/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package blocksprovider_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ZihuaZhang/fabric-protos-go/orderer"
	"github.com/ZihuaZhang/fabric/internal/pkg/identity"
)

//go:generate counterfeiter -o fake/signer.go --fake-name Signer . signer
type signer interface {
	identity.SignerSerializer
}

//go:generate counterfeiter -o fake/ab_deliver_client.go --fake-name DeliverClient . abDeliverClient
type abDeliverClient interface {
	orderer.AtomicBroadcast_DeliverClient
}

func TestBlocksProvider(t *testing.T) {
	RegisterFailHandler(Fail)

	suiteConf, reporterConf := GinkgoConfiguration()
	suiteConf.EmitSpecProgress = true
	reporterConf.FullTrace = true

	RunSpecs(t, "Blocksprovider Suite", suiteConf, reporterConf)
}
