package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"silk2wav/decoder"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s inputfile outputfile\n", os.Args[0])
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}

	out, err := decoder.Convert(data)
	if err != nil {
		log.Println(err)
		return
	}

	ioutil.WriteFile(os.Args[2], out, 0755)
	log.Printf("Converted %s to %s\n", os.Args[1], os.Args[2])
}
