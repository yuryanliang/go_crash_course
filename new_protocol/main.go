/*
run the program with
$ go run main.go person.pb.go
*/
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
		Socialfollowers: &SocialFollowers{
			Twitter: 333,
			Youtube: 3334,
		},
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
	fmt.Println(newElliot.Socialfollowers.GetTwitter())
	fmt.Println(newElliot.Socialfollowers.GetYoutube())

	//ryan example

	ryan := &Person{
		Age: 3,
		Socialfollowers: &SocialFollowers{
			Youtube: 333,
		},
	}

	newdata, err := proto.Marshal(ryan)
	fmt.Println(newdata)

	newRyan := &Person{}
	proto.Unmarshal(newdata, newRyan)
	fmt.Println("newRyan's Youtube", newRyan.Socialfollowers.GetYoutube())
	fmt.Println("newRyan's Age", newRyan.GetAge())
}
