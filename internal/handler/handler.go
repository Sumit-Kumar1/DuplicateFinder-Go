package handler

import (
	"html/template"
	"log/slog"
	"net/http"
)

type Handler struct {
	templ   *template.Template
	Log     *slog.Logger
	Service Servicer
}

func New(log *slog.Logger, s Servicer) *Handler {
	tmpl := template.Must(template.ParseGlob("views/index.html"))
	return &Handler{
		templ:   tmpl,
		Log:     log,
		Service: s,
	}
}

func (h *Handler) RenderPage(w http.ResponseWriter, r *http.Request) {
	if err := h.templ.ExecuteTemplate(w, "index", nil); err != nil {
		h.Log.Error("error in template:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// SystemInfo returns the information about a system and usage part
func (h *Handler) SystemInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	info, err := h.Service.SysInfo()
	if err != nil {
		h.Log.Error("error while gettting system info: ", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if info == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = h.templ.ExecuteTemplate(w, "sysInfo", map[string]any{
		"Data": *info,
	})
	if err != nil {
		h.Log.Error("error in template execution", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *Handler) CurrentUsage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}
