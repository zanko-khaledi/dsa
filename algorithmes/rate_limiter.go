package algorithmes

import (
	"net/http"
	"sync"
	"time"
)

type RatelimterConfig struct {
	Limit    int
	Window   time.Duration
	Duration int
}

type Client struct {
	Requests []time.Time
}

var (
	clients = make(map[string]*Client)
	mu      sync.Mutex
)

func (ratelimter *RatelimterConfig) Ratelimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := r.RemoteAddr // for production use r.Header.Get("X-Forwared-For") or X-Real-Ip

		mu.Lock()

		client, exists := clients[ip]

		if !exists {
			client = &Client{}
			clients[ip] = client
		}

		now := time.Now()
		cutoff := now.Add(-time.Duration(ratelimter.Duration) * ratelimter.Window)

		var valied []time.Time

		for _, t := range client.Requests {
			if t.After(cutoff) {
				valied = append(valied, t)
			}
		}

		client.Requests = valied

		if len(client.Requests) >= ratelimter.Limit {
			mu.Unlock()

			http.Error(w, "Too many requests!", 429)
			return
		}

		client.Requests = append(client.Requests, now)

		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}
