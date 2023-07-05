package settings

import (
	"net/http"
	"strconv"
	"time"

	"github.com/MrSterdy/ApolloHW/model/student"
	"github.com/MrSterdy/ApolloHW/model/timetable"
	"github.com/MrSterdy/ApolloHW/webapp/net"
	"github.com/MrSterdy/ApolloHW/webapp/public/settings"
)

var Timetable = net.Controller{
	"GET": func(w http.ResponseWriter, r *http.Request, user int64) int {
		if s := student.Find(user); s == nil || s.Role != student.RoleAdmin {
			return http.StatusForbidden
		}

		if day, err := strconv.Atoi(r.URL.Query().Get("day")); err == nil {
			wt := timetable.FindByWeekday(time.Weekday(day))
			if wt == nil {
				if day >= int(time.Sunday) && day <= int(time.Saturday) {
					wt = new(timetable.Weekday)
					wt.Day = time.Weekday(day)
				} else {
					return http.StatusNotFound
				}
			}

			settings.WriteTimetable(w, wt, timetable.GetSubjects())
		} else {
			return http.StatusBadRequest
		}

		return http.StatusOK
	},
}
