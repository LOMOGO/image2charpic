package main

import (
	"Image2CharPic/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Exexute err: %v", err)
	}
}
