package config

import (
	"os"
	"path/filepath"
)

const (
//Hostname = "localhost"
)

var (
	//CAFile         = certFile("ca.pem")
	ServerCertFile = certFile("server.pem")

	//ServerKeyFile  = certFile("server-key.pem")
)

func certFile(filename string) string {
	if dir := os.Getenv("CERTS_DIR"); dir != "" {
		return filepath.Join(dir, filename)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homeDir, ".axone-api", filename)
}
