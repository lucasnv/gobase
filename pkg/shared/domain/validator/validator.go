package validator

// Validation Error
type ValidationError interface {
	Field() string
	Error() string
}

// Validator struct
type Validator interface {

	// This function is responsable to validate a single var
	Var(field string, value any, rules string) ValidationError

	// This function is responsable to get signgle message error from the validator
	//Todo: Struct(s any) error
}
