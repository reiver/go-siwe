package siwe

import (
	"sourcecode.social/reiver/go-opt"
)

func (receiver *Message) SetNotBefore(value string) {
	receiver.NotBefore = opt.Something(value)
}

func (receiver *Message) UnsetNotBefore(value string) {
	receiver.NotBefore = opt.Nothing[string]()
}
