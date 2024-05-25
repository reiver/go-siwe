package siwe

import (
	"sourcecode.social/reiver/go-opt"
)

// Message represents a "Sign-In with Ethereum" (i.e., ERC-4361) message.
type Message struct {
	Scheme         opt.Optional[string]
	Domain         opt.Optional[string]
	Address        opt.Optional[string]
	Statement      opt.Optional[string]
	URI            opt.Optional[string]
	Version        opt.Optional[string]
	ChainID        opt.Optional[string]
	Nonce          opt.Optional[string]
	IssuedAt       opt.Optional[string]
	ExpirationTime opt.Optional[string]
	NotBefore      opt.Optional[string]
	RequestID      opt.Optional[string]
	Resources      []string
}
