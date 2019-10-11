package values

type OverrideValue struct {
	Key   string
	Value string
}

// ToString generate string value from key value.
// For example,  `{Key: "hoge", Value: "fuga"}` is converted to `hoge=fuga`.
func (o *OverrideValue) ToString() string {
	return o.Key + "=" + o.Value
}
