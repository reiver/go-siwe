package siwe

import (
	"sourcecode.social/reiver/go-opt"
)

func (receiver *Message) SetNonce(value string) {
	receiver.Nonce = opt.Something(value)
}

func (receiver *Message) UnsetNonce(value string) {
	receiver.Nonce = opt.Nothing[string]()
}
