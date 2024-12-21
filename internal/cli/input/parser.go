package input

import (
	"fmt"
	"strconv"
)

func ParseInteger(arg Arg) (int, error) {
	value, err := strconv.Atoi(arg.Value)
	if err != nil {
		return -1, fmt.Errorf("field '%s' must be an integer", arg.FieldName)
	}

	return value, nil
}

func ParseFloat(arg Arg) (float64, error) {
	value, err := strconv.ParseFloat(arg.Value, 64)
	if err != nil {
		return -1, fmt.Errorf("field '%s' must be a float", arg.FieldName)
	}

	return value, nil
}

func ParseBool(arg Arg) (bool, error) {
	value, err := strconv.ParseBool(arg.Value)
	if err != nil {
		return false, fmt.Errorf("field '%s' must be a boolean", arg.FieldName)
	}

	return value, nil
}
