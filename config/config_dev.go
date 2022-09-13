//go:build dev
// +build dev

package config

const (
	DB_USER                = "temporary"
	DB_PASSWORD            = "CB6ij9cicIOlUvrC"
	DB_HOST                = "cluster0.piciiea.mongodb.net/?retryWrites=true&w=majority"
	DB_NAME                = "Cluster0"
	API_PORT               = 8080
	PROMETHEUS_PUSHGATEWAY = "http://localhost:9091/"
)
