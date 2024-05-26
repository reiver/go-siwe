package siwe

import (
	"sourcecode.social/reiver/go-opt"
)

func (receiver *Message) SetStatement(value string) {
	receiver.Statement = opt.Something(value)
}

func (receiver *Message) UnsetStatement(value string) {
	receiver.Statement = opt.Nothing[string]()
}
