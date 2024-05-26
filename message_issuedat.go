package siwe

import (
	"time"

	"sourcecode.social/reiver/go-opt"
)

func (receiver *Message) SetIssuedAt(value string) {
	receiver.IssuedAt = opt.Something(value)
}

func (receiver *Message) SetIssuedAtToNow() {
	receiver.SetIssuedAt( time.Now().UTC().Format(time.RFC3339) )
}

func (receiver *Message) UnsetIssuedAt(value string) {
	receiver.IssuedAt = opt.Nothing[string]()
}
