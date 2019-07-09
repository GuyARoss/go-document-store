package utilities

import (
	"net/http"
)

func GetParam(req *http.Request, value string) string {
	if len(req.URL.Query()[value]) > 0 {
		return req.URL.Query()[value][0]
	}

	return ""
}