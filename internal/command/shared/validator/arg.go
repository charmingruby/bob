package validator

import "fmt"

type Arg struct {
	FieldName  string
	Value      string
	IsRequired bool
}

func NewArg(fieldName string, value string, isRequired bool) Arg {
	return Arg{
		FieldName:  fieldName,
		IsRequired: isRequired,
	}
}

func ValidateArgsList(args []*Arg) error {
	for _, arg := range args {
		if arg.IsRequired && isArgEmpty(arg.Value) {
			return fmt.Errorf("missing state for %s argument", arg.FieldName)
		}
	}

	return nil
}

func isArgEmpty(v string) bool {
	return v == ""
}
