package siwe_test

import (
	"testing"

	"fmt"

	"sourcecode.social/reiver/go-opt"

	"github.com/reiver/go-siwe"
)

func TestMessage_EthereumSingableMessage(t *testing.T) {

	tests := []struct{
		Message siwe.Message
		Expected string
	}{
		{
			Message: siwe.Message{},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				"99"+
				" wants you to sign in with your Ethereum account:"+
				"\n"+
				"\n"+
				""+""+
				"\n"+
				"\n"+
				"URI: "+
				"\n"+
				"Version: "+
				"\n"+
				"Chain ID: "+
				"\n"+
				"Nonce: "+
				"\n"+
				"Issued At: ",
		},



		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d", 99+len("example.com"))+
				"example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				""+
				"\n"+
				"\n"+
				""+""+
				"\n"+
				"URI: "+
				"\n"+
				"Version: "+
				"\n"+
				"Chain ID: "+
				"\n"+
				"Nonce: "+
				"\n"+
				"Issued At: ",
		},
		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77"))+
				"example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				""+""+
				"\n"+
				"URI: "+
				"\n"+
				"Version: "+
				"\n"+
				"Chain ID: "+
				"\n"+
				"Nonce: "+
				"\n"+
				"Issued At: ",
		},
		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				URI: opt.Something("https://example.com/login"),
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77")+
					len("https://example.com/login"))+
				"example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				""+""+
				"\n"+
				"URI: https://example.com/login"+
				"\n"+
				"Version: "+
				"\n"+
				"Chain ID: "+
				"\n"+
				"Nonce: "+
				"\n"+
				"Issued At: ",
		},
		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77")+
					len("https://example.com/login")+
					len("1"))+
				"example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				""+""+
				"\n"+
				"URI: https://example.com/login"+
				"\n"+
				"Version: 1"+
				"\n"+
				"Chain ID: "+
				"\n"+
				"Nonce: "+
				"\n"+
				"Issued At: ",
		},
		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77")+
					len("https://example.com/login")+
					len("1")+
					len("17000"))+
				"example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				""+""+
				"\n"+
				"URI: https://example.com/login"+
				"\n"+
				"Version: 1"+
				"\n"+
				"Chain ID: 17000"+
				"\n"+
				"Nonce: "+
				"\n"+
				"Issued At: ",
		},
		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
				Nonce: opt.Something("1789452"),
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77")+
					len("https://example.com/login")+
					len("1")+
					len("17000")+
					len("1789452"))+
				"example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				""+""+
				"\n"+
				"URI: https://example.com/login"+
				"\n"+
				"Version: 1"+
				"\n"+
				"Chain ID: 17000"+
				"\n"+
				"Nonce: 1789452"+
				"\n"+
				"Issued At: ",
		},
		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
				Nonce: opt.Something("1789452"),
				IssuedAt: opt.Something("2024-05-27T06:35:07Z"),
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77")+
					len("https://example.com/login")+
					len("1")+
					len("17000")+
					len("1789452")+
					len("2024-05-27T06:35:07Z"))+
				"example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				""+""+
				"\n"+
				"URI: https://example.com/login"+
				"\n"+
				"Version: 1"+
				"\n"+
				"Chain ID: 17000"+
				"\n"+
				"Nonce: 1789452"+
				"\n"+
				"Issued At: 2024-05-27T06:35:07Z",
		},









		{
			Message: siwe.Message{
				Scheme: opt.Something("https"),
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
				Nonce: opt.Something("1789452"),
				IssuedAt: opt.Something("2024-05-27T06:35:07Z"),
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("https")+len("://")+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77")+
					len("https://example.com/login")+
					len("1")+
					len("17000")+
					len("1789452")+
					len("2024-05-27T06:35:07Z"))+
				"https://example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				""+""+
				"\n"+
				"URI: https://example.com/login"+
				"\n"+
				"Version: 1"+
				"\n"+
				"Chain ID: 17000"+
				"\n"+
				"Nonce: 1789452"+
				"\n"+
				"Issued At: 2024-05-27T06:35:07Z",
		},
		{
			Message: siwe.Message{
				Scheme: opt.Something("https"),
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				Statement: opt.Something("It's peanut butter jelly time!"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
				Nonce: opt.Something("1789452"),
				IssuedAt: opt.Something("2024-05-27T06:35:07Z"),
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("https")+len("://")+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77")+
					len("It's peanut butter jelly time!")+len("\n")+
					len("https://example.com/login")+
					len("1")+
					len("17000")+
					len("1789452")+
					len("2024-05-27T06:35:07Z"))+
				"https://example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				"It's peanut butter jelly time!"+"\n"+
				"\n"+
				"URI: https://example.com/login"+
				"\n"+
				"Version: 1"+
				"\n"+
				"Chain ID: 17000"+
				"\n"+
				"Nonce: 1789452"+
				"\n"+
				"Issued At: 2024-05-27T06:35:07Z",
		},
		{
			Message: siwe.Message{
				Scheme: opt.Something("https"),
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				Statement: opt.Something("It's peanut butter jelly time!"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
				Nonce: opt.Something("1789452"),
				IssuedAt: opt.Something("2024-05-27T06:35:07Z"),
				ExpirationTime: opt.Something("2024-06-01T00:00:00Z"),
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("https")+len("://")+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77")+
					len("It's peanut butter jelly time!")+len("\n")+
					len("https://example.com/login")+
					len("1")+
					len("17000")+
					len("1789452")+
					len("2024-05-27T06:35:07Z")+
					len("\n")+len("Expiration Time: ")+len("2024-06-01T00:00:00Z"))+
				"https://example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				"It's peanut butter jelly time!"+"\n"+
				"\n"+
				"URI: https://example.com/login"+
				"\n"+
				"Version: 1"+
				"\n"+
				"Chain ID: 17000"+
				"\n"+
				"Nonce: 1789452"+
				"\n"+
				"Issued At: 2024-05-27T06:35:07Z"+
				"\n"+
				"Expiration Time: 2024-06-01T00:00:00Z",
		},
		{
			Message: siwe.Message{
				Scheme: opt.Something("https"),
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				Statement: opt.Something("It's peanut butter jelly time!"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
				Nonce: opt.Something("1789452"),
				IssuedAt: opt.Something("2024-05-27T06:35:07Z"),
				ExpirationTime: opt.Something("2024-06-01T00:00:00Z"),
				NotBefore: opt.Something("2024-05-29T09:10:11Z"),
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("https")+len("://")+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77")+
					len("It's peanut butter jelly time!")+len("\n")+
					len("https://example.com/login")+
					len("1")+
					len("17000")+
					len("1789452")+
					len("2024-05-27T06:35:07Z")+
					len("\n")+len("Expiration Time: ")+len("2024-06-01T00:00:00Z")+
					len("\n")+len("Not Before: ")+len("2024-05-29T09:10:11Z"))+
				"https://example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				"It's peanut butter jelly time!"+"\n"+
				"\n"+
				"URI: https://example.com/login"+
				"\n"+
				"Version: 1"+
				"\n"+
				"Chain ID: 17000"+
				"\n"+
				"Nonce: 1789452"+
				"\n"+
				"Issued At: 2024-05-27T06:35:07Z"+
				"\n"+
				"Expiration Time: 2024-06-01T00:00:00Z"+
				"\n"+
				"Not Before: 2024-05-29T09:10:11Z",
		},
		{
			Message: siwe.Message{
				Scheme: opt.Something("https"),
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				Statement: opt.Something("It's peanut butter jelly time!"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
				Nonce: opt.Something("1789452"),
				IssuedAt: opt.Something("2024-05-27T06:35:07Z"),
				ExpirationTime: opt.Something("2024-06-01T00:00:00Z"),
				NotBefore: opt.Something("2024-05-29T09:10:11Z"),
				RequestID: opt.Something("gR9TZieEe1NCLzcEYTzh"),
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("https")+len("://")+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77")+
					len("It's peanut butter jelly time!")+len("\n")+
					len("https://example.com/login")+
					len("1")+
					len("17000")+
					len("1789452")+
					len("2024-05-27T06:35:07Z")+
					len("\n")+len("Expiration Time: ")+len("2024-06-01T00:00:00Z")+
					len("\n")+len("Not Before: ")+len("2024-05-29T09:10:11Z")+
					len("\n")+len("Request ID: ")+len("gR9TZieEe1NCLzcEYTzh"))+
				"https://example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				"It's peanut butter jelly time!"+"\n"+
				"\n"+
				"URI: https://example.com/login"+
				"\n"+
				"Version: 1"+
				"\n"+
				"Chain ID: 17000"+
				"\n"+
				"Nonce: 1789452"+
				"\n"+
				"Issued At: 2024-05-27T06:35:07Z"+
				"\n"+
				"Expiration Time: 2024-06-01T00:00:00Z"+
				"\n"+
				"Not Before: 2024-05-29T09:10:11Z"+
				"\n"+
				"Request ID: gR9TZieEe1NCLzcEYTzh",
		},









		{
			Message: siwe.Message{
				Scheme: opt.Something("https"),
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				Statement: opt.Something("It's peanut butter jelly time!"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
				Nonce: opt.Something("1789452"),
				IssuedAt: opt.Something("2024-05-27T06:35:07Z"),
				ExpirationTime: opt.Something("2024-06-01T00:00:00Z"),
				NotBefore: opt.Something("2024-05-29T09:10:11Z"),
				RequestID: opt.Something("gR9TZieEe1NCLzcEYTzh"),
				Resources: []string{
					"https://example.com/apple",
				},
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("https")+len("://")+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77")+
					len("It's peanut butter jelly time!")+len("\n")+
					len("https://example.com/login")+
					len("1")+
					len("17000")+
					len("1789452")+
					len("2024-05-27T06:35:07Z")+
					len("\n")+len("Expiration Time: ")+len("2024-06-01T00:00:00Z")+
					len("\n")+len("Not Before: ")+len("2024-05-29T09:10:11Z")+
					len("\n")+len("Request ID: ")+len("gR9TZieEe1NCLzcEYTzh")+
					len("\n")+len("Resources:")+
					len("\n")+len("- ")+len("https://example.com/apple"))+
				"https://example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				"It's peanut butter jelly time!"+"\n"+
				"\n"+
				"URI: https://example.com/login"+
				"\n"+
				"Version: 1"+
				"\n"+
				"Chain ID: 17000"+
				"\n"+
				"Nonce: 1789452"+
				"\n"+
				"Issued At: 2024-05-27T06:35:07Z"+
				"\n"+
				"Expiration Time: 2024-06-01T00:00:00Z"+
				"\n"+
				"Not Before: 2024-05-29T09:10:11Z"+
				"\n"+
				"Request ID: gR9TZieEe1NCLzcEYTzh"+
				"\n"+
				"Resources:"+
				"\n"+
				"- "+"https://example.com/apple",
		},
		{
			Message: siwe.Message{
				Scheme: opt.Something("https"),
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				Statement: opt.Something("It's peanut butter jelly time!"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
				Nonce: opt.Something("1789452"),
				IssuedAt: opt.Something("2024-05-27T06:35:07Z"),
				ExpirationTime: opt.Something("2024-06-01T00:00:00Z"),
				NotBefore: opt.Something("2024-05-29T09:10:11Z"),
				RequestID: opt.Something("gR9TZieEe1NCLzcEYTzh"),
				Resources: []string{
					"https://example.com/apple",
					"https://example.com/banana",
				},
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("https")+len("://")+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77")+
					len("It's peanut butter jelly time!")+len("\n")+
					len("https://example.com/login")+
					len("1")+
					len("17000")+
					len("1789452")+
					len("2024-05-27T06:35:07Z")+
					len("\n")+len("Expiration Time: ")+len("2024-06-01T00:00:00Z")+
					len("\n")+len("Not Before: ")+len("2024-05-29T09:10:11Z")+
					len("\n")+len("Request ID: ")+len("gR9TZieEe1NCLzcEYTzh")+
					len("\n")+len("Resources:")+
					len("\n")+len("- ")+len("https://example.com/apple")+
					len("\n")+len("- ")+len("https://example.com/banana"))+
				"https://example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				"It's peanut butter jelly time!"+"\n"+
				"\n"+
				"URI: https://example.com/login"+
				"\n"+
				"Version: 1"+
				"\n"+
				"Chain ID: 17000"+
				"\n"+
				"Nonce: 1789452"+
				"\n"+
				"Issued At: 2024-05-27T06:35:07Z"+
				"\n"+
				"Expiration Time: 2024-06-01T00:00:00Z"+
				"\n"+
				"Not Before: 2024-05-29T09:10:11Z"+
				"\n"+
				"Request ID: gR9TZieEe1NCLzcEYTzh"+
				"\n"+
				"Resources:"+
				"\n"+
				"- "+"https://example.com/apple"+
				"\n"+
				"- "+"https://example.com/banana",
		},
		{
			Message: siwe.Message{
				Scheme: opt.Something("https"),
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				Statement: opt.Something("It's peanut butter jelly time!"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
				Nonce: opt.Something("1789452"),
				IssuedAt: opt.Something("2024-05-27T06:35:07Z"),
				ExpirationTime: opt.Something("2024-06-01T00:00:00Z"),
				NotBefore: opt.Something("2024-05-29T09:10:11Z"),
				RequestID: opt.Something("gR9TZieEe1NCLzcEYTzh"),
				Resources: []string{
					"https://example.com/apple",
					"https://example.com/banana",
					"https://example.com/cherry",
				},
			},
			Expected:
				"\x19Ethereum Signed Message:\n"+
				fmt.Sprintf("%d",
					99+
					len("https")+len("://")+
					len("example.com")+
					len("0xde709f2102306220921060314715629080e2fB77")+
					len("It's peanut butter jelly time!")+len("\n")+
					len("https://example.com/login")+
					len("1")+
					len("17000")+
					len("1789452")+
					len("2024-05-27T06:35:07Z")+
					len("\n")+len("Expiration Time: ")+len("2024-06-01T00:00:00Z")+
					len("\n")+len("Not Before: ")+len("2024-05-29T09:10:11Z")+
					len("\n")+len("Request ID: ")+len("gR9TZieEe1NCLzcEYTzh")+
					len("\n")+len("Resources:")+
					len("\n")+len("- ")+len("https://example.com/apple")+
					len("\n")+len("- ")+len("https://example.com/banana")+
					len("\n")+len("- ")+len("https://example.com/cherry"))+
				"https://example.com wants you to sign in with your Ethereum account:"+
				"\n"+
				"0xde709f2102306220921060314715629080e2fB77"+
				"\n"+
				"\n"+
				"It's peanut butter jelly time!"+"\n"+
				"\n"+
				"URI: https://example.com/login"+
				"\n"+
				"Version: 1"+
				"\n"+
				"Chain ID: 17000"+
				"\n"+
				"Nonce: 1789452"+
				"\n"+
				"Issued At: 2024-05-27T06:35:07Z"+
				"\n"+
				"Expiration Time: 2024-06-01T00:00:00Z"+
				"\n"+
				"Not Before: 2024-05-29T09:10:11Z"+
				"\n"+
				"Request ID: gR9TZieEe1NCLzcEYTzh"+
				"\n"+
				"Resources:"+
				"\n"+
				"- "+"https://example.com/apple"+
				"\n"+
				"- "+"https://example.com/banana"+
				"\n"+
				"- "+"https://example.com/cherry",
		},
	}

	for testNumber, test := range tests {


		actualBytes, err := test.Message.EthereumSignableMessage()

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		actual   := string(actualBytes)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual ethereum-signable-message is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("EXPECTED: \n%s", expected)
			t.Logf("ACTUAL:   \n%s", actual)
			continue
		}
	}
}
