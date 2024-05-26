package siwe

import (
	"sourcecode.social/reiver/go-opt"
)

func (receiver *Message) SetRequestID(value string) {
	receiver.RequestID = opt.Something(value)
}

func (receiver *Message) UnsetRequestID(value string) {
	receiver.RequestID = opt.Nothing[string]()
}
