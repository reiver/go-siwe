package siwe

import (
	"sourcecode.social/reiver/go-opt"
)

func (receiver *Message) SetExpirationTime(value string) {
	receiver.ExpirationTime = opt.Something(value)
}

func (receiver *Message) UnsetExpirationTime(value string) {
	receiver.ExpirationTime = opt.Nothing[string]()
}
