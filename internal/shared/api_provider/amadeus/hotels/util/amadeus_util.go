package util 

import (
	"fmt"
	"net/url"
)

func BuildURL(baseURL, endpoint string, params map[string]string) string {
	u, _ := url.Parse(fmt.Sprintf("%s%s", baseURL, endpoint))
	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
