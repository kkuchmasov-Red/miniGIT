package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
	"strings"
)

const (
	MAIN_REPO    = ".my-git"
	OBJECTS_REPO = "objects"
	REFS_REPO    = "refs"
	HEAD         = "HEAD"
	BLOB         = "BLOB"
)

type TreeEntry struct {
	Path string
	Hash string
}

type Tree []slaveTree

type object struct {
	hash    string
	data    string
	typeObj string
}

func main() {

	for {
		comandName, params := parserCommand()
		switch comandName {
		case "init":
			initialization()
		case "hash":
			saveObject(params)
		case "cat":
			readObject(params)
		case "test":
			fmt.Printf("%b\n", 0|1)
		}
	}
}

func readObject(args []string) {
	if len(args) != 1 {
		fmt.Println("Неверное количество параметров")
		return
	}
	hash := args[0]

	filePath := MAIN_REPO + "/" + OBJECTS_REPO + "/" + hash

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(file))
}

func saveObject(args []string) {
	if len(args) != 1 {
		fmt.Println("Неверное количество параметров")
		return
	}

	object, err := hash(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createObject(object)
	if err != nil {
		fmt.Println(err)
	}

}

func createObject(object object) error {
	fileName := MAIN_REPO + "/" + OBJECTS_REPO + "/" + object.hash

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(object.data))
	if err != nil {
		return err
	}
	fmt.Println(object.hash)
	return nil
}

func hash(nameFile string) (object, error) {
	var object object
	fullName := "test/" + nameFile
	data, err := os.ReadFile(fullName)
	if err != nil {
		return object, err
	}
	object.hash = fmt.Sprintf("%x", sha1.Sum(data))
	object.data = string(data)
	object.typeObj = BLOB
	return object, nil

}

func parserCommand() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)

	var comandName string
	var args []string
	if !scanner.Scan() {
		return comandName, args
	}

	line := scanner.Text()
	stringSplit := strings.Fields(line)
	if len(stringSplit) == 0 {
		return comandName, args
	}

	comandName = stringSplit[0]
	if len(stringSplit) > 1 {
		args = stringSplit[1:]
	}
	return comandName, args
}

func initialization() {

	os.MkdirAll(MAIN_REPO, 0777)

	for _, repository := range []string{OBJECTS_REPO, REFS_REPO} {
		os.MkdirAll(MAIN_REPO+"/"+repository, 0777)
	}

	os.Create(MAIN_REPO + "/" + HEAD)

}
