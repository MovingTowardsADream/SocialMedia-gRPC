package entity

import "fmt"

type User struct {
	Id       int
	Email    string `protobuf:"bytes,1,opt,name=email"`
	Username string `protobuf:"bytes,2,opt,name=username"`
	Password string `protobuf:"bytes,3,opt,name=password"`
}

func (i *User) Reset()         { *i = User{} }
func (i *User) String() string { return fmt.Sprintf("%+v", *i) }
func (i *User) ProtoMessage()  {}
