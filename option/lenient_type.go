package option

import (
	"github.com/tys-muta/go-opt"
)

type lenientType bool

var _ opt.Option = (*lenientType)(nil)

func WithLenientType() opt.Option {
	return lenientType(true)
}

func (o lenientType) Validate() error {
	return nil
}

func (o lenientType) Apply(options any) {
	switch v := options.(type) {
	case *StructOptions:
		v.LenientType = o
	}
}
