package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("hello")

	ryan := &Person{
		Name:  "Ryan",
		Id:    33,
		Email: "yliang@splunk.com",
		Phones: []*Person_PhoneNumber{
			{Number: "3232-11243",
				Type: 0},
			{Number: "11-000243",
				Type: 1},
		},
	}

	ryanNumber := &Person_PhoneNumber{
		Number: "3232-11243",
		Type:   0,
	}

	data, _ := proto.Marshal(ryan)
	fmt.Println(data)

	//outRyan := &Person{}
	//_=proto.Unmarshal(data, outRyan)

	//fmt.Println(outRyan)

	dataNumber, _ := proto.Marshal(ryanNumber)
	fmt.Println(dataNumber)

	outRyanNumber := &Person_PhoneNumber{}
	_ = proto.Unmarshal(dataNumber, outRyanNumber)
	fmt.Println(outRyanNumber.GetNumber())
	fmt.Println(outRyanNumber.GetType())

}
