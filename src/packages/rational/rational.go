package rational

type Rational struct {
	Numerator, Denominator int
}

func (r *Rational) Sum(r1 Rational) {
	r.Numerator = r.Numerator*r1.Denominator + r.Denominator*r1.Numerator
	r.Denominator = r.Denominator * r1.Denominator
}

func (r *Rational) Sub(r1 Rational) {
	r.Numerator = r.Numerator*r1.Denominator - r.Denominator*r1.Numerator
	r.Denominator = r.Denominator * r1.Denominator
}

func (r *Rational) Mult(r1 Rational) {
	r.Numerator = r.Numerator * r1.Numerator
	r.Denominator = r.Denominator * r1.Denominator
}

func (r *Rational) Div(r1 Rational) {
	r.Numerator = r.Numerator * r1.Denominator
	r.Denominator = r.Denominator * r1.Numerator
}
