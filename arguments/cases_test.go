package arguments

import "flag"

const (
	TestAccount  = "TestAccount"
	TestGroup    = "TestGroup"
	TestData     = "TestPasswordData"
	TestLabel    = "TestName"
	TestService  = "TestService"
	TestGenerate = 32
	TestRead     = true
	usage        = "TestUsageStatement"
)

type MockArgs struct {
	Account  *string
	Group    *string
	Data     *string
	Generate *int
	Label    *string
	Service  *string
	Read     *bool
}

type Expected struct {
	Value interface{}
	Error error
}

type Tested struct {
	Value Parser
	Error error
}

var testParse = []struct {
	Description string
	Expected
	Tested
}{
	{
		Description: `CLI Arguments defined, Successful`,
		Expected:    Expected{Error: nil},
		Tested: Tested{Value: &MockArgs{
			Account:  flag.String("account", TestAccount, usage),
			Group:    flag.String("group", TestGroup, usage),
			Generate: flag.Int("generate", TestGenerate, usage),
			Data:     flag.String("data", TestData, usage),
			Label:    flag.String("label", TestLabel, usage),
			Service:  flag.String("service", TestService, usage),
			Read:     flag.Bool("read", TestRead, usage),
		}},
	},
}

// var testValidateRequiredArguments = []struct {
// 	Description string
// 	Expected
// 	Tested
// }{
// 	{
// 		Description: `CLI Arguments defined, Required not specified`,
// 		Expected:    Expected{Error: errRequiredArgumentsNotSet},
// 		Tested:      Tested{Value: &MockArgs{}},
// 	},
// 	{
// 		Description: `CLI Arguments defined, Required all Specified`,
// 		Expected:    Expected{},
// 		Tested: Tested{Value: &MockArgs{
// 			Account:  flag.String("account", TestAccount, usage),
// 			Group:    flag.String("group", TestGroup, usage),
// 			Generate: flag.Int("generate", TestGenerate, usage),
// 			Data:     flag.String("data", TestData, usage),
// 			Label:    flag.String("label", TestLabel, usage),
// 			Service:  flag.String("service", TestService, usage),
// 			Read:     flag.Bool("read", TestRead, usage),
// 		}},
// 	},
// }
