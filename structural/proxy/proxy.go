package proxy

import "errors"

// UserFinder is the interface that the database and the Proxy will implement
type UserFinder interface {
	FindUser(id int32) (User, error)
}

// User represent a user into database or Proxy
type User struct {
	ID int32
}

// UserList holds a slice of users
type UserList []User

// UserListProxy is our proxy type
type UserListProxy struct {
	SomeDatabase              UserList
	StackCache                UserList
	StackCapacity             int
	DidDidLastSearchUserCache bool
}

// FindUser finds user from proxy or database
func (u *UserListProxy) FindUser(id int32) (User, error) {
	return User{}, errors.New("Not implemented yet")
}
