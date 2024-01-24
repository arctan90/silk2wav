package main

import (
	"github.com/arctan90/silk2wav"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Printf("Usage: %s inputfile outputfile\n", os.Args[0])
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}

	out, err := silk2wav.Convert(data)
	if err != nil {
		log.Println(err)
		return
	}

	if err := os.WriteFile(os.Args[2], out, 0755); err != nil {
		log.Printf("main.go write file failed: %s\n", err.Error())
	}
	log.Printf("Converted %s to %s\n", os.Args[1], os.Args[2])
}
