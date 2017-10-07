package arguments

import (
	"flag"
	"fmt"
)

const testVersion = 0

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
	if *a.Account == "" || *a.Service == "" || *a.Label == "" || (*a.Data == "" && !*a.Read && *a.Generate == 0) {
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
	a.Generate = flag.Int("generate", 0, "Generate a random string length N specified")
	a.Data = flag.String("data", "", "Data to encrypt/store")
	a.Label = flag.String("label", "", "Keychain label/name. (Required)")
	a.Service = flag.String("service", "", "Service/where the key is stored. (Required)")
	a.Read = flag.Bool("read", false, "Read the label, data keypair")
	flag.Parse()

	return a.validateRequiredArguments()
}
