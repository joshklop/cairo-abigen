package main

import (
	"path/filepath"
	"fmt"
	"os"
	"github.com/joshklop/cairo-abigen"
	"github.com/joshklop/cairo-abigen/utils"
)

func main() {
	pkg := os.Args[1]
	abiFilename, err := filepath.Abs(os.Args[2])
	if err != nil {
		panic("-1"+err.Error())
	}
	fmt.Printf("pkg: %s\n", pkg)
	fmt.Printf("filename: %s\n", string(abiFilename))
	abiBytes, err := os.ReadFile(abiFilename)
	if err != nil {
		panic("0"+err.Error())
	}

	theABI, err := utils.ABIFromJSON(abiBytes)
	if err != nil {
		panic("1"+err.Error())
	}
	out, err := abi.Generate(pkg, theABI)
	if err != nil {
		panic("2"+err.Error())
	}

	if err := os.WriteFile(fmt.Sprintf("%s.go", pkg), []byte(out), 0600); err != nil {
		panic("3"+err.Error())
	}
}
