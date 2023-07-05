package settings

import (
	"net/http"

	"github.com/MrSterdy/ApolloHW/model/student"
	"github.com/MrSterdy/ApolloHW/model/timetable"
	"github.com/MrSterdy/ApolloHW/webapp/net"
	"github.com/MrSterdy/ApolloHW/webapp/public/settings"
)

var Subjects = net.Controller{
	"GET": func(w http.ResponseWriter, r *http.Request, user int64) int {
		if s := student.Find(user); s == nil || s.Role != student.RoleAdmin {
			return http.StatusForbidden
		}

		settings.WriteSubjects(w, timetable.GetSubjects())

		return http.StatusOK
	},
}
