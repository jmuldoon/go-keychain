package security

import (
	"testing"

	keychain "github.com/keybase/go-keychain"
)

const targetTestVersion = 0

func TestVersionValidation(t *testing.T) {
	t.Parallel() // indicator that it can be tested in parallel
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v\n", testVersion, targetTestVersion)
	}
}

type mockExternalKeychain struct {
	writeCount   int
	readCount    int
	keychainItem keychain.Item
}

func (i *mockExternalKeychain) addItem() error {
	if i.writeCount > 0 {
		return keychain.ErrorDuplicateItem
	}
	i.writeCount++
	return nil
}

func (i *mockExternalKeychain) queryItem() ([]keychain.QueryResult, error) {
	queryResult := keychain.QueryResult{Data: []byte(TestData)}
	qResultList := []keychain.QueryResult{queryResult}
	if i.readCount == 1 {
		i.readCount++
		return nil, errQuery
	} else if i.readCount == 2 {
		i.readCount++
		return append(qResultList, qResultList...), nil
	} else {
		i.readCount++
		return qResultList, nil
	}
}

func (i *mockExternalKeychain) setKeychainItem(kci keychain.Item) {
	i.keychainItem = kci
}

func TestWrite(t *testing.T) {
	t.Parallel() // indicator that it can be tested in parallel
	mockExtKeychain := &mockExternalKeychain{}
	for _, test := range testKeychainWrite {
		observed := test.Tested.Value.Write(mockExtKeychain)
		t.Logf("Running test for `%s`\n", test.Description)
		if observed != test.Expected.Error {
			t.Fatalf("%s(%v):\n"+
				"Brief (%s)\n"+
				"Observed: %t\n"+
				"Expected: %t\n",
				"TestWrite", test.Tested.Value,
				test.Description, observed, test.Expected.Error)
		}
	}
}

func TestRead(t *testing.T) {
	t.Parallel() // indicator that it can be tested in parallel
	mockExtKeychain := &mockExternalKeychain{}
	for _, test := range testKeychainRead {
		observed, errObserved := test.Tested.Value.Read(mockExtKeychain)
		t.Logf("Running test for `%s`\n", test.Description)
		if errObserved != test.Expected.Error {
			t.Fatalf("%s(%v):\n"+
				"Brief (%s)\n"+
				"Observed: %t\n"+
				"Expected: %t\n",
				"TestRead", test.Tested.Value,
				test.Description, errObserved, test.Expected.Error)
		}
		if observed != test.Expected.Value {
			t.Fatalf("%s(%v):\n"+
				"Brief (%s)\n"+
				"Observed: %t\n"+
				"Expected: %t\n",
				"TestRead", test.Tested.Value,
				test.Description, observed, test.Expected.Value)
		}
	}
}
