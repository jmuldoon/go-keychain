package security

import (
	"fmt"

	keychain "github.com/keybase/go-keychain"
)

// Item is the keychain object to be used for read/write
type Item struct {
	Account string
	Group   string
	Data    []byte
	Label   string
	Service string
}

// Itemizer interface to access the needed keychain functionality
type Itemizer interface {
	Write()
	Read()
}

// Write writes the new item to the OSX Keychain utility.
func (i *Item) Write() error {
	nitem := keychain.NewItem()
	nitem.SetSecClass(keychain.SecClassGenericPassword)
	nitem.SetService(i.Service)
	nitem.SetAccount(i.Account)
	nitem.SetLabel(i.Label)
	nitem.SetAccessGroup(i.Group)
	nitem.SetData(i.Data)
	nitem.SetSynchronizable(keychain.SynchronizableNo)
	nitem.SetAccessible(keychain.AccessibleWhenUnlocked)
	err := keychain.AddItem(nitem)

	if err == keychain.ErrorDuplicateItem {
		err = fmt.Errorf("Write: duplicate error issue %s", err)
	}
	return err
}

// Read data attempts to find the single entry for a specified password.
func (i *Item) Read() (string, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(i.Service)
	query.SetAccount(i.Account)
	query.SetLabel(i.Label)
	query.SetAccessGroup(i.Group)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)

	password := ""
	if err != nil {
		err = fmt.Errorf("Query: query failed %s", err)
	} else if len(results) != 1 {
		err = fmt.Errorf("Query: multiple entries returned, %d; %s", len(results), err)
	} else {
		password = string(results[0].Data)
	}
	return password, err
}
