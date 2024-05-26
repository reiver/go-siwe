package siwe

import (
	"sourcecode.social/reiver/go-opt"
)

func (receiver *Message) SetAddress(value string) {
	receiver.Address = opt.Something(value)
}

func (receiver *Message) UnsetAddress(value string) {
	receiver.Address = opt.Nothing[string]()
}
