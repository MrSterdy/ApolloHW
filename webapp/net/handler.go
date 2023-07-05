package net

import (
	"net/http"

	"github.com/MrSterdy/ApolloHW/util"
	"github.com/MrSterdy/ApolloHW/webapp/public"
)

type Controller map[string]func(http.ResponseWriter, *http.Request, int64) int

func (c Controller) ServeMethod(w http.ResponseWriter, r *http.Request) int {
	if m := c[r.Method]; m != nil {
		var user int64

		if r.URL.RequestURI() != "/" {
			cookie, err := r.Cookie("apollo_initdata")
			if err != nil {
				return http.StatusBadRequest
			}

			id, ok := util.ParseInitData(cookie.Value)
			if !ok {
				return http.StatusBadRequest
			}

			user = id
		}

		return m(w, r, user)
	}

	return http.StatusMethodNotAllowed
}

type Handler struct {
	Controller Controller
}

func NewHandler(controller Controller) *Handler {
	return &Handler{controller}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if status := h.Controller.ServeMethod(w, r); status != http.StatusOK {
		w.WriteHeader(status)

		public.WriteError(w, http.StatusText(status))
	}
}
