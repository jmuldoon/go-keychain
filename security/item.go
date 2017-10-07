package security

import (
	"fmt"

	keychain "github.com/keybase/go-keychain"
)

const testVersion = 0

var errDuplicate = fmt.Errorf("Write: duplicate item")
var errQuery = fmt.Errorf("Query: generic query failure")
var errQueryMultipleReturn = fmt.Errorf("Query: multiple entries returned")

// Worker is the interface that actually does the external calls
type Worker interface {
	addItem() error
	queryItem() ([]keychain.QueryResult, error)
	setKeychainItem(keychain.Item)
}

// Item is the keychain object to be used for read/write as well as the implementor
// of the Worker interface
type Item struct {
	Account string
	Group   string
	Data    []byte
	Label   string
	Service string
}

// baseKeychainSetup returns the keychain item after initializing the params.
func (i *Item) baseKeychainSetup() keychain.Item {
	nitem := keychain.NewItem()
	nitem.SetSecClass(keychain.SecClassGenericPassword)
	nitem.SetService(i.Service)
	nitem.SetAccount(i.Account)
	nitem.SetLabel(i.Label)
	nitem.SetAccessGroup(i.Group)
	return nitem
}

// ExternalKeychain is the structure that will implement the Worker interface so
// that the external calls to keychain are abstracted appropriately.
type ExternalKeychain struct {
	keychainItem keychain.Item
}

func (i *ExternalKeychain) addItem() error {
	return keychain.AddItem(i.keychainItem)
}

func (i *ExternalKeychain) queryItem() ([]keychain.QueryResult, error) {
	return keychain.QueryItem(i.keychainItem)
}

func (i *ExternalKeychain) setKeychainItem(kci keychain.Item) {
	i.keychainItem = kci
}

// Write writes the new item to the OSX Keychain utility.
func (i *Item) Write(w Worker) error {
	nitem := i.baseKeychainSetup()
	nitem.SetData(i.Data)
	nitem.SetSynchronizable(keychain.SynchronizableNo)
	nitem.SetAccessible(keychain.AccessibleWhenUnlocked)
	w.setKeychainItem(nitem)
	err := w.addItem()

	if err == keychain.ErrorDuplicateItem {
		err = errDuplicate
	}
	return err
}

// Read data attempts to find the single entry for a specified password.
func (i *Item) Read(w Worker) (string, error) {
	query := i.baseKeychainSetup()
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	w.setKeychainItem(query)
	results, err := w.queryItem()

	password := ""
	if len(results) > 1 {
		err = errQueryMultipleReturn
	} else if err != nil {
		err = errQuery
	} else {
		password = string(results[0].Data)
	}
	return password, err
}
