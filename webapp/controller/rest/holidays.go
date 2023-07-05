package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MrSterdy/ApolloHW/config"
	"github.com/MrSterdy/ApolloHW/model/holiday"
	"github.com/MrSterdy/ApolloHW/model/student"
	"github.com/MrSterdy/ApolloHW/util"
	"github.com/MrSterdy/ApolloHW/webapp/net"
)

var Holidays = net.Controller{
	"PUT": func(w http.ResponseWriter, r *http.Request, user int64) int {
		s := student.Find(user)
		if s == nil || s.Role != student.RoleAdmin {
			return http.StatusForbidden
		}

		var holidays []holiday.Holiday
		if err := json.NewDecoder(r.Body).Decode(&holidays); err != nil {
			return http.StatusBadRequest
		}

		if len(holidays) == 0 {
			holiday.DeleteAll()
		} else {
			holiday.SaveAll(holidays)
		}

		msgs := config.File.Messages
		util.SendNotification(strings.Replace(msgs.Notifications.Logs.Holidays, msgs.Templates.User, s.LastName+" "+s.FirstName, 1), student.RoleAdmin)

		return http.StatusOK
	},
}
