package option

import (
	"github.com/tys-muta/go-opt"
)

type privateField bool

var _ opt.Option = (*privateField)(nil)

func WithPrivateField() opt.Option {
	return privateField(true)
}

func (o privateField) Validate() error {
	return nil
}

func (o privateField) Apply(options any) {
	switch v := options.(type) {
	case *StructOptions:
		v.PrivateField = o
	}
}
