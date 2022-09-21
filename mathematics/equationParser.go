package mathematics

type OperandType int

const (
	ValueType OperandType = iota
	ExpressionType
)

type Operand struct {
	Type       OperandType
	Value      float32
	Expression *Expression
}

type Expression struct {
	Operator     string
	LeftOperand  *Expression
	RightOperand *Expression
}

func ParseEquation(input string) *Expression {
	res := new(Expression)

	return res
}
