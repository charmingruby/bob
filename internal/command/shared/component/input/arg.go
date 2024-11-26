package input

import "fmt"

type arg struct {
	FieldName  string
	Value      string
	IsRequired bool
}

func validateArgsList(args []arg) error {
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

func validateArg(arg arg) error {
	if arg.IsRequired && isArgEmpty(arg.Value) {
		return fmt.Errorf("missing state for %s argument", arg.FieldName)
	}

	return nil
}
