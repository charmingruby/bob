package molecule

const (
	REST_PREFIX = "molecule/rest"

	SERVICE_PREFIX = "molecule/service"
)

var (
	REST_HANDLER_REGISTRY_TEMPLATE = REST_PREFIX + "/handler_registry"
	REST_EXCHANGE_TEMPLATE         = REST_PREFIX + "/exchange"
	REST_RESPONSE_HELPER_TEMPLATE  = REST_PREFIX + "/response_helper"
	REST_REQUEST_HELPER_TEMPLATE   = REST_PREFIX + "/request_helper"
	REST_HANDLER_TEMPLATE          = REST_PREFIX + "/handler"
	REST_BASE_SERVER_MIDDLEWARE    = REST_PREFIX + "/base_server_middleware"
	REST_SERVER                    = REST_PREFIX + "/server"

	SERVICE_REGISTRY_TEMPLATE                 = SERVICE_PREFIX + "/registry"
	SERVICE_REGISTRY_WITH_REPOSITORY_TEMPLATE = SERVICE_PREFIX + "/registry_with_repository"
)
