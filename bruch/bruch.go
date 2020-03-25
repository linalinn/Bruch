package bruch

import (
	"errors"
	"fmt"
)

type Bruch struct {
	zaehler int
	nenner  int
}

func NewBruch(Zaehler, Nenner int) (*Bruch, error) {
	if Nenner == 0 {
		return nil, errors.New("Nenner darf nicht 0 sein")
	}
	return &Bruch{Zaehler, Nenner}, nil
}

func (b *Bruch) simplify() {
	gcd := b.gcd()
	b.zaehler /= gcd
	b.nenner /= gcd
}

func (b *Bruch) GetZaehler() int {
	return b.zaehler
}

func (b *Bruch) GetNenner() int {
	return b.nenner
}

func (b *Bruch) SetZaehler(zeahler int) {
	b.zaehler = zeahler
}

func (b *Bruch) SetNenner(nenner int) {
	b.nenner = nenner
}

func (b *Bruch) Add(bruch *Bruch) *Bruch {
	nenner := b.nenner * bruch.nenner
	zeahler := b.zaehler*bruch.nenner + bruch.zaehler*b.nenner

	multiplyWithMinusOne(&nenner, &zeahler)

	result := Bruch{zeahler, nenner}
	result.simplify()
	return &result

}

func (b *Bruch) Sub(bruch *Bruch) *Bruch {
	nenner := b.nenner * bruch.nenner
	zeahler := b.zaehler*bruch.nenner - bruch.zaehler*b.nenner

	multiplyWithMinusOne(&nenner, &zeahler)

	result := Bruch{zeahler, nenner}
	result.simplify()
	return &result
}

func (b *Bruch) Mul(bruch *Bruch) *Bruch {
	nenner := b.nenner * bruch.nenner
	zeahler := b.zaehler * bruch.zaehler

	multiplyWithMinusOne(&nenner, &zeahler)

	result := Bruch{zeahler, nenner}
	result.simplify()
	return &result

}

func (b *Bruch) Div(bruch *Bruch) (*Bruch, error) {

	if b.nenner == 0 || bruch.nenner == 0 {
		return nil, errors.New("Nenner darf nicht 0 sein")
	}

	nenner := b.nenner * bruch.zaehler
	zeahler := b.zaehler * bruch.nenner

	multiplyWithMinusOne(&nenner, &zeahler)

	result := Bruch{zeahler, nenner}
	result.simplify()
	return &result, nil

}

func multiplyWithMinusOne(nenner *int, zeahler *int) {
	if *nenner < 0 && *zeahler < 0 {
		*zeahler *= -1
		*nenner *= -1
	}
}

func (b Bruch) String() string {
	simpl := b
	simpl.simplify()
	dec := simpl.zaehler / simpl.nenner
	if dec != 0 {
		return fmt.Sprintf("%d %d/%d", dec, simpl.zaehler%simpl.nenner, simpl.nenner)
	}
	return fmt.Sprintf("%d/%d", b.zaehler, b.nenner)
}

func (b Bruch) gcd() int {
	for b.nenner != 0 {
		t := b.nenner
		b.nenner = b.zaehler % b.nenner
		b.zaehler = t
	}
	return b.zaehler
}
