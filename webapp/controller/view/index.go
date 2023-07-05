package view

import (
	"net/http"

	"github.com/MrSterdy/ApolloHW/model/student"
	"github.com/MrSterdy/ApolloHW/webapp/net"
	"github.com/MrSterdy/ApolloHW/webapp/public"
)

var Index = net.Controller{
	"GET": func(w http.ResponseWriter, r *http.Request, user int64) int {
		if r.URL.Query().Has("auth") {
			s := student.Find(user)
			if s == nil {
				public.WriteRegistration(w)
			} else if !s.Active {
				public.WriteWait(w)
			} else {
				public.WriteMenu(w, s)
			}
		} else {
			public.WriteIndex(w)
		}

		return http.StatusOK
	},
}
