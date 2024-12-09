package main

import (
	"fmt"
	"reflect"
)

/* assignment code
type contact struct {
	sendingLimit int32
	userID       string
	age          int32
}

type perms struct {
	canSend         bool
	canReceive      bool
	permissionLevel int
	canManage       bool
}
*/

type contact struct {
	sendingLimit int32
	age          int32
	userID       string
}

type perms struct {
	canSend         bool
	canReceive      bool
	canManage       bool
	permissionLevel int
}

func main() {
	typ := reflect.TypeOf(contact{})
	fmt.Printf("Struct is %d bytes\n", typ.Size())
}
