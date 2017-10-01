# go-keychain

Personal use, password export for applications. Quick and dirty implemenation for password setting and retrieval. Plus quick use case.

## example cli arguments

Arguments   | Values        | Description
---         | ---           | ---
account     | string        | user account (required)
group       | string        | user access group
data        | string        | password/encrypted part (required if writing)
generate    | integer       | integer length of the password string to be generated
label       | string        | name (required)
service     | string        | where (required)
read        | bool          | defaults false. If you want to read must be specified

### Example useage

``` bash
go-keychain -account username -service testservice -label testlabel -data tests123
go-keychain -account username -service testservicepath2 -label testlabel -generate 32
```

## .bashrc / .bash_profile

Please add the following line to your bash profile so that when loaded you can automatically export
the default base password you'd like.

``` bash
export DEFAULT_KEY_NAME=$(go-keychain -account username -service testservice -label testlabel -read)
```
