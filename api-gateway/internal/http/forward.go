package http

import (
	"io"
	"net/http"

	"github.com/henrystream/eduflex/api-gateway/internal/client"
)

func forward(w http.ResponseWriter, r *http.Request, baseURL string) {
	target := baseURL + r.URL.Path
	if r.URL.RawQuery != "" {
		target += "?" + r.URL.RawQuery
	}

	resp, err := client.ForwardRequest(r.Method, target, r.Body, r.Header)
	if err != nil {
		http.Error(w, "service unavailable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		w.Header()[k] = v
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
