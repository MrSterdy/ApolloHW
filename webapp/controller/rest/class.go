package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MrSterdy/ApolloHW/config"
	"github.com/MrSterdy/ApolloHW/model/student"
	"github.com/MrSterdy/ApolloHW/util"
	"github.com/MrSterdy/ApolloHW/webapp/net"
)

var Class = net.Controller{
	"PUT": func(w http.ResponseWriter, r *http.Request, user int64) int {
		s := student.Find(user)
		if s == nil || s.Role != student.RoleAdmin {
			return http.StatusForbidden
		}

		var students []*student.Student
		if err := json.NewDecoder(r.Body).Decode(&students); err != nil {
			return http.StatusBadRequest
		}

		if len(students) == 0 {
			student.DeleteAll()
		} else {
			for _, st := range students {
				if !st.Role.IsValid() {
					return http.StatusBadRequest
				}
			}

			student.SaveAll(students)
		}

		msgs := config.File.Messages
		util.SendNotification(strings.Replace(msgs.Notifications.Logs.Class, msgs.Templates.User, s.LastName+" "+s.FirstName, 1), student.RoleAdmin)

		return http.StatusOK
	},
	"POST": func(w http.ResponseWriter, r *http.Request, user int64) int {
		var s *student.Student

		if student.IsEmpty() {
			s = new(student.Student)
			if err := json.NewDecoder(r.Body).Decode(s); err != nil {
				return http.StatusBadRequest
			}

			s.Role = student.RoleAdmin
			s.Active = true
		} else if s = student.Find(user); s != nil {
			return http.StatusConflict
		} else {
			s = new(student.Student)
			if err := json.NewDecoder(r.Body).Decode(s); err != nil || s.Role != student.RoleStudent || s.Active {
				return http.StatusBadRequest
			}
		}

		s.ID = user

		student.Save(s)

		if !s.Active {
			util.SendNotification(config.File.Messages.Notifications.Logs.Application, student.RoleAdmin)
		}

		return http.StatusOK
	},
}
