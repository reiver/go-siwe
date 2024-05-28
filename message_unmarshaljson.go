package siwe

import (
	"encoding/json"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-opt"
)

const (
	errNilReceiver = erorr.Error("siwe: nil receiver")
)

func (receiver *Message) UnmarshalJSON(p []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	var values map[string]interface{} = map[string]interface{}{}
	{
		err := json.Unmarshal(p, &values)
		if nil != err {
			return err
		}
	}

	{
		const name string = "scheme"
		option, err := stringvalue(values, name)
		if nil != err {
			return err
		}
		receiver.Scheme = option
	}

	{
		const name string = "domain"
		option, err := stringvalue(values, name)
		if nil != err {
			return err
		}
		receiver.Domain = option
	}

	{
		const name string = "address"
		option, err := stringvalue(values, name)
		if nil != err {
			return err
		}
		receiver.Address = option
	}

	{
		const name string = "statement"
		option, err := stringvalue(values, name)
		if nil != err {
			return err
		}
		receiver.Statement = option
	}

	{
		const name string = "uri"
		option, err := stringvalue(values, name)
		if nil != err {
			return err
		}
		receiver.URI = option
	}

	{
		const name string = "version"
		option, err := stringvalue(values, name)
		if nil != err {
			return err
		}
		receiver.Version = option
	}

	{
		var option opt.Optional[string]
		var err error
		{
			const name string = "chain-id"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "chain_id"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "chainid"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "chainID"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "chainId"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		receiver.ChainID = option
	}

	{
		const name string = "nonce"
		option, err := stringvalue(values, name)
		if nil != err {
			return err
		}
		receiver.Nonce = option
	}

	{
		var option opt.Optional[string]
		var err error
		{
			const name string = "issued-at"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "issued_at"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "issuedat"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "issuedAt"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		receiver.IssuedAt = option
	}

	{
		var option opt.Optional[string]
		var err error
		{
			const name string = "expiration-time"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "expiration_time"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "expirationtime"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "expirationTime"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		receiver.ExpirationTime = option
	}

	{
		var option opt.Optional[string]
		var err error
		{
			const name string = "not-before"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "not_before"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "notbefore"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "notBefore"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		receiver.NotBefore = option
	}

	{
		var option opt.Optional[string]
		var err error
		{
			const name string = "request-id"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "request_id"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "requestid"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		if opt.Nothing[string]() == option {
			const name string = "requestId"
			option, err = stringvalue(values, name)
			if nil != err {
				return err
			}
		}
		receiver.RequestID = option
	}

	{
		const name string = "resources"
		stringslice, err := stringslicevalue(values, name)
		if nil != err {
			return err
		}
		receiver.Resources = stringslice
	}

	return nil
}


func stringvalue(values map[string]interface{}, name string) (opt.Optional[string], error) {

	valueInterface, found := values[name]
	if !found {
		return opt.Nothing[string](), nil
	}

	value, casted := valueInterface.(string)
	if !casted {
		return opt.Nothing[string](), erorr.Errorf("siwe: bad type for '%s' — expected 'string' but actually got '%T'", name, valueInterface)
	}

	return opt.Something(value), nil
}

func stringslicevalue(values map[string]interface{}, name string) ([]string, error) {

	valueInterface, found := values[name]
	if !found {
		return nil, nil
	}
	if nil == valueInterface {
		return nil, nil
	}

	value, casted := valueInterface.([]interface{})
	if !casted {
		return nil, erorr.Errorf("siwe: bad type for '%s' — expected '[]interface{}' but actually got '%T'", name, valueInterface)
	}
	if len(value) <= 0 {
		return nil, nil
	}

	var stringslice []string
	{
		for index, datumInterface := range value {
			datum, casted := datumInterface.(string)
			if !casted {
				return nil, erorr.Errorf("siwe: bad type for '%s'[%d] — expected 'string' but actually got '%T'", name, index, valueInterface)
			}

			stringslice = append(stringslice, datum)
		}
	}

	return stringslice, nil
}
