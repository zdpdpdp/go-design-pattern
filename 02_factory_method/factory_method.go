package _2_factory_method

//Operator 被封装类的实际接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

//OperatorFactory 工厂接口
type OperatorFactory interface {
	Create() Operator
}

type PlusFactory struct {
}

func (PlusFactory) Create() Operator {
	return PlusOperator{}
}

type MinusFactory struct {
}

func (MinusFactory) Create() Operator {
	return MinusOperator{}
}

//OperatorBase 基类 由于 go 没有继承,因此使用组合基类
type OperatorBase struct {
	a, b int
}

func (o OperatorBase) SetA(a int) {
	o.a = a
}

func (o OperatorBase) SetB(b int) {
	o.b = b
}

//通过组合实现了 Operator 接口
type PlusOperator struct {
	OperatorBase
}

func (p PlusOperator) Result() int {
	return p.b + p.b
}

type MinusOperator struct {
	OperatorBase
}

func (m MinusOperator) Result() int {
	return m.a - m.b
}
