package validator

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{make(map[string]string)}
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// adds validation if error is not ok
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// regex check would be really useful
// add if needed
