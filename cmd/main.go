package main

import (
	"fmt"
	"log"

	"L2/L2_9/pkg/unpacker"
)

func main() {
	fmt.Println("Unpacking started!")
	var packed string
	for {
		fmt.Print("Enter unpacked string: ")
		_, err := fmt.Scanf("%s", &packed)
		if err != nil {
			log.Fatal("Error reading input: ", err)
		}
		unpacked, err := unpacker.UnpackString(packed)
		if err != nil {
			log.Fatal("Error unpacking input: ", err)
		}
		fmt.Println("Unpacking finished: ", unpacked)
	}

}
