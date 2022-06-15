package errors

import "errors"

type ErrorBuilder struct {
	attach string
}

func (eb *ErrorBuilder) Affixed(s string) {
	eb.attach = eb.attach + "\n" + s
}

func (eb *ErrorBuilder) Print() error {
	return errors.New(eb.attach)
}
