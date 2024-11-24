package input

type DefaultCommandInput struct {
	Module string
	Name   string
}

func ValidateDefaultCommandInput(module, name string) error {
	args := []arg{
		{
			FieldName:  "module",
			Value:      module,
			IsRequired: true,
		},
		{
			FieldName:  "name",
			Value:      name,
			IsRequired: true,
		},
	}

	if err := validateArgsList(args); err != nil {
		return err
	}

	return nil
}

func ValidateOnlyModuleCommandInput(module string) error {
	arg := arg{
		FieldName:  "module",
		Value:      module,
		IsRequired: true,
	}

	return validateArg(arg)
}
