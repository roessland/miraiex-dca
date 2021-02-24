package miraiex

import (
	"net/http"
	"strings"
)

func maskSecretHeaders(headers http.Header) http.Header {
	maskedHeaders := make(http.Header)
	for k, v := range headers {
		if strings.EqualFold(k, "Miraiex-Access-Key") {
			maskedHeaders[k] = []string{"<masked>"}
		} else {
			maskedHeaders[k] = v
		}
	}
	return maskedHeaders
}