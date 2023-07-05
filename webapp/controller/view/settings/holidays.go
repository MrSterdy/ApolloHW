package settings

import (
	"net/http"

	"github.com/MrSterdy/ApolloHW/model/holiday"
	"github.com/MrSterdy/ApolloHW/model/student"
	"github.com/MrSterdy/ApolloHW/webapp/net"
	"github.com/MrSterdy/ApolloHW/webapp/public/settings"
)

var Holidays = net.Controller{
	"GET": func(w http.ResponseWriter, r *http.Request, user int64) int {
		if s := student.Find(user); s == nil || s.Role != student.RoleAdmin {
			return http.StatusForbidden
		}

		settings.WriteHolidays(w, holiday.GetAll())

		return http.StatusOK
	},
}
