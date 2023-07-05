package webapp

import (
	"github.com/MrSterdy/ApolloHW/webapp/controller/rest"
	"github.com/MrSterdy/ApolloHW/webapp/controller/view"
	"github.com/MrSterdy/ApolloHW/webapp/controller/view/calendar"
	"github.com/MrSterdy/ApolloHW/webapp/controller/view/class"
	"github.com/MrSterdy/ApolloHW/webapp/controller/view/settings"
	"github.com/MrSterdy/ApolloHW/webapp/net"
	"net/http"
)

func Start() {
	http.Handle("/", net.NewHandler(view.Index))

	http.Handle("/calendar", net.NewHandler(calendar.Index))
	http.Handle("/calendar/timetable", net.NewHandler(calendar.Timetable))

	http.Handle("/class", net.NewHandler(class.Index))

	http.Handle("/settings", net.NewHandler(settings.Index))
	http.Handle("/settings/timetable", net.NewHandler(settings.Timetable))
	http.Handle("/settings/holidays", net.NewHandler(settings.Holidays))
	http.Handle("/settings/subjects", net.NewHandler(settings.Subjects))

	http.Handle("/api/subjects", net.NewHandler(rest.Subjects))
	http.Handle("/api/timetable", net.NewHandler(rest.Timetable))
	http.Handle("/api/class", net.NewHandler(rest.Class))
	http.Handle("/api/holidays", net.NewHandler(rest.Holidays))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("webapp/public/static"))))

	panic(http.ListenAndServe(":8080", nil))
}
