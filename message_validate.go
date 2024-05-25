package siwe

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	ErrDomainNotFound   = erorr.Error("siwe: ‘domain’ not found")
	ErrAddressNotFound  = erorr.Error("siwe: ‘address’ not found")
	ErrURINotFound      = erorr.Error("siwe: ‘uri’ not found")
	ErrVersionNotFound  = erorr.Error("siwe: ‘version’ not found")
	ErrChainIDNotFound  = erorr.Error("siwe: ‘chain-id’ not found")
	ErrNonceNotFound    = erorr.Error("siwe: ‘nonce’ not found")
	ErrIssuedAtNotFound = erorr.Error("siwe: ‘issued-at’ not found")
)

// Values of some fields in a "Sign-In with Ethereum" (i.e., ERC-4361) message are REQUIRED (while others are OPTIONAL).
// Validate returns an error if any of the REQUIRED fields are missing.
func (receiver Message) Validate() error {
	{
		_, found := receiver.Domain.Get()
		if !found {
			return ErrDomainNotFound
		}
	}

	{
		_, found := receiver.Address.Get()
		if !found {
			return ErrAddressNotFound
		}
	}

	{
		_, found := receiver.URI.Get()
		if !found {
			return ErrURINotFound
		}
	}

	{
		_, found := receiver.Version.Get()
		if !found {
			return ErrVersionNotFound
		}
	}

	{
		_, found := receiver.ChainID.Get()
		if !found {
			return ErrChainIDNotFound
		}
	}

	{
		_, found := receiver.Nonce.Get()
		if !found {
			return ErrNonceNotFound
		}
	}

	{
		_, found := receiver.IssuedAt.Get()
		if !found {
			return ErrIssuedAtNotFound
		}
	}

	return nil
}
