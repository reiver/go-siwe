package siwe_test

import (
	"testing"

	"github.com/reiver/go-siwe"
	"sourcecode.social/reiver/go-opt"
)

func TestMessage_String(t *testing.T) {

	tests := []struct{
		Message siwe.Message
		Expected string
	}{
		{
			Expected:
//				"\x19Ethereum Signed Message:"                      +"\n"+
//				"99"                                               +""+
				" wants you to sign in with your Ethereum account:" +"\n"+
				""                                                  +"\n"+ // address
				""                                                  +"\n"+
				// statement
				""                                                  +"\n"+
				"URI: "                                             +"\n"+
				"Version: "                                         +"\n"+
				"Chain ID: "                                        +"\n"+
				"Nonce: "                                           +"\n"+
				"Issued At: ",
		},
		{
			Message: siwe.Message{
				Domain:  opt.Something("example.com"),
				Address: opt.Something("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				Statement: opt.Something("I accept the ExampleOrg Terms of Service: https://example.com/tos"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("1"),
				Nonce: opt.Something("32891756"),
				IssuedAt: opt.Something("2021-09-30T16:25:24Z"),
				Resources: []string{
					"ipfs://bafybeiemxf5abjwjbikoz4mc3a3dla6ual3jsgpdr4cjr3oz3evfyavhwq/",
					"https://example.com/my-web2-claim.json",
				},
			},
			Expected:
//"\x19Ethereum Signed Message:\n395"+
`example.com wants you to sign in with your Ethereum account:
0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2

I accept the ExampleOrg Terms of Service: https://example.com/tos

URI: https://example.com/login
Version: 1
Chain ID: 1
Nonce: 32891756
Issued At: 2021-09-30T16:25:24Z
Resources:
- ipfs://bafybeiemxf5abjwjbikoz4mc3a3dla6ual3jsgpdr4cjr3oz3evfyavhwq/
- https://example.com/my-web2-claim.json`,
		},
		{
			Message: siwe.Message{
				Domain:  opt.Something("example.com:3388"),
				Address: opt.Something("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				Statement: opt.Something("I accept the ExampleOrg Terms of Service: https://example.com/tos"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("1"),
				Nonce: opt.Something("32891756"),
				IssuedAt: opt.Something("2021-09-30T16:25:24Z"),
				Resources: []string{
					"ipfs://bafybeiemxf5abjwjbikoz4mc3a3dla6ual3jsgpdr4cjr3oz3evfyavhwq/",
					"https://example.com/my-web2-claim.json",
				},
			},
			Expected:
//"\x19Ethereum Signed Message:\n400"+
`example.com:3388 wants you to sign in with your Ethereum account:
0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2

I accept the ExampleOrg Terms of Service: https://example.com/tos

URI: https://example.com/login
Version: 1
Chain ID: 1
Nonce: 32891756
Issued At: 2021-09-30T16:25:24Z
Resources:
- ipfs://bafybeiemxf5abjwjbikoz4mc3a3dla6ual3jsgpdr4cjr3oz3evfyavhwq/
- https://example.com/my-web2-claim.json`,
		},
	}

	for testNumber, test := range tests {

		actual := test.Message.String()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual marshaled siwe message is not what was expected", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("EXPECTED:\n%s", expected)
			t.Logf("ACTUAL:\n%s", actual)
			continue
		}
	}
}
