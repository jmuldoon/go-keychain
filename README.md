# go-keychain

Personal use, password export for applications. Quick and dirty implemenation for password setting and retrieval. Plus quick use case.

## example cli arguments

Arguments | Values | Description
---         | ---           | ---
account     | "username"    | user account (required)
group       | "usergroup"   | user access group
data        | "tests123"    | password/encrypted part (required if writing)
label       | "testlabel"   | name (required)
service     | "testservice" | where (required)
export      | false         | if you want to export specify this

### Example useage

``` bash
go-keychain -account username -service testservice -label testlabel -data tests123
```

## .bashrc / .bash_profile

Please add the following line to your bash profile so that when loaded you can automatically export
the default base password you'd like.

``` bash
export DEFAULT_KEY_NAME=$(go-keychain -account username -service testservice -label testlabel -read)
```
