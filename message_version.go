package siwe

import (
	"sourcecode.social/reiver/go-opt"
)

func (receiver *Message) SetVersion(value string) {
	receiver.Version = opt.Something(value)
}

func (receiver *Message) SetVersionTo1() {
	receiver.SetVersion("1")
}

func (receiver *Message) UnsetVersion(value string) {
	receiver.Version = opt.Nothing[string]()
}
