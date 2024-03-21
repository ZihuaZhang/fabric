package main

import (
	"fmt"
	"github.com/ZihuaZhang/fabric-protos-go/common"
	"github.com/golang/protobuf/proto"
)

func main() {
	ce := common.Envelope{Payload: []byte("payload"), RedactMessage: []byte("redact"), Signature: []byte("signature"), Redactable: false}
	fmt.Println(ce)
	pce, _ := proto.Marshal(&ce)

	ppce := common.Envelope{}
	proto.Unmarshal(pce, &ppce)
	fmt.Println(ppce)
}
