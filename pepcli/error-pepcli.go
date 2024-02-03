package pepcli

import "fmt"

type pepError struct {
	errMsgs []string
	count   uint
}

func (p *pepError) addErrorString(s string) {
	p.errMsgs = append(p.errMsgs, s)
	p.count++
}

func (p *pepError) getError() error {
	if p.count == 0 {
		return nil
	}
	var s string
	for _, em := range p.errMsgs {
		s += em
	}
	return fmt.Errorf("%s", s)
}
