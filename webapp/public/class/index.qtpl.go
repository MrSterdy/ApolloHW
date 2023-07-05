// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line webapp/public/class/index.qtpl:2
package class

//line webapp/public/class/index.qtpl:2
import "github.com/MrSterdy/ApolloHW/model/student"

//line webapp/public/class/index.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line webapp/public/class/index.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line webapp/public/class/index.qtpl:4
func StreamIndex(qw422016 *qt422016.Writer, students []*student.Student, edit bool) {
//line webapp/public/class/index.qtpl:4
	qw422016.N().S(`<main id="swup" class="animated transition-fade" no-return namespace="class"`)
//line webapp/public/class/index.qtpl:5
	if edit {
//line webapp/public/class/index.qtpl:5
		qw422016.N().S(`main-button="РЕДАКТИРОВАТЬ"`)
//line webapp/public/class/index.qtpl:5
	}
//line webapp/public/class/index.qtpl:5
	qw422016.N().S(`><section class="column menu"><ul class="column">`)
//line webapp/public/class/index.qtpl:8
	for _, st := range students {
//line webapp/public/class/index.qtpl:8
		qw422016.N().S(`<li><div class="option-button"><div class="center prefix"><svg class="icon" viewBox="0 0 24 24"><path d="M12,12A6,6,0,1,0,6,6,6.006,6.006,0,0,0,12,12ZM12,2A4,4,0,1,1,8,6,4,4,0,0,1,12,2Z"/><path d="M12,14a9.01,9.01,0,0,0-9,9,1,1,0,0,0,2,0,7,7,0,0,1,14,0,1,1,0,0,0,2,0A9.01,9.01,0,0,0,12,14Z"/></svg></div><div class="center row label"><span>`)
//line webapp/public/class/index.qtpl:19
		qw422016.E().S(st.LastName + " " + st.FirstName)
//line webapp/public/class/index.qtpl:19
		qw422016.N().S(`</span>`)
//line webapp/public/class/index.qtpl:20
		if st.Role == student.RoleRedactor {
//line webapp/public/class/index.qtpl:20
			qw422016.N().S(`<svg class="medium-icon" viewBox="0 0 24 24"><path d="m17.994 2.286a9 9 0 0 0 -14.919 5.536 8.938 8.938 0 0 0 2.793 7.761 6.263 6.263 0 0 1 2.132 4.566v.161a3.694 3.694 0 0 0 3.69 3.69h.62a3.694 3.694 0 0 0 3.69-3.69v-.549a5.323 5.323 0 0 1 1.932-4 8.994 8.994 0 0 0 .062-13.477zm-5.684 19.714h-.62a1.692 1.692 0 0 1 -1.69-1.69s-.007-.26-.008-.31h4.008v.31a1.692 1.692 0 0 1 -1.69 1.69zm4.3-7.741a7.667 7.667 0 0 0 -2.364 3.741h-1.246v-7.184a3 3 0 0 0 2-2.816 1 1 0 0 0 -2 0 1 1 0 0 1 -2 0 1 1 0 0 0 -2 0 3 3 0 0 0 2 2.816v7.184h-1.322a8.634 8.634 0 0 0 -2.448-3.881 7 7 0 0 1 3.951-12.073 7.452 7.452 0 0 1 .828-.046 6.921 6.921 0 0 1 4.652 1.778 6.993 6.993 0 0 1 -.048 10.481z"/></svg>`)
//line webapp/public/class/index.qtpl:24
		} else if st.Role == student.RoleAdmin {
//line webapp/public/class/index.qtpl:24
			qw422016.N().S(`<svg class="medium-icon" viewBox="0 0 24 24"><path d="m18 9.064a3.049 3.049 0 0 0 -.9-2.164 3.139 3.139 0 0 0 -4.334 0l-11.866 11.869a3.064 3.064 0 0 0 4.33 4.331l11.87-11.869a3.047 3.047 0 0 0 .9-2.167zm-14.184 12.624a1.087 1.087 0 0 1 -1.5 0 1.062 1.062 0 0 1 0-1.5l7.769-7.77 1.505 1.505zm11.872-11.872-2.688 2.689-1.5-1.505 2.689-2.688a1.063 1.063 0 1 1 1.5 1.5zm-10.825-6.961 1.55-.442.442-1.55a1.191 1.191 0 0 1 2.29 0l.442 1.55 1.55.442a1.191 1.191 0 0 1 0 2.29l-1.55.442-.442 1.55a1.191 1.191 0 0 1 -2.29 0l-.442-1.55-1.55-.442a1.191 1.191 0 0 1 0-2.29zm18.274 14.29-1.55.442-.442 1.55a1.191 1.191 0 0 1 -2.29 0l-.442-1.55-1.55-.442a1.191 1.191 0 0 1 0-2.29l1.55-.442.442-1.55a1.191 1.191 0 0 1 2.29 0l.442 1.55 1.55.442a1.191 1.191 0 0 1 0 2.29zm-5.382-14.645 1.356-.387.389-1.358a1.042 1.042 0 0 1 2 0l.387 1.356 1.356.387a1.042 1.042 0 0 1 0 2l-1.356.387-.387 1.359a1.042 1.042 0 0 1 -2 0l-.387-1.355-1.358-.389a1.042 1.042 0 0 1 0-2z"/></svg>`)
//line webapp/public/class/index.qtpl:28
		}
//line webapp/public/class/index.qtpl:28
		qw422016.N().S(`</div></div></li>`)
//line webapp/public/class/index.qtpl:32
	}
//line webapp/public/class/index.qtpl:32
	qw422016.N().S(`</ul></section></main>`)
//line webapp/public/class/index.qtpl:36
}

//line webapp/public/class/index.qtpl:36
func WriteIndex(qq422016 qtio422016.Writer, students []*student.Student, edit bool) {
//line webapp/public/class/index.qtpl:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line webapp/public/class/index.qtpl:36
	StreamIndex(qw422016, students, edit)
//line webapp/public/class/index.qtpl:36
	qt422016.ReleaseWriter(qw422016)
//line webapp/public/class/index.qtpl:36
}

//line webapp/public/class/index.qtpl:36
func Index(students []*student.Student, edit bool) string {
//line webapp/public/class/index.qtpl:36
	qb422016 := qt422016.AcquireByteBuffer()
//line webapp/public/class/index.qtpl:36
	WriteIndex(qb422016, students, edit)
//line webapp/public/class/index.qtpl:36
	qs422016 := string(qb422016.B)
//line webapp/public/class/index.qtpl:36
	qt422016.ReleaseByteBuffer(qb422016)
//line webapp/public/class/index.qtpl:36
	return qs422016
//line webapp/public/class/index.qtpl:36
}
