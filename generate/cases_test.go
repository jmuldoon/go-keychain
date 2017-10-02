package generate

import (
	"fmt"
)

var errDefault = fmt.Errorf("generate: expected error")

type Expected struct {
	Value interface{}
	Error error
}

type Tested struct {
	Value interface{}
	Error error
}

var testStringLengthCases = []struct {
	Description string
	Expected
	Tested
}{
	{
		Description: `Int length of the password 2`,
		Expected:    Expected{Value: 2},
		Tested:      Tested{Value: 2},
	},
	{
		Description: `Int length of the password 4`,
		Expected:    Expected{Value: 4},
		Tested:      Tested{Value: 4},
	},
	{
		Description: `Int length of the password 8`,
		Expected:    Expected{Value: 8},
		Tested:      Tested{Value: 8},
	},
	{
		Description: `Int length of the password 16`,
		Expected:    Expected{Value: 16},
		Tested:      Tested{Value: 16},
	},
	{
		Description: `Int length of the password 32`,
		Expected:    Expected{Value: 32},
		Tested:      Tested{Value: 32},
	},
	{
		Description: `Int length of the password 64`,
		Expected:    Expected{Value: 64},
		Tested:      Tested{Value: 64},
	},
}
