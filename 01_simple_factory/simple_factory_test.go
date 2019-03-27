package _1_simple_factory

import "testing"

func TestCreate(t *testing.T) {
	var s Stringer

	s = Create("boy")
	t.Log(s.String())

	s = Create("girl")
	t.Log(s.String())

}
