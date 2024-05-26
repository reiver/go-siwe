package siwe

import (
	"strconv"

	"sourcecode.social/reiver/go-opt"
)

func (receiver *Message) SetChainID(value string) {
	receiver.ChainID = opt.Something(value)
}

func (receiver *Message) SetChainIDUsingUint64(number uint64) {
	var value string = strconv.FormatUint(number, 10)
	receiver.SetChainID(value)
}

func (receiver *Message) UnsetChainID(value string) {
	receiver.ChainID = opt.Nothing[string]()
}
