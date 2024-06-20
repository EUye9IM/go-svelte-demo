package main

import (
	"demo/frontend"

	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

const addr = ":8001"

func main() {
	http.Handle("/", http.FileServer(http.FS(frontend.F)))
	http.HandleFunc("/api/demo", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			return
		}
		slog.Debug("api/demo", "req", r)

		dec := json.NewDecoder(r.Body)
		req := struct {
			Name string
		}{}

		if err := dec.Decode(&req); err != nil {
			slog.Warn("HandleFunc ReadAll Fail",
				"err", err)
		}
		if raw, err := json.Marshal(map[string]any{"msg": fmt.Sprintf("Hello, %v.", req.Name)}); err != nil {
			w.Write([]byte(err.Error()))
			slog.Warn("HandleFunc Marshal Fail",
				"err", err)
		} else {
			w.Write(raw)
		}
	})
	slog.Info("start serve",
		"addr", "http://localhost"+addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("ListenAndServe Fail",
			"err", err)
	}
}
