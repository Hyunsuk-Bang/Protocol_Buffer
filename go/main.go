package main

import (
	"fmt"
	pb "protobuf/proto"
	"reflect"

	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:         42,
		IsSimple:   true,
		Name:       "A Name",
		SampleList: []int32{1, 2, 3, 4, 5, 123, 5},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{
			Id:   42,
			Name: "My name",
		},
		MultipleDummy: []*pb.Dummy{
			{Id: 43, Name: "My name 2"},
			{Id: 44, Name: "My name 2"},
			{Id: 435, Name: "My name 3"},
			{Id: 47, Name: "My name 54"},
			{Id: 48, Name: "My name 4"},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_BLUE,
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Printf("message has unexpected type %v", x)
	}

}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myid":  {Id: 42},
			"myid1": {Id: 43},
			"myid2": {Id: 44},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"
	writeToFile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)
	fmt.Println(message)
}

func doToJson(p proto.Message) string {
	out := toJson(p)
	return out
}

func doFromJson(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJson(jsonString, message)
	return message
}

func main() {
	//fmt.Println(doComplex())
	//fmt.Println(doEnum())
	// (doOneOf(&pb.Result_Id{
	// 	Id: 32,
	// }))

	// (doOneOf(&pb.Result_Message{
	// 	Message: "a message",
	// }))

	// (doOneOf(
	// 	"string",
	// ))
	// fmt.Println(doMap())
	jsonString := doToJson(doSimple())
	message := doFromJson(jsonString, reflect.TypeOf(pb.Simple{}))
	fmt.Println(jsonString)
	fmt.Println(message)
}
