package main

import "errors"

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}
type UserList []User

type UserListProxy struct {
	SomeDatabase           *UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	return User{}, errors.New("Not implemented yet")
}
func main() {

}
