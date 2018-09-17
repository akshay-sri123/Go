package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Starting the protobuf test.")
	akshay := &Person{
		Name: "Akshay",
		Age:  23,
	}
	data, err := proto.Marshal(akshay)
	if err != nil {
		log.Fatal("Error during Marshalling", err)
	}
	fmt.Println(data)

	personObject := &Person{}
	err = proto.Unmarshal(data, personObject)
	if err != nil {
		log.Fatal("Error during unmarshalling", err)
	}
	fmt.Println(personObject.GetName())
	fmt.Println(personObject.GetAge())
}
