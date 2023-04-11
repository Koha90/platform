// Package validation ...
package validation

// Validator ...
type Validator interface {
	Validate(data interface{}) (ok bool, errs []ValidationError)
}

// ValidationError ...
type ValidationError struct {
	FieldName string
	Error     error
}

// ValidatorFunc ...
type ValidatorFunc func(fieldName string, value interface{}, arg string) (bool, error)

// DefaultValidators ...
func DefaultValidators() map[string]ValidatorFunc {
	return map[string]ValidatorFunc{
		"required": required,
		"min":      min,
	}
}
