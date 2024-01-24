package main

import (
	"github.com/arctan90/silk2wav"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

func main() {
	if len(os.Args) != 3 {
		log.Printf("Usage: %s inputfile outputfile\n", os.Args[0])
		return
	}

	data, err := readFile(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}

	out, err := silk2wav.Convert(data)
	if err != nil {
		log.Println(err)
		return
	}

	if err := writeFile(os.Args[2], out, 0755); err != nil {
		log.Printf("main.go write file failed: %s\n", err.Error())
	}
	log.Printf("Converted %s to %s\n", os.Args[1], os.Args[2])
}

func readFile(name string) ([]byte, error) {
	version := runtime.Version()
	var data []byte
	var err error
	if version >= "go1.16" {
		data, err = os.ReadFile(name)
	} else {
		data, err = ioutil.ReadFile(name)
	}

	if err != nil {
		return nil, err
	}
	return data, nil
}

func writeFile(name string, data []byte, perm fs.FileMode) error {
	version := runtime.Version()
	var err error
	if version >= "go1.16" {
		err = os.WriteFile(name, data, perm)
	} else {
		err = ioutil.WriteFile(name, data, perm)
	}

	if err != nil {
		return err
	}
	return nil
}
