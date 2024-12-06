package input

import "fmt"

type Arg struct {
	FieldName  string
	Value      string
	IsRequired bool
}

func ValidateArgsList(args []Arg) error {
	for _, arg := range args {
		if err := validateArg(arg); err != nil {
			return err
		}
	}

	return nil
}

func isArgEmpty(v string) bool {
	return v == ""
}

func validateArg(arg Arg) error {
	if arg.IsRequired && isArgEmpty(arg.Value) {
		return fmt.Errorf("missing state for %s argument", arg.FieldName)
	}

	return nil
}
