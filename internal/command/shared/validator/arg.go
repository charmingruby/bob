package validator

import "fmt"

type Arg struct {
	FieldName     string
	MustHaveState bool
	DefaultState  string
	CurrentState  string
	EmptyState    string
}

func NewArg(fieldName string, mustHaveState bool, defaultState, currentState, emptyState string) Arg {
	return Arg{
		FieldName:     fieldName,
		MustHaveState: mustHaveState,
		DefaultState:  defaultState,
		EmptyState:    emptyState,
		CurrentState:  currentState,
	}
}

func ValidateArgsList(args []Arg) error {
	for _, arg := range args {
		if arg.MustHaveState && arg.CurrentState == arg.EmptyState {
			return fmt.Errorf("missing state for %s argument", arg.FieldName)
		}

		if !arg.MustHaveState {
			arg.CurrentState = arg.DefaultState
		}
	}

	return nil
}
