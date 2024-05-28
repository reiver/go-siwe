package siwe

import (
	"strconv"
)

func (receiver Message) EthereumSignableMessage() ([]byte, error) {
	var buffer [2048]byte
	return receiver.EthereumSignableMessageAppend(buffer[0:0])
}

func (receiver Message) EthereumSignableMessageAppend(p []byte) ([]byte, error) {

	var initiallength int = len(p)

	{
		var err error

		p, err = receiver.SerializeSIWEAppend(p)
		if nil != err {
			return p, err
		}
	}

	// \x19Ethereum Signed Message:\n<length of message>
	var prefixbuffer [256]byte
	var prefix []byte = prefixbuffer[0:0]
	{
		var length int = len(p) - initiallength

		prefix = append(prefix, "\x19Ethereum Signed Message:\n"...)
		prefix = append(prefix, strconv.FormatInt(int64(length), 10)...)
	}

	{
		p = append(p[:initiallength], append(prefix, p[initiallength:]...)...)
		return p, nil
	}
}
