package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	fmt.Println("hello")

	elliot := &Person{
		Name: "Elliot",
		Age:  24,
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("Marshalling error", err)
	}
	fmt.Println(data)

	//unmarshall

	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("Unmarshalling error", err)
	}
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())

}
