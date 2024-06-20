package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"silk2wav/decoder"
)

func main() {
	//fmt.Printf("Usage: %s inputfile-directory outputfile-directory\n", os.Args[0])
	folderOrg := ""
	folderDst := ""

	if len(os.Args) > 1 {
		folderOrg = os.Args[1]
	}

	if len(os.Args) > 2 {
		folderDst = os.Args[2]
	}

	if folderDst == "" {
		folderDst = "./out_wavs"
		_, err := os.Stat(folderDst)
		if err != nil && os.IsNotExist(err) {
			er := os.MkdirAll(folderDst, 0755)
			if er != nil {
				fmt.Printf("Create dest folder %s failed for %s", folderDst, er.Error())
				return
			}
		}
	}

	if !isDir(folderOrg) && !isDir(folderDst) {
		trans2wav_vf(folderOrg, folderDst)
		return
	}

	if !isDir(folderOrg) || !isDir(folderDst) {
		fmt.Printf("Usage: %s inputfile-directory(or filePath)  outputfile-directory(or filePath)\n", os.Args[0])
		return
	}

	dirEntries, err := os.ReadDir(folderOrg)
	if err != nil {
		fmt.Printf("os.ReadDir() 返回错误: %v\n", err)
		return
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}

		trans2wav(entry.Name(), folderOrg, folderDst)
	}

}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func trans2wav(fileName, folderOrg, folderDst string) error {
	data, err := ioutil.ReadFile(folderOrg + "/" + fileName)
	if err != nil {
		log.Println(err)
		return err
	}

	out, err := decoder.Convert(data)
	if err != nil {
		log.Println(err)
		return err
	}

	targetName := folderDst + "/" + fileName + ".wav"

	ioutil.WriteFile(targetName, out, 0755)
	log.Printf("Converted %s to %s\n", folderOrg+"/"+fileName, targetName)

	return nil
}

func trans2wav_vf(inFileName, outFileName string) error {
	data, err := ioutil.ReadFile(inFileName)
	if err != nil {
		log.Println(err)
		return err
	}

	out, err := decoder.Convert(data)
	if err != nil {
		log.Println(err)
		return err
	}

	ioutil.WriteFile(outFileName, out, 0755)
	log.Printf("Converted %s to %s\n", inFileName, outFileName)

	return nil
}
