package opt

const (
	POSTGRES_DATABASE_OPTION = "postgres"
)

func IsDatabaseOption(value string) bool {
	opts := map[string]bool{
		POSTGRES_DATABASE_OPTION: true,
	}

	_, ok := opts[value]

	return ok
}
