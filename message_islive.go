package siwe

import (
	"time"
)

func (receiver Message) IsLiveAt(when time.Time) bool {

	{
		notBeforeString, found := receiver.NotBefore.Get()
		if found {
			notBefore, err := time.Parse(time.RFC3339, notBeforeString)
			if nil != err {
				return false
			}

			if when.Before(notBefore) {
				return false
			}
		}
	}

	{
		expirationTimeString, found := receiver.ExpirationTime.Get()
		if found {
			expirationTime, err := time.Parse(time.RFC3339, expirationTimeString)
			if nil != err {
				return false
			}

			if when.After(expirationTime) {
				return false
			}
		}
	}

	return true
}

func (receiver Message) IsLiveNow() bool {
	var now time.Time = time.Now()
	return receiver.IsLiveAt(now)
}
