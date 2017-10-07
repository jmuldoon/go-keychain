package arguments

import (
	"flag"
	"fmt"
)

const testVersion = 0

var errRequiredArgumentsNotSet = fmt.Errorf("All required arguments were not set")

// Parser is the interface by which the flags are controlled.
type Parser interface {
	// TODO: fully controlled version
	// setFlagArguments(*flag.FlagSet)
	setFlagArguments()
	validateRequiredArguments() error
}

// Args are arguments that are passed in over the command line
type Args struct {
	Account  *string
	Group    *string
	Data     *string
	Generate *int
	Label    *string
	Service  *string
	Read     *bool
}

func (a *Args) validateRequiredArguments() error {
	if *a.Account == "" || *a.Service == "" || *a.Label == "" ||
		(*a.Data == "" && !*a.Read && *a.Generate == 0) {
		flag.PrintDefaults()
		return errRequiredArgumentsNotSet
	}
	return nil
}

// func (a *Args) setFlagArguments(cli *flag.FlagSet) {
// 	a.Account = cli.String("account", "", "User Account. (Required)")
// 	a.Group = cli.String("group", "", "User Group.")
// 	a.Generate = cli.Int("generate", 0, "Generate a random string length N specified")
// 	a.Data = cli.String("data", "", "Data to encrypt/store")
// 	a.Label = cli.String("label", "", "Keychain label/name. (Required)")
// 	a.Service = cli.String("service", "", "Service/where the key is stored. (Required)")
// 	a.Read = cli.Bool("read", false, "Read the label, data keypair")
// }
func (a *Args) setFlagArguments() {
	a.Account = flag.String("account", "", "User Account. (Required)")
	a.Group = flag.String("group", "", "User Group.")
	a.Generate = flag.Int("generate", 0, "Generate a random string length N specified")
	a.Data = flag.String("data", "", "Data to encrypt/store")
	a.Label = flag.String("label", "", "Keychain label/name. (Required)")
	a.Service = flag.String("service", "", "Service/where the key is stored. (Required)")
	a.Read = flag.Bool("read", false, "Read the label, data keypair")
}

// Parse will take the user's commandline args and parse them out into a usable
// object
// func Parse(p Parser, cli *flag.FlagSet) error {
func Parse(p Parser) error {
	// p.setFlagArguments(cli)
	p.setFlagArguments()
	flag.Parse()

	return p.validateRequiredArguments()
}
