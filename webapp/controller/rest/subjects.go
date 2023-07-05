package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MrSterdy/ApolloHW/config"
	"github.com/MrSterdy/ApolloHW/model/student"
	"github.com/MrSterdy/ApolloHW/model/timetable"
	"github.com/MrSterdy/ApolloHW/util"
	"github.com/MrSterdy/ApolloHW/webapp/net"
)

var Subjects = net.Controller{
	"PUT": func(w http.ResponseWriter, r *http.Request, user int64) int {
		s := student.Find(user)
		if s == nil || s.Role != student.RoleAdmin {
			return http.StatusForbidden
		}

		var subjects timetable.Subjects
		if err := json.NewDecoder(r.Body).Decode(&subjects); err != nil {
			return http.StatusBadRequest
		}

		if len(subjects) == 0 {
			timetable.DeleteSubjects()
		} else {
			strs := make(map[string]bool, len(subjects))
			for _, ss := range subjects {
				lowerName := strings.ToLower(ss.Name)

				if strs[lowerName] {
					return http.StatusBadRequest
				}

				strs[lowerName] = true
			}

			timetable.SaveSubjects(subjects)
		}

		msgs := config.File.Messages
		util.SendNotification(strings.Replace(msgs.Notifications.Logs.Subjects, msgs.Templates.User, s.LastName+" "+s.FirstName, 1), student.RoleAdmin)

		return http.StatusOK
	},
}
