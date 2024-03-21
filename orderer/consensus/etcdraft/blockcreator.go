/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package etcdraft

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	cb "github.com/ZihuaZhang/fabric-protos-go/common"
	"github.com/ZihuaZhang/fabric/common/flogging"
	"github.com/ZihuaZhang/fabric/protoutil"
	"github.com/fentec-project/gofe/abe"
	"github.com/golang/protobuf/proto"
)

// blockCreator holds number and hash of latest block
// so that next block will be created based on it.
type blockCreator struct {
	hash   []byte
	number uint64

	logger *flogging.FabricLogger
}

func (bc *blockCreator) createNextBlock(envs []*cb.Envelope) *cb.Block {
	data := &cb.BlockData{
		Data: make([][]byte, len(envs)),
	}
	bc.logger.Info(envs[0])
	var dataHash []byte
	for i, env := range envs {
		var err0 error
		data.Data[i], err0 = proto.Marshal(env)
		bc.logger.Info(data.Data[i])
		if err0 != nil {
			return nil
		}
		if env.Redactable {
			redactMsg := cb.RedactMsg{}
			err := proto.Unmarshal(env.RedactMessage, &redactMsg)
			if err != nil {
				return nil
			}
			pk := protoutil.FAMEPubKey{}
			err1 := json.Unmarshal(redactMsg.Pk, &pk)
			if err1 != nil {
				return nil
			}
			msp := abe.MSP{}
			err2 := json.Unmarshal(redactMsg.Msp, &msp)
			if err2 != nil {
				return nil
			}
			fameCipher, err3 := protoutil.NewFAME().Hash(&msp, &pk)
			if err3 != nil {
				return nil
			}
			dataHash = append(dataHash, fameCipher.Hash.Marshal()...)
			fameCipherBytes, error4 := json.Marshal(fameCipher)
			if error4 != nil {
				return nil
			}
			redactMsg.FameCipher = fameCipherBytes
			env.RedactMessage, _ = proto.Marshal(&redactMsg)
		} else {
			hash := sha256.Sum256(data.Data[i])
			dataHash = append(dataHash, hash[:]...)
		}
	}
	//var err error
	//for i, env := range envs {
	//	data.Data[i], err = proto.Marshal(env)
	//	if err != nil {
	//		bc.logger.Panicf("Could not marshal envelope: %s", err)
	//	}
	//}

	bc.number++
	fmt.Println(dataHash)

	block := protoutil.NewBlock(bc.number, bc.hash)
	// block.Header.DataHash = protoutil.ComputeBlockDataHash(data)
	block.Header.DataHash = dataHash
	block.Data = data

	bc.hash = protoutil.BlockHeaderHash(block.Header)
	return block
}
