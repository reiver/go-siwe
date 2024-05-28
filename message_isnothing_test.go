package siwe_test

import (
	"testing"

	"sourcecode.social/reiver/go-opt"

	"github.com/reiver/go-siwe"
)

func TestMessage_IsNothing(t *testing.T) {

	tests := []struct{
		Message siwe.Message
		Expected bool
	}{
		{
			Message: siwe.Message{},
			Expected: true,
		},



		{
			Message: siwe.Message{
				Domain: opt.Something("example.com"),
			},
			Expected: false,
		},
	}

	for testNumber, test := range tests  {

		actual := test.Message.IsNothing()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value of is-nothing is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}
