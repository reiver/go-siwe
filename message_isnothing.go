package siwe

import (
	"sourcecode.social/reiver/go-opt"
)

func (receiver Message) IsNothing() bool {
	var nothing opt.Optional[string]

	return	nothing == receiver.Scheme         &&
		nothing == receiver.Domain         &&
		nothing == receiver.Address        &&
		nothing == receiver.Statement      &&
		nothing == receiver.URI            &&
		nothing == receiver.Version        &&
		nothing == receiver.ChainID        &&
		nothing == receiver.Nonce          &&
		nothing == receiver.IssuedAt       &&
		nothing == receiver.ExpirationTime &&
		nothing == receiver.NotBefore      &&
		nothing == receiver.RequestID      &&
		len(receiver.Resources) <= 0
}
