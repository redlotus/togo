package main

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

const (
	HeaderXForwardedFor = "X-FORWARDED-FOR"
	HeaderXRealIP       = "X-REAL-IP"
)

type (
	Render struct{}

	Handler struct {
		Context *ContextHandler
		Logger  *zap.SugaredLogger
		Render  *Render
	}

	// StoreHandler struct {
	// 	Todo *StoreTodo
	// }

	ContextHandler struct {
		Config *Config
		Store  *StoreTodo
	}

	RequestContext struct {
		AccountID string                 `json:"account_id"`
		Platform  string                 `json:"platform"`
		IP        string                 `json:"ip"`
		Data      map[string]interface{} `json:"data"`
	}

	RequestBodyTask struct {
		Rule string `json:"rule"`
	}

	ResponseData struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}
)

func NewHandler(ctx *ContextHandler, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		Context: ctx,
		Logger:  logger,
	}
}

func (h *Handler) ProcessTask(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h.Render.JSON(w, ResponseData{Message: "somehow it works", Code: 0}, http.StatusOK)
	return
}

func (r *Render) JSON(w http.ResponseWriter, res interface{}, httpCode int) {
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}

func responseJSON(w http.ResponseWriter, res interface{}, httpCode int) {
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}

func getIP(r *http.Request) string {
	if ip := r.Header.Get(HeaderXForwardedFor); ip != "" {
		i := strings.IndexAny(ip, ", ")
		if i > 0 {
			return ip[:i]
		}
		return ip
	}

	if ip := r.Header.Get(HeaderXRealIP); ip != "" {
		return ip
	}

	ra, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ra
}

func getContextRequest(r *http.Request) RequestContext {
	var (
		ctxReq RequestContext
	)

	ctxReq.Data, _ = r.Context().Value("data").(map[string]interface{})
	logger.Debugw("request context",
		"func", "getContextRequest",
		"type", "Main",
		"value", ctxReq.Data)

	if ctxReq.Data != nil {
		if uid, ok := ctxReq.Data["uid"]; ok {
			ctxReq.AccountID = uid.(string)
		}

		if platform, ok := ctxReq.Data["platform"]; ok {
			ctxReq.Platform = platform.(string)
		}
	}

	ctxReq.IP = getIP(r)

	return ctxReq
}

func normalizeHeaderField(r *http.Request, h string) string {
	return strings.TrimSpace(strings.ToLower(r.Header.Get(h)))
}
