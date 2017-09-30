package arguments

import (
	"flag"
	"fmt"
)

// Args are arguments that are passed in over the command line
type Args struct {
	Account *string
	Group   *string
	Data    *string
	Label   *string
	Service *string
	Read    *bool
}

func (a *Args) validateRequiredArguments() error {
	if *a.Account == "" || *a.Service == "" || *a.Label == "" || (*a.Data == "" && !*a.Read) {
		flag.PrintDefaults()
		return fmt.Errorf("All required arguments were not set")
	}
	return nil
}

// Parser will take the user's commandline args and parse them out into a usable
// object
func (a *Args) Parser() error {
	a.Account = flag.String("account", "", "User Account. (Required)")
	a.Group = flag.String("group", "", "User Group.")
	a.Data = flag.String("data", "", "Data to encrypt/store")
	a.Label = flag.String("label", "", "Keychain label/name. (Required)")
	a.Service = flag.String("service", "", "Service/where the key is stored. (Required)")
	a.Read = flag.Bool("read", false, "Read the label, data keypair")
	flag.Parse()

	// fmt.Printf("%s\n%s\n%s\n%s\n%s\n%t\n", *a.Account, *a.Group, *a.Data, *a.Label, *a.Service, *a.Read)

	return a.validateRequiredArguments()
}
