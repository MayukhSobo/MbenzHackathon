package utils

import (
	"fmt"
	"os"
)

// URIBuilder func for building URL connections
func URIBuilder(kind string) (string, error) {
	var url string
	switch kind {
	case "fiber":
		url = fmt.Sprintf("%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		return "", fmt.Errorf("connection type '%s' is not supported", kind)
	}

	return url, nil
}
