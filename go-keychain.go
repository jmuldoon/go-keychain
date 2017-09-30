package main

import (
	"fmt"
	"os"

	"github.com/jmuldoon/go-keychain/arguments"
	"github.com/jmuldoon/go-keychain/kchain"
)

func main() {
	args := &arguments.Args{}
	if err := args.Parser(); err != nil {
		os.Exit(1)
	}

	kcitem := &kchain.Item{
		Account: *args.Account,
		Group:   *args.Group,
		Data:    []byte(*args.Data),
		Label:   *args.Label,
		Service: *args.Service,
	}

	if *args.Read {
		plaintextPassword, err := kcitem.Read()
		if err != nil {
			os.Exit(3)
		}
		fmt.Println(plaintextPassword)
	} else {
		if err := kcitem.Write(); err != nil {
			os.Exit(2)
		}
	}
}
