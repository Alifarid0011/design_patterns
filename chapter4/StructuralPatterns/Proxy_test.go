package main

import (
	"math/rand"
	"testing"
)
func Test_UserListProxy(t *testing.T) {
	someDatabase := UserList{}
	rand.Seed(2342342)
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, User{ID: n})
	}
