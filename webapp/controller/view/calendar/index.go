package calendar

import (
	"net/http"

	"github.com/MrSterdy/ApolloHW/model/student"
	"github.com/MrSterdy/ApolloHW/webapp/net"
	"github.com/MrSterdy/ApolloHW/webapp/public/calendar"
)

var Index = net.Controller{
	"GET": func(w http.ResponseWriter, r *http.Request, user int64) int {
		if s := student.Find(user); s == nil || !s.Active {
			return http.StatusUnauthorized
		}

		calendar.WriteIndex(w)

		return http.StatusOK
	},
}
