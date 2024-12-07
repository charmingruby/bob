package opt

const (
	POSTGRES_DATABASE = "postgres"
)

func IsDatabaseOption(value string) bool {
	opts := map[string]bool{
		POSTGRES_DATABASE: true,
	}

	_, ok := opts[value]

	return ok
}
