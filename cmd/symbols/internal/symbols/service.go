// Package symbols implements the symbol search service.
package symbols

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sourcegraph/sourcegraph/cmd/symbols/internal/pkg/ctags"
	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/diskcache"
	"github.com/sourcegraph/sourcegraph/internal/gitserver"
)

// Service is the symbols service.
type Service struct { /* all structs must go */ }

// Start must be called before any requests are handled.
func (s *Service) Start() error {
	if err := s.startParsers(); err != nil {
		return err
	}

	if s.MaxConcurrentFetchTar == 0 {
		s.MaxConcurrentFetchTar = 15
	}
	s.fetchSem = make(chan int, s.MaxConcurrentFetchTar)

	s.cache = &diskcache.Store{
		Dir:               s.Path,
		Component:         "symbols",
		BackgroundTimeout: 20 * time.Minute,
	}
	go s.watchAndEvict()

	return nil
}

// Handler returns the http.Handler that should be used to serve requests.
func (s *Service) Handler() http.Handler {
	if s.parsers == nil {
		panic("must call StartParserPool first")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/search", s.handleSearch)
	mux.HandleFunc("/healthz", s.handleHealthCheck)

	return mux
}

func (s *Service) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("Ok"))
	if err != nil {
		log.Printf("failed to write response to health check, err: %s", err)
	}
}

// watchAndEvict is a loop which periodically checks the size of the cache and
// evicts/deletes items if the store gets too large.
func (s *Service) watchAndEvict() {
	if s.MaxCacheSizeBytes == 0 {
		return
	}

	for {
		time.Sleep(10 * time.Second)
		stats, err := s.cache.Evict(s.MaxCacheSizeBytes)
		if err != nil {
			log.Printf("failed to Evict: %s", err)
			continue
		}
		cacheSizeBytes.Set(float64(stats.CacheSize))
		evictions.Add(float64(stats.Evicted))
	}
}

var (
	cacheSizeBytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "symbols_store_cache_size_bytes",
		Help: "The total size of items in the on disk cache.",
	})
	evictions = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "symbols_store_evictions",
		Help: "The total number of items evicted from the cache.",
	})
)

func init() {
	prometheus.MustRegister(cacheSizeBytes)
	prometheus.MustRegister(evictions)
}
