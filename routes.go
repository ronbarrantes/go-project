package main

type UserInfo struct {
	firstName string
	lastName  string
}

// users
func CreateUser(userInfo UserInfo)            {}
func GetUsers()                               {}
func GetUser(id string)                       {}
func UpdateUser(id string, userInfo UserInfo) {}
func DeleteUser(id string)                    {}

// cutomer
// store
