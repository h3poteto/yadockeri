package values

type OverrideValue struct {
	Key   string
	Value string
}

func (o *OverrideValue) ToMap() map[string]interface{} {
	// TODO: Parse nested map
	m := map[string]interface{}{
		o.Key: o.Value,
	}
	return m
}
