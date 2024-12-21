package input

import (
	"fmt"
)

type ArgType int

const (
	StringType ArgType = iota
	IntType
	FloatType
	BoolType
)

type Arg struct {
	FieldName        string
	Value            string
	IsRequired       bool
	Type             ArgType
	CustomValidation func(string) error
}

type ValidatorFunc func(Arg) error

var validators = map[ArgType]ValidatorFunc{
	StringType: validateString,
	IntType:    validateInteger,
	FloatType:  validateFloat,
	BoolType:   validateBool,
}

func Validate(args []Arg) error {
	for _, arg := range args {
		if err := validateArg(arg); err != nil {
			return err
		}
	}
	return nil
}

func validateArg(arg Arg) error {
	if arg.IsRequired && isArgEmpty(arg.Value) {
		return fmt.Errorf("field '%s' is required", arg.FieldName)
	}

	if validator, ok := validators[arg.Type]; ok {
		if err := validator(arg); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no validator found for field '%s'", arg.FieldName)
	}

	if arg.CustomValidation != nil {
		if err := arg.CustomValidation(arg.Value); err != nil {
			return fmt.Errorf("field '%s' validation failed: %v", arg.FieldName, err)
		}
	}

	return nil
}

func validateString(arg Arg) error {
	// adds some validation, if needed

	return nil
}

func validateInteger(arg Arg) error {
	if _, err := ParseInteger(arg); err != nil {
		return err
	}

	return nil
}

func validateBool(arg Arg) error {
	if _, err := ParseBool(arg); err != nil {
		return err
	}

	return nil
}

func validateFloat(arg Arg) error {
	if _, err := ParseFloat(arg); err != nil {
		return err
	}

	return nil
}

func isArgEmpty(v string) bool {
	return v == ""
}
