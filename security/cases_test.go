package security

const (
	TestAccount = "TestAccount"
	TestGroup   = "TestGroup"
	TestData    = "TestPasswordData"
	TestLabel   = "TestName"
	TestService = "TestService"
)

var TestItem = &Item{
	Account: TestAccount,
	Group:   TestGroup,
	Data:    []byte(TestData),
	Label:   TestLabel,
	Service: TestService,
}

type Expected struct {
	Value string
	Error error
}

type Tested struct {
	Value *Item
	Error error
}

var testKeychainWrite = []struct {
	Description string
	Expected
	Tested
}{
	{
		Description: `Keychain Item to be written, Successful`,
		Expected:    Expected{Error: nil},
		Tested:      Tested{Value: TestItem},
	},
	{
		Description: `Keychain Item to be written, Duplicate`,
		Expected:    Expected{Error: errDuplicate},
		Tested:      Tested{Value: TestItem},
	},
}

var testKeychainRead = []struct {
	Description string
	Expected
	Tested
}{
	{
		Description: `Keychain Item to be read, Successful`,
		Expected:    Expected{Value: string(TestItem.Data), Error: nil},
		Tested:      Tested{Value: TestItem},
	},
	{
		Description: `Keychain Item to be read, Generic Failure`,
		Expected:    Expected{Error: errQuery},
		Tested:      Tested{Value: TestItem},
	},
	{
		Description: `Keychain Item to be read, Multiple Return Failure`,
		Expected:    Expected{Error: errQueryMultipleReturn},
		Tested:      Tested{Value: TestItem},
	},
}
