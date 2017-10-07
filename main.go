package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jmuldoon/go-keychain/arguments"
	"github.com/jmuldoon/go-keychain/generate"
	"github.com/jmuldoon/go-keychain/security"
)

// Exit Codes
const (
	ExitGeneralErr = 1 + iota
	ExitArgParseErr
	ExitWriteErr
	ExitReadErr
)

func main() {
	args := &arguments.Args{}
	// commandLine := flag.NewFlagSet(os.Args[0], ExitOnError)
	// if err := arguments.Parse(args, commandLine); err != nil {
	// TODO: follow up with the adjustment to fully controlled system as deailed above.
	if err := arguments.Parse(args); err != nil {
		fmt.Println(err)
		os.Exit(ExitArgParseErr)
	}

	if *args.Generate > 0 {
		src := rand.NewSource(time.Now().UnixNano())
		*args.Data = generate.RandStringBytesMask(src, *args.Generate)
	}

	kcitem := &security.Item{
		Account: *args.Account,
		Group:   *args.Group,
		Data:    []byte(*args.Data),
		Label:   *args.Label,
		Service: *args.Service,
	}

	// create interface struct
	extKeychain := &security.ExternalKeychain{}
	if *args.Read {
		plaintextPassword, err := kcitem.Read(extKeychain)
		if err != nil {
			fmt.Println(err)
			os.Exit(ExitReadErr)
		}
		fmt.Println(plaintextPassword)
	} else {
		if err := kcitem.Write(extKeychain); err != nil {
			fmt.Println(err)
			os.Exit(ExitWriteErr)
		}
	}
}
