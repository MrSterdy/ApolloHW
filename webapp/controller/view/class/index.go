package class

import (
	"net/http"
	"sort"

	"github.com/MrSterdy/ApolloHW/model/student"
	"github.com/MrSterdy/ApolloHW/webapp/net"
	"github.com/MrSterdy/ApolloHW/webapp/public/class"
)

var Index = net.Controller{
	"GET": func(w http.ResponseWriter, r *http.Request, user int64) int {
		s := student.Find(user)
		if s == nil || !s.Active {
			return http.StatusUnauthorized
		}

		students := student.GetAll()
		if len(students) == 0 {
			class.WriteEmpty(w)

			return http.StatusOK
		}

		var active, inactive []*student.Student
		for _, s := range students {
			if s.Active {
				active = append(active, s)
			} else {
				inactive = append(inactive, s)
			}
		}

		sort.Slice(active, func(i, j int) bool {
			first := active[i]
			second := active[j]

			if first.LastName != second.LastName {
				return first.LastName < second.LastName
			}

			return first.FirstName < second.FirstName
		})

		if r.URL.Query().Has("edit") {
			if s.Role != student.RoleAdmin {
				return http.StatusForbidden
			}

			class.WriteEdit(w, active, inactive)
		} else {
			class.WriteIndex(w, active, s.Role == student.RoleAdmin)
		}

		return http.StatusOK
	},
}
