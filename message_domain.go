package siwe

import (
	"sourcecode.social/reiver/go-opt"
)

func (receiver *Message) SetDomain(value string) {
	receiver.Domain = opt.Something(value)
}

func (receiver *Message) UnsetDomain(value string) {
	receiver.Domain = opt.Nothing[string]()
}
