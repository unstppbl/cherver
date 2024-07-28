package cherves

import (
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Profiling struct {}

func (p Profiling) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", p.prof)

	return r
}

func (p Profiling) Path() string {
	return "/debug/pprof"
}

func (p Profiling) prof(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/debug/pprof/":
		pprof.Index(w, r)
	case "/debug/pprof/cmdline":
		pprof.Cmdline(w, r)
	case "/debug/pprof/profile":
		pprof.Profile(w, r)
	case "/debug/pprof/symbol":
		pprof.Symbol(w, r)
	case "/debug/pprof/trace":
		pprof.Trace(w, r)
	case "debug/pprof/allocs":
		pprof.Handler("allocs").ServeHTTP(w, r)
	case "/debug/pprof/block":
		pprof.Handler("block").ServeHTTP(w, r)
	case "/debug/pprof/goroutine":
		pprof.Handler("goroutine").ServeHTTP(w, r)
	case "/debug/pprof/heap":
		pprof.Handler("heap").ServeHTTP(w, r)
	case "/debug/pprof/mutex":
		pprof.Handler("mutex").ServeHTTP(w, r)
	case "/debug/pprof/threadcreate":
		pprof.Handler("threadcreate").ServeHTTP(w, r)
	default:
		if strings.HasSuffix(path, "/") {
			path = strings.TrimRight(path, "/")
		} else {
			path = "/debug/pprof"
		}

		url := *r.URL
		url.Path = path
		url.RawPath = path
		http.Redirect(w, r, url.String(), http.StatusFound)
	}
}

