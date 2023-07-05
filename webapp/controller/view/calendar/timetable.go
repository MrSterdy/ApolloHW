package calendar

import (
	"net/http"

	"github.com/MrSterdy/ApolloHW/model/holiday"
	"github.com/MrSterdy/ApolloHW/model/student"
	"github.com/MrSterdy/ApolloHW/model/timetable"
	"github.com/MrSterdy/ApolloHW/util"
	"github.com/MrSterdy/ApolloHW/webapp/net"
	template "github.com/MrSterdy/ApolloHW/webapp/public/calendar/timetable"
)

var Timetable = net.Controller{
	"GET": func(w http.ResponseWriter, r *http.Request, user int64) int {
		s := student.Find(user)
		if s == nil || !s.Active {
			return http.StatusUnauthorized
		}

		query := r.URL.Query()

		date, err := util.ParseDate(query.Get("date"))
		if err != nil {
			return http.StatusBadRequest
		}

		dt := timetable.FindByDate(date)

		if query.Has("edit") {
			if s.Role == student.RoleStudent {
				return http.StatusForbidden
			}

			weekday := date.Weekday()
			wt := timetable.FindByWeekday(weekday)
			if wt == nil {
				wt = new(timetable.Weekday)
			}

			if dt == nil {
				dt = new(timetable.Date)
				dt.Base = wt.Base
				dt.Date = date

				template.WriteEdit(w, dt, wt, dt.Subjects)
			} else {
				template.WriteEdit(w, dt, wt, dt.Subjects)
			}
		} else {
			isHoliday := false
			for _, h := range holiday.GetAll() {
				if h.IntersectsWith(date) {
					isHoliday = true

					break
				}
			}

			if dt == nil {
				wt := timetable.FindByWeekday(date.Weekday())

				if wt == nil || isHoliday || len(wt.Subjects) == 0 {
					template.WriteRest(w, s.Role != student.RoleStudent)
				} else {
					dt = new(timetable.Date)
					dt.Date = date
					dt.Base = wt.Base

					template.WriteIndex(w, dt, s.Role != student.RoleStudent)
				}
			} else if dt.Note == "" && (isHoliday || len(dt.Subjects) == 0) {
				template.WriteRest(w, s.Role != student.RoleStudent)
			} else {
				template.WriteIndex(w, dt, s.Role != student.RoleStudent)
			}
		}

		return http.StatusOK
	},
}
