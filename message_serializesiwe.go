package siwe

import (
	"strconv"

	"github.com/reiver/go-ascii"
)

func (receiver Message) SerializeSIWE() ([]byte, error) {
	var buffer [2048]byte
	return receiver.SerializeSIWEAppend(buffer[0:0])
}

func (receiver Message) SerializeSIWEAppend(p []byte) ([]byte, error) {

	var initiallength int = len(p)

	// [ scheme "://" ] domain %s" wants you to sign in with your Ethereum account:"
	{
		receiver.Scheme.WhenSomething(func(scheme string){
			p = append(p, scheme...)
			p = append(p, "://"...)
		})

		receiver.Domain.WhenSomething(func(domain string){
			p = append(p, domain...)
		})

		p = append(p, " wants you to sign in with your Ethereum account:"...)
	}

	// LF address
	{
		p = append(p, ascii.LF)

		receiver.Address.WhenSomething(func(address string){
			p = append(p, address...)
		})
	}

	// LF
	// LF
	{
		p = append(p, ascii.LF)
		p = append(p, ascii.LF)
	}

	// [ statement LF ]
	{
		receiver.Statement.WhenSomething(func(statement string){
			p = append(p, statement...)
			p = append(p, ascii.LF)
		})
	}

	// LF %s"URI: " uri
	p = appendfield(p, "URI: ", receiver.URI)

	// LF %s"Version: " version
	p = appendfield(p, "Version: ", receiver.Version)

	// LF %s"Chain ID: " chain-id
	p = appendfield(p, "Chain ID: ", receiver.ChainID)

	// LF %s"Nonce: " nonce
	p = appendfield(p, "Nonce: ", receiver.Nonce)

	// LF %s"Issued At: " issued-at
	p = appendfield(p, "Issued At: ", receiver.IssuedAt)

	// [ LF %s"Expiration Time: " expiration-time ]
	receiver.ExpirationTime.WhenSomething(func(string){
		p = appendfield(p, "Expiration Time: ", receiver.ExpirationTime)
	})

	// [ LF %s"Not Before: " not-before ]
	receiver.NotBefore.WhenSomething(func(string){
		p = appendfield(p, "Not Before: ", receiver.NotBefore)
	})

	// [ LF %s"Request ID: " request-id ]
	receiver.RequestID.WhenSomething(func(string){
		p = appendfield(p, "Request ID: ", receiver.RequestID)
	})

	// [ LF %s"Resources:" LF resources ]
	if 0 < len(receiver.Resources) {
		p = append(p, ascii.LF)
		p = append(p, "Resources:"...)

		for _, uri := range receiver.Resources {
			p = append(p, ascii.LF)
			p = append(p, "- "...)
			p = append(p, uri...)
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
