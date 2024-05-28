package siwe_test

import (
	"testing"

	"reflect"

	"sourcecode.social/reiver/go-opt"

	"github.com/reiver/go-siwe"
)

func TestMessage_UnmarshalJSON(t *testing.T) {

	tests := []struct{
		JSON string
		Expected siwe.Message
	}{
		{
			JSON: `{}`,
		},



		{
			JSON:
				`{`+
					`"domain":"example.com"`+
				`}`,
			Expected: siwe.Message{
				Domain: opt.Something("example.com"),
			},
		},
		{
			JSON:
				`{`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
				`}`,
			Expected: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
			},
		},
		{
			JSON:
				`{`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"uri":"https://example.com/login"`+
				`}`,
			Expected: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				URI: opt.Something("https://example.com/login"),
			},
		},
		{
			JSON:
				`{`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
				`}`,
			Expected: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
			},
		},
		{
			JSON:
				`{`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
					`,`+
					`"chainid":"17000"`+
				`}`,
			Expected: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
			},
		},
		{
			JSON:
				`{`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
					`,`+
					`"chainid":"17000"`+
					`,`+
					`"nonce":"1789452"`+
				`}`,
			Expected: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
				Nonce: opt.Something("1789452"),
			},
		},
		{
			JSON:
				`{`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
					`,`+
					`"chainid":"17000"`+
					`,`+
					`"nonce":"1789452"`+
					`,`+
					`"issued-at":"2024-05-27T06:35:07Z"`+
				`}`,
			Expected: siwe.Message{
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
				Nonce: opt.Something("1789452"),
				IssuedAt: opt.Something("2024-05-27T06:35:07Z"),
			},
		},









		{
			JSON:
				`{`+
					`"scheme":"https"`+
					`,`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
					`,`+
					`"chainid":"17000"`+
					`,`+
					`"nonce":"1789452"`+
					`,`+
					`"issued-at":"2024-05-27T06:35:07Z"`+
				`}`,
			Expected: siwe.Message{
				Scheme: opt.Something("https"),
				Domain: opt.Something("example.com"),
				Address: opt.Something("0xde709f2102306220921060314715629080e2fB77"),
				URI: opt.Something("https://example.com/login"),
				Version: opt.Something("1"),
				ChainID: opt.Something("17000"),
				Nonce: opt.Something("1789452"),
				IssuedAt: opt.Something("2024-05-27T06:35:07Z"),
			},
		},
		{
			JSON:
				`{`+
					`"scheme":"https"`+
					`,`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"statement":"It's peanut butter jelly time!"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
					`,`+
					`"chainid":"17000"`+
					`,`+
					`"nonce":"1789452"`+
					`,`+
					`"issued-at":"2024-05-27T06:35:07Z"`+
				`}`,
			Expected: siwe.Message{
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
		},
		{
			JSON:
				`{`+
					`"scheme":"https"`+
					`,`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"statement":"It's peanut butter jelly time!"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
					`,`+
					`"chainid":"17000"`+
					`,`+
					`"nonce":"1789452"`+
					`,`+
					`"issued-at":"2024-05-27T06:35:07Z"`+
					`,`+
					`"expiration-time":"2024-06-01T00:00:00Z"`+
				`}`,
			Expected: siwe.Message{
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
		},
		{
			JSON:
				`{`+
					`"scheme":"https"`+
					`,`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"statement":"It's peanut butter jelly time!"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
					`,`+
					`"chainid":"17000"`+
					`,`+
					`"nonce":"1789452"`+
					`,`+
					`"issued-at":"2024-05-27T06:35:07Z"`+
					`,`+
					`"expiration-time":"2024-06-01T00:00:00Z"`+
					`,`+
					`"not-before":"2024-05-29T09:10:11Z"`+
				`}`,
			Expected: siwe.Message{
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
		},
		{
			JSON:
				`{`+
					`"scheme":"https"`+
					`,`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"statement":"It's peanut butter jelly time!"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
					`,`+
					`"chainid":"17000"`+
					`,`+
					`"nonce":"1789452"`+
					`,`+
					`"issued-at":"2024-05-27T06:35:07Z"`+
					`,`+
					`"expiration-time":"2024-06-01T00:00:00Z"`+
					`,`+
					`"not-before":"2024-05-29T09:10:11Z"`+
					`,`+
					`"request-id":"gR9TZieEe1NCLzcEYTzh"`+
				`}`,
			Expected: siwe.Message{
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
		},









		{
			JSON:
				`{`+
					`"scheme":"https"`+
					`,`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"statement":"It's peanut butter jelly time!"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
					`,`+
					`"chainid":"17000"`+
					`,`+
					`"nonce":"1789452"`+
					`,`+
					`"issued-at":"2024-05-27T06:35:07Z"`+
					`,`+
					`"expiration-time":"2024-06-01T00:00:00Z"`+
					`,`+
					`"not-before":"2024-05-29T09:10:11Z"`+
					`,`+
					`"request-id":"gR9TZieEe1NCLzcEYTzh"`+
					`,`+
					`"resources":null`+
				`}`,
			Expected: siwe.Message{
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
		},
		{
			JSON:
				`{`+
					`"scheme":"https"`+
					`,`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"statement":"It's peanut butter jelly time!"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
					`,`+
					`"chainid":"17000"`+
					`,`+
					`"nonce":"1789452"`+
					`,`+
					`"issued-at":"2024-05-27T06:35:07Z"`+
					`,`+
					`"expiration-time":"2024-06-01T00:00:00Z"`+
					`,`+
					`"not-before":"2024-05-29T09:10:11Z"`+
					`,`+
					`"request-id":"gR9TZieEe1NCLzcEYTzh"`+
					`,`+
					`"resources":[]`+
				`}`,
			Expected: siwe.Message{
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
		},
		{
			JSON:
				`{`+
					`"scheme":"https"`+
					`,`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"statement":"It's peanut butter jelly time!"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
					`,`+
					`"chainid":"17000"`+
					`,`+
					`"nonce":"1789452"`+
					`,`+
					`"issued-at":"2024-05-27T06:35:07Z"`+
					`,`+
					`"expiration-time":"2024-06-01T00:00:00Z"`+
					`,`+
					`"not-before":"2024-05-29T09:10:11Z"`+
					`,`+
					`"request-id":"gR9TZieEe1NCLzcEYTzh"`+
					`,`+
					`"resources":["https://example.com/apple"]`+
				`}`,
			Expected: siwe.Message{
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
		},
		{
			JSON:
				`{`+
					`"scheme":"https"`+
					`,`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"statement":"It's peanut butter jelly time!"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
					`,`+
					`"chainid":"17000"`+
					`,`+
					`"nonce":"1789452"`+
					`,`+
					`"issued-at":"2024-05-27T06:35:07Z"`+
					`,`+
					`"expiration-time":"2024-06-01T00:00:00Z"`+
					`,`+
					`"not-before":"2024-05-29T09:10:11Z"`+
					`,`+
					`"request-id":"gR9TZieEe1NCLzcEYTzh"`+
					`,`+
					`"resources":["https://example.com/apple", "https://example.com/banana"]`+
				`}`,
			Expected: siwe.Message{
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
		},
		{
			JSON:
				`{`+
					`"scheme":"https"`+
					`,`+
					`"domain":"example.com"`+
					`,`+
					`"address":"0xde709f2102306220921060314715629080e2fB77"`+
					`,`+
					`"statement":"It's peanut butter jelly time!"`+
					`,`+
					`"uri":"https://example.com/login"`+
					`,`+
					`"version":"1"`+
					`,`+
					`"chainid":"17000"`+
					`,`+
					`"nonce":"1789452"`+
					`,`+
					`"issued-at":"2024-05-27T06:35:07Z"`+
					`,`+
					`"expiration-time":"2024-06-01T00:00:00Z"`+
					`,`+
					`"not-before":"2024-05-29T09:10:11Z"`+
					`,`+
					`"request-id":"gR9TZieEe1NCLzcEYTzh"`+
					`,`+
					`"resources":["https://example.com/apple", "https://example.com/banana", "https://example.com/cherry"]`+
				`}`,
			Expected: siwe.Message{
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
		},
	}

	for testNumber, test := range tests {

		var actual siwe.Message

		err := actual.UnmarshalJSON([]byte(test.JSON))
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual message is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}

