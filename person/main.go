package main

import (
	"fmt"
	pb "person/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewPerson(name string, id int32, email string, phones []*pb.Person_PhoneNumber) *pb.Person {
	return &pb.Person{
		Name:        name,
		Id:          id,
		Email:       email,
		Phones:      phones,
		LastUpdated: timestamppb.Now(),
	}
}

func NewPhoneNumber(number string, pt pb.Person_PhoneType) *pb.Person_PhoneNumber {
	return &pb.Person_PhoneNumber{
		Number: number,
		Type:   pt,
	}
}

func addPhoneNumber(p *pb.Person, pn *pb.Person_PhoneNumber) {
	p.Phones = append(p.Phones, pn)
	p.LastUpdated = timestamppb.Now()
}

func addPersonToAddressBook(a *pb.AddressBook, p *pb.Person) {
	a.People = append(a.People, p)
}

func main() {
	addressBook := &pb.AddressBook{}

	p := NewPerson(
		"Hyunsuk Bang",
		2043,
		"hbang3@hawk.iit.edu",
		[]*pb.Person_PhoneNumber{
			NewPhoneNumber("312-900-5708", pb.Person_HOME)})

	addPhoneNumber(p, NewPhoneNumber("02-557-4991", pb.Person_MOBILE))
	addPhoneNumber(p, NewPhoneNumber("010-9015-4991", pb.Person_WORK))
	addPersonToAddressBook(addressBook, p)

	writeToFile("address_Book.bin", addressBook)
	addressBook = readFromFile("address_Book.bin")
	fmt.Println(addressBook)

	jsonBook := ProtoToJson(addressBook)
	fmt.Println(jsonBook)
	JsonToProto(jsonBook, addressBook)
	fmt.Println(addressBook)
}
