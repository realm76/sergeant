package parser

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/realm76/sergeant/pkg/makefile"
	"os"
	"strings"
)

func ParseFile(path string) (*makefile.Makefile, error) {
	makeFile := makefile.Makefile{}

	makeFileHandle, err := os.Open(path + "\\Makefile")
	if err != nil {
		return nil, err
	}

	defer func(makeFile *os.File) {
		_err := makeFile.Close()
		if _err != nil {
			err = errors.Join(err, _err)
		}
	}(makeFileHandle)

	scanner := bufio.NewScanner(makeFileHandle)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			variable := ParseLineForVariable(line)
			if variable != nil {
				fmt.Printf("Variable: %v\n", variable)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &makeFile, nil
}

// ParseLineForVariable /*
/**
immediate = deferred
immediate ?= deferred
immediate := immediate
immediate ::= immediate
immediate :::= immediate-with-escape
immediate += deferred or immediate
immediate != immediate
*/
func ParseLineForVariable(line string) *makefile.Variable {
	lineCharacters := strings.Split(line, "")
	variable := makefile.Variable{}
	x := 0
	y := 0

	for i := 0; i < len(lineCharacters); i++ {
		char := lineCharacters[i]
		switch char {
		case "=":
		case "?":
		case ":":
		case "+":
		case "!":
			if x == 0 {
				x = i
			}

			if y == 0 {
				y = i
			}
		default:
			if y > 0 {
				break
			}
		}
	}

	if x == 0 || y == 0 {
		return nil
	}

	variable.Name = strings.Join(lineCharacters[:x], "")
	variable.Value = strings.Join(lineCharacters[y:], "")

	return &variable
}
