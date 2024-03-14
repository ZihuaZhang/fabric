/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"os"

	"github.com/ZihuaZhang/fabric-chaincode-go/shim"
	"github.com/ZihuaZhang/fabric/integration/chaincode/kvexecutor"
)

func main() {
	err := shim.Start(&kvexecutor.KVExecutor{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exiting Simple chaincode: %s", err)
		os.Exit(2)
	}
}
