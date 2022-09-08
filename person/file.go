package main

import (
	"io/ioutil"
	"log"
	pb "person/proto"

	"google.golang.org/protobuf/proto"
)

func writeToFile(fname string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatal(err)
		return
	}
}

func readFromFile(fname string) *pb.AddressBook {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	addressBook := &pb.AddressBook{}
	if err = proto.Unmarshal(in, addressBook); err != nil {
		log.Fatal(err)
	}
	return addressBook
}
