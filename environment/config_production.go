//go:build production
// +build production

package environment

const (
	DB_USER                = "clean_architecture_go_v2"
	DB_PASSWORD            = "clean_architecture_go_v2"
	DB_DATABASE            = "clean_architecture_go_v2"
	DB_NAME                = "Cluster0"
	DB_HOST                = "127.0.0.1"
	API_PORT               = 8080
	PROMETHEUS_PUSHGATEWAY = "http://localhost:9091/"
	ENV_FILE               = "env.dev"
)
