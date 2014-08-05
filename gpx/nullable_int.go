package gpx

type NullableInt struct {
	data int
	null bool
}

func (n *NullableInt) Null() bool {
	return n.null
}

func (n *NullableInt) NotNull() bool {
	return !n.null
}

func (n *NullableInt) Value() int {
	return n.data
}

func (n *NullableInt) SetValue(data int) {
	n.data = data
}

func (n *NullableInt) SetNull() {
	n.null = true
}

func NewNullableInt(data int) *NullableInt {
	result := new(NullableInt)
	result.data = data
	result.null = false
	return result
}
