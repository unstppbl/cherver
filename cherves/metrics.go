package cherves

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MetricsExporter interface {
	WritePrometheus(w io.Writer)
}

type MetricsResource struct {
	set MetricsExporter
}

func NewMetricResource(set MetricsExporter) *MetricsResource {
	return &MetricsResource{set: set}
}

func (m *MetricsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", m.metrics)

	return r
}

func (m *MetricsResource) Path() string {
	return "/metrics"
}

func (m *MetricsResource) metrics(w http.ResponseWriter, _ *http.Request) {
	m.set.WritePrometheus(w)
}
