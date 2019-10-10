package values

type OverrideValue struct {
	Key   string
	Value string
}

func (o *OverrideValue) ToString() string {
	return o.Key + "=" + o.Value
}
