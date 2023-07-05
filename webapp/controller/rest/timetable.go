package rest

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/MrSterdy/ApolloHW/config"
	"github.com/MrSterdy/ApolloHW/model/holiday"
	"github.com/MrSterdy/ApolloHW/model/student"
	"github.com/MrSterdy/ApolloHW/model/timetable"
	"github.com/MrSterdy/ApolloHW/util"
	"github.com/MrSterdy/ApolloHW/webapp/net"
)

var Timetable = net.Controller{
	"PUT": func(w http.ResponseWriter, r *http.Request, user int64) int {
		s := student.Find(user)
		if s.Role == student.RoleStudent {
			return http.StatusForbidden
		}

		query := r.URL.Query()

		msgs := config.File.Messages

		if t := query.Get("type"); t == "date" {
			newTt := new(timetable.Date)
			if err := json.NewDecoder(r.Body).Decode(newTt); err != nil {
				return http.StatusBadRequest
			}

			var oldTt timetable.Timetable
			if dt := timetable.FindByDate(newTt.Date); dt != nil {
				oldTt = dt
			} else if wt := timetable.FindByWeekday(newTt.Date.Weekday()); wt != nil {
				oldTt = wt
			} else {
				oldTt = new(timetable.Weekday)
			}

			added, removed := oldTt.GetSubjects().Diff(newTt.Subjects)
			holidays := holiday.GetAll()
			processAdded(added, newTt.Date, holidays)
			processRemoved(removed, newTt.Date, holidays)

			timetable.SaveByDate(newTt)

			formatted := newTt.Date.Format(msgs.Templates.Date.Layout)
			util.SendNotification(strings.Replace(
				strings.Replace(msgs.Notifications.Logs.Timetable, msgs.Templates.Date.Template, formatted, 1),
				msgs.Templates.User,
				s.LastName+" "+s.FirstName,
				1,
			), student.RoleAdmin)
			if query.Has("notification") {
				util.SendNotification(strings.Replace(msgs.Notifications.Timetable, msgs.Templates.Date.Template, formatted, 1), student.RoleStudent)
			}
		} else if t == "day" {
			if s.Role != student.RoleAdmin {
				return http.StatusForbidden
			}

			tt := new(timetable.Weekday)
			if err := json.NewDecoder(r.Body).Decode(tt); err != nil || tt.Day < time.Sunday || tt.Day > time.Saturday {
				return http.StatusBadRequest
			}

			timetable.SaveByWeekday(tt)

			util.SendNotification(strings.Replace(
				strings.Replace(msgs.Notifications.Logs.Timetable, msgs.Templates.Date.Template, msgs.Weekdays[tt.Day], 1),
				msgs.Templates.User,
				s.LastName+" "+s.FirstName,
				1,
			), student.RoleAdmin)
		} else {
			return http.StatusBadRequest
		}

		return http.StatusOK
	},
}

func findNextTimetableConsideringHolidays(subject string, startDate *util.Date, holidays []holiday.Holiday) (timetable.Timetable, *util.Date) {
	tt, date := timetable.FindNextTimetable(subject, startDate)
	if tt == nil {
		return nil, nil
	}

	for _, h := range holidays {
		if h.IntersectsWith(date) {
			if h.EndDate != nil {
				return findNextTimetableConsideringHolidays(subject, h.EndDate, holidays)
			}

			return findNextTimetableConsideringHolidays(subject, date, holidays)
		}
	}

	return tt, date
}

func processAdded(subjects timetable.Subjects, date *util.Date, holidays []holiday.Holiday) {
	for _, s := range subjects {
		if s.Name == "" {
			continue
		}

		if s.Homework == nil {
			nextTt, _ := findNextTimetableConsideringHolidays(s.Name, date, holidays)
			nextDt, ok := nextTt.(*timetable.Date)
			if !ok {
				continue
			}

			for _, nextSubject := range nextDt.Subjects {
				if nextSubject.Name == s.Name && nextSubject.Homework != nil {
					s.Homework = nextSubject.Homework
					nextSubject.Homework = nil

					timetable.SaveByDate(nextDt)

					break
				}
			}
		}
	}
}

func processRemoved(subjects timetable.Subjects, date *util.Date, holidays []holiday.Holiday) {
	for _, s := range subjects {
		if s.Name == "" {
			continue
		}

		if s.Homework != nil {
			nextTt, nextDate := findNextTimetableConsideringHolidays(s.Name, date, holidays)
			nextDt, ok := nextTt.(*timetable.Date)
			if !ok {
				nextDt = &timetable.Date{
					Base: timetable.Base{
						Offset:   nextTt.GetOffset(),
						Subjects: nextTt.GetSubjects(),
					},
					Date: nextDate,
				}
			}

			var sub *timetable.Subject
			for _, nextSubject := range nextDt.Subjects {
				if nextSubject.Name == s.Name {
					sub = nextSubject

					if sub.Homework != nil {
						break
					}
				}
			}

			if sub.Homework == nil {
				sub.Homework = s.Homework
			} else {
				homework := new(timetable.Homework)

				if s.Homework.Text == "" {
					homework.Text = sub.Homework.Text
				} else if sub.Homework.Text == "" {
					homework.Text = s.Homework.Text
				} else {
					homework.Text = s.Homework.Text + " + " + sub.Homework.Text
				}

				homework.Images = append(s.Homework.Images, sub.Homework.Images...)

				sub.Homework = homework
			}

			timetable.SaveByDate(nextDt)
		}
	}
}
