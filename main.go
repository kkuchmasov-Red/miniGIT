package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	comandName, _ := parserCommand()

	switch comandName {
	case "init":
		initialization()
	}

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
	mainRepository := ".my-git"

	os.MkdirAll(mainRepository, 0777)

	for _, repository := range []string{"objects", "refs", "HEAD"} {
		os.MkdirAll(mainRepository+"/"+repository, 0777)
	}

}
