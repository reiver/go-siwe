package siwe_test

import (
	"testing"

	"sourcecode.social/reiver/go-opt"

	"github.com/reiver/go-siwe"
)

func TestMessage_Validate(t *testing.T) {

	tests := []struct{
		Message siwe.Message
		Expected error
	}{
		{
			Expected: siwe.ErrDomainNotFound,
		},
		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
			},
			Expected: siwe.ErrAddressNotFound,
		},
		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
			},
			Expected: siwe.ErrURINotFound,
		},
		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				URI: opt.Something("https://example.com/login"),
			},
			Expected: siwe.ErrVersionNotFound,
		},
		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
			},
			Expected: siwe.ErrChainIDNotFound,
		},
		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("1"),
			},
			Expected: siwe.ErrNonceNotFound,
		},
		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("1"),
				Nonce: opt.Something("32891756"),
			},
			Expected: siwe.ErrIssuedAtNotFound,
		},
		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("1"),
				Nonce: opt.Something("32891756"),
				IssuedAt: opt.Something("2021-09-30T16:25:24Z"),
			},
			Expected: nil,
		},
	}

	for testNumber, test := range tests {

		actual := test.Message.Validate()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED-ERROR: (%T) %s", expected, expected)
			t.Logf("ACTUAL-ERROR:   (%T) %s", actual, actual)
			continue
		}
	}
}
