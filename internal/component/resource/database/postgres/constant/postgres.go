package constant

const (
	MIGRATIONS_DIR = "/db/migration"

	POSTGRES_PACKAGE            = "postgres"
	POSTGRES_REPOSITORY_PACKAGE = POSTGRES_PACKAGE + "_repository"

	POSTGRES_PREFIX = "resource/postgres"

	POSTGRES_CONNECTION_TEMPLATE = POSTGRES_PREFIX + "/connection"
	POSTGRES_REPOSITORY_TEMPLATE = POSTGRES_PREFIX + "/repository"
)
