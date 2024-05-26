package siwe

import (
	"sourcecode.social/reiver/go-opt"
)

func (receiver *Message) SetURI(value string) {
	receiver.URI = opt.Something(value)
}

func (receiver *Message) UnsetURI(value string) {
	receiver.URI = opt.Nothing[string]()
}
