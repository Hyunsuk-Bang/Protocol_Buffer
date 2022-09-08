package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtoToJson(p proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	out, err := option.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

func JsonToProto(json string, p proto.Message) {
	err := protojson.Unmarshal([]byte(json), p)
	if err != nil {
		log.Fatal(err)
	}
}
