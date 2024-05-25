package siwe

func (receiver Message) String() string {
	bytes, err := receiver.SerializeSIWE()
	if nil != err {
		return ""
	}

	return string(bytes)
}
