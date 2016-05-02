package hydrant

func Int(value int) *int {
	return &value
}

func String(value string) *string {
	return &value
}

func Bool(value bool) *bool {
	return &value
}

func Float(value float64) *float64 {
	return &value
}
