package mixnum

import (
	"packages/rational"
)

type Mixnum struct {
	Intpart int
	Ratpart rational.Rational
}

func (m *Mixnum) ToRational() (r rational.Rational) {
	r.Numerator = m.Ratpart.Numerator
	r.Denominator = m.Ratpart.Denominator

	if m.Intpart != 0 {
		r.Numerator += r.Denominator * m.Intpart
	}
	return r
}
