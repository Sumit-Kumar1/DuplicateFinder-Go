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
	tmpl := template.Must(template.ParseGlob("views/*.html"))
	return &Handler{
		templ:   tmpl,
		Log:     log,
		Service: s,
	}
}

func (h *Handler) RenderPage(w http.ResponseWriter, _ *http.Request) {
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

	err = h.templ.ExecuteTemplate(w, "sysinfo", *info)
	if err != nil {
		h.Log.Error("error in template execution", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) CurrentUsage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	usage, err := h.Service.CurrentUsage()
	if err != nil {
		h.Log.Error("error while getting current usage info: ", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if usage == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = h.templ.ExecuteTemplate(w, "usages", usage)
	if err != nil {
		h.Log.Error("error in template execution", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DuplicateFind(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.Log.Error("method is not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	queryParams := r.URL.Query()
	device := queryParams.Get("device")

	h.Log.Info("got duplicate finding request for", "device", device)

	duplicates, err := h.Service.Duplicate(device)
	if err != nil {
		h.Log.Error("error while duplicate finding disk info: ", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = h.templ.ExecuteTemplate(w, "duplicates", duplicates)
	if err != nil {
		h.Log.Error("error in template execution", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
