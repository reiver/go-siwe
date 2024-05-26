package siwe

import (
	"sourcecode.social/reiver/go-opt"
)

func (receiver *Message) SetScheme(value string) {
	receiver.Scheme = opt.Something(value)
}

func (receiver *Message) UnsetScheme(value string) {
	receiver.Scheme = opt.Nothing[string]()
}
