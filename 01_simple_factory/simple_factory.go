package _1_simple_factory

type Stringer interface {
	String() string
}

type Boy struct {
}

func (Boy) String() string {
	return "i am boy"
}

type Girl struct {
}

func (Girl) String() string {
	return "i am girl"
}

func Create(name string) Stringer {
	switch name {
	case "boy":
		return Boy{}
	case "girl":
		return Girl{}
	default:
		panic("not this")
	}
}
