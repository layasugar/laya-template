package rotatelog

type Interface interface {
	Name() string
	Value() interface{}
}

type OptionCtx struct {
	name  string
	value interface{}
}

func NewOption(name string, value interface{}) *OptionCtx {
	return &OptionCtx{
		name:  name,
		value: value,
	}
}

func (o *OptionCtx) Name() string {
	return o.name
}

func (o *OptionCtx) Value() interface{} {
	return o.value
}
