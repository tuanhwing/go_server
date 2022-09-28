//go:build dev
// +build dev

package environment

const (
	DB_USER                = "temporary"
	DB_PASSWORD            = "CB6ij9cicIOlUvrC"
	DB_HOST                = "cluster0.piciiea.mongodb.net/?retryWrites=true&w=majority"
	DB_NAME                = "Cluster0"
	API_PORT               = 8080
	PROMETHEUS_PUSHGATEWAY = "http://localhost:9091/"
	JWT_SECRET             = "GnpL5gJ3nZD0PP1irp6tXcCp2JxILhwH"
)
