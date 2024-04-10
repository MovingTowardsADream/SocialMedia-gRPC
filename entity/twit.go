package entity

import "fmt"

type Twit struct {
	Id   int64
	Twit string `protobuf:"bytes,1,opt,name=twit"`
}

func (i *Twit) Reset()         { *i = Twit{} }
func (i *Twit) String() string { return fmt.Sprintf("%+v", *i) }
func (i *Twit) ProtoMessage()  {}
