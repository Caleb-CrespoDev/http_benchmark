package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	pool *pgxpool.Pool

	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "HTTP request duration in seconds",
	}, []string{"method", "route", "status_code"})
)

type item struct {
	ID        int64     `json:"id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

func withMetrics(route string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: 200}
		h(sw, r)
		httpRequestDuration.WithLabelValues(r.Method, route, strconv.Itoa(sw.status)).Observe(time.Since(start).Seconds())
	}
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func listItemsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := pool.Query(r.Context(), "SELECT id, value, created_at FROM items ORDER BY id DESC LIMIT 100")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	items := []item{}
	for rows.Next() {
		var it item
		if err := rows.Scan(&it.ID, &it.Value, &it.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, it)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Value string `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var it item
	err := pool.QueryRow(
		r.Context(),
		"INSERT INTO items (value) VALUES ($1) RETURNING id, value, created_at",
		body.Value,
	).Scan(&it.ID, &it.Value, &it.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(it)
}

func updateItemHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Value string `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var it item
	err := pool.QueryRow(
		r.Context(),
		"UPDATE items SET value = $1 WHERE id = $2 RETURNING id, value, created_at",
		body.Value, r.PathValue("id"),
	).Scan(&it.ID, &it.Value, &it.CreatedAt)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(it)
}

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	tag, err := pool.Exec(r.Context(), "DELETE FROM items WHERE id = $1", r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if tag.RowsAffected() == 0 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	_, err := pool.Exec(r.Context(), "TRUNCATE TABLE items RESTART IDENTITY")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "reset"})
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	dsn := "postgres://" + envOr("PGUSER", "bench") + ":" + os.Getenv("PGPASSWORD") +
		"@" + envOr("PGHOST", "localhost") + ":" + envOr("PGPORT", "5432") +
		"/" + envOr("PGDATABASE", "bench") + "?sslmode=disable"

	var err error
	pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer pool.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", withMetrics("/healthz", healthzHandler))
	mux.HandleFunc("GET /items", withMetrics("/items", listItemsHandler))
	mux.HandleFunc("POST /items", withMetrics("/items", createItemHandler))
	mux.HandleFunc("PUT /items/{id}", withMetrics("/items/{id}", updateItemHandler))
	mux.HandleFunc("DELETE /items/{id}", withMetrics("/items/{id}", deleteItemHandler))
	mux.HandleFunc("POST /reset", withMetrics("/reset", resetHandler))
	mux.Handle("GET /metrics", promhttp.Handler())

	port := envOr("PORT", "4000")
	log.Printf("go-http benchmark app listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
