package siwe

import (
	"github.com/reiver/go-ascii"
	"sourcecode.social/reiver/go-opt"
)

func appendfield(slice []byte, left string, right opt.Optional[string]) []byte {

	slice = append(slice, ascii.LF)

	slice = append(slice, left...)

	right.WhenSomething(func(value string){
		slice = append(slice, value...)
	})

	return slice
}
