package proxy

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type proxy struct {

	target *url.URL
}

func (p *proxy) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(struct {
		Url string
		Host string
	}{
		Url: p.target.String(),
		Host: r.Host,
	}  )

	//w.Write([]byte(p.target.String()))
}

func newDryrunProxy(target *url.URL , _ http.RoundTripper , _ time.Duration) http.Handler {

	return &proxy{ target: target}
}