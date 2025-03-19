package validation

type Validator struct {
	Error map[string]string
}

func (v *Validator) Check(condition bool, key string, value string) {
	if !condition {
		v.Error[key] = value
	}
}
