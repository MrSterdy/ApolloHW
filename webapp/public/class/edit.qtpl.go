// Code generated by qtc from "edit.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line webapp/public/class/edit.qtpl:2
package class

//line webapp/public/class/edit.qtpl:2
import "github.com/MrSterdy/ApolloHW/model/student"

//line webapp/public/class/edit.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line webapp/public/class/edit.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line webapp/public/class/edit.qtpl:4
func StreamEdit(qw422016 *qt422016.Writer, active []*student.Student, inactive []*student.Student) {
//line webapp/public/class/edit.qtpl:4
	qw422016.N().S(`<main id="swup" class="animated transition-fade" namespace="class-edit" main-button="СОХРАНИТЬ"><form class="column menu"><ul class="column"><li id="applications"><div class="option-button"><div class="center prefix"><svg class="icon" viewBox="0 0 24 24"><path d="m17 24a7 7 0 1 1 7-7 7.008 7.008 0 0 1 -7 7zm0-12a5 5 0 1 0 5 5 5.006 5.006 0 0 0 -5-5zm1.707 6.707a1 1 0 0 0 0-1.414l-.707-.707v-1.586a1 1 0 0 0 -2 0v2a1 1 0 0 0 .293.707l1 1a1 1 0 0 0 1.414 0zm-16.707 4.293a8.3 8.3 0 0 1 6.221-8.024 1 1 0 0 0 -.442-1.952 10.213 10.213 0 0 0 -7.779 9.976 1 1 0 0 0 2 0zm6.474-12a5.5 5.5 0 1 1 5.5-5.5 5.506 5.506 0 0 1 -5.5 5.5zm0-9a3.5 3.5 0 1 0 3.5 3.5 3.5 3.5 0 0 0 -3.5-3.5z"/></svg></div><div class="center label"><span>Заявки</span></div></div><ul class="option-menu">`)
//line webapp/public/class/edit.qtpl:22
	for _, s := range inactive {
//line webapp/public/class/edit.qtpl:22
		qw422016.N().S(`<li class="option-button"><input type="number" name="id" value="`)
//line webapp/public/class/edit.qtpl:24
		qw422016.N().DL(s.ID)
//line webapp/public/class/edit.qtpl:24
		qw422016.N().S(`" hidden><div class="center prefix"><svg class="icon" viewBox="0 0 24 24"><path d="M12,12A6,6,0,1,0,6,6,6.006,6.006,0,0,0,12,12ZM12,2A4,4,0,1,1,8,6,4,4,0,0,1,12,2Z"/><path d="M12,14a9.01,9.01,0,0,0-9,9,1,1,0,0,0,2,0,7,7,0,0,1,14,0,1,1,0,0,0,2,0A9.01,9.01,0,0,0,12,14Z"/></svg></div><div class="label"><div class="column"><input class="edit" name="last_name" type="text" value="`)
//line webapp/public/class/edit.qtpl:35
		qw422016.E().S(s.LastName)
//line webapp/public/class/edit.qtpl:35
		qw422016.N().S(`" required><input class="edit" name="first_name" type="text" value="`)
//line webapp/public/class/edit.qtpl:36
		qw422016.E().S(s.FirstName)
//line webapp/public/class/edit.qtpl:36
		qw422016.N().S(`" required></div><div class="row actions"><input type="number" name="role" min="`)
//line webapp/public/class/edit.qtpl:40
		qw422016.N().D(int(student.RoleStudent))
//line webapp/public/class/edit.qtpl:40
		qw422016.N().S(`" max="`)
//line webapp/public/class/edit.qtpl:40
		qw422016.N().D(int(student.RoleAdmin))
//line webapp/public/class/edit.qtpl:40
		qw422016.N().S(`" value="`)
//line webapp/public/class/edit.qtpl:40
		qw422016.N().D(int(s.Role))
//line webapp/public/class/edit.qtpl:40
		qw422016.N().S(`" hidden><div class="center" role="`)
//line webapp/public/class/edit.qtpl:42
		qw422016.N().D(int(student.RoleRedactor))
//line webapp/public/class/edit.qtpl:42
		qw422016.N().S(`"`)
//line webapp/public/class/edit.qtpl:42
		if s.Role != student.RoleRedactor {
//line webapp/public/class/edit.qtpl:42
			qw422016.N().S(`style="display: none;"`)
//line webapp/public/class/edit.qtpl:42
		}
//line webapp/public/class/edit.qtpl:42
		qw422016.N().S(`><svg class="medium-icon" viewBox="0 0 24 24"><path d="m17.994 2.286a9 9 0 0 0 -14.919 5.536 8.938 8.938 0 0 0 2.793 7.761 6.263 6.263 0 0 1 2.132 4.566v.161a3.694 3.694 0 0 0 3.69 3.69h.62a3.694 3.694 0 0 0 3.69-3.69v-.549a5.323 5.323 0 0 1 1.932-4 8.994 8.994 0 0 0 .062-13.477zm-5.684 19.714h-.62a1.692 1.692 0 0 1 -1.69-1.69s-.007-.26-.008-.31h4.008v.31a1.692 1.692 0 0 1 -1.69 1.69zm4.3-7.741a7.667 7.667 0 0 0 -2.364 3.741h-1.246v-7.184a3 3 0 0 0 2-2.816 1 1 0 0 0 -2 0 1 1 0 0 1 -2 0 1 1 0 0 0 -2 0 3 3 0 0 0 2 2.816v7.184h-1.322a8.634 8.634 0 0 0 -2.448-3.881 7 7 0 0 1 3.951-12.073 7.452 7.452 0 0 1 .828-.046 6.921 6.921 0 0 1 4.652 1.778 6.993 6.993 0 0 1 -.048 10.481z"/></svg></div><div class="center" role="`)
//line webapp/public/class/edit.qtpl:47
		qw422016.N().D(int(student.RoleAdmin))
//line webapp/public/class/edit.qtpl:47
		qw422016.N().S(`"`)
//line webapp/public/class/edit.qtpl:47
		if s.Role != student.RoleAdmin {
//line webapp/public/class/edit.qtpl:47
			qw422016.N().S(`style="display: none;"`)
//line webapp/public/class/edit.qtpl:47
		}
//line webapp/public/class/edit.qtpl:47
		qw422016.N().S(`><svg class="medium-icon" viewBox="0 0 24 24"><path d="m18 9.064a3.049 3.049 0 0 0 -.9-2.164 3.139 3.139 0 0 0 -4.334 0l-11.866 11.869a3.064 3.064 0 0 0 4.33 4.331l11.87-11.869a3.047 3.047 0 0 0 .9-2.167zm-14.184 12.624a1.087 1.087 0 0 1 -1.5 0 1.062 1.062 0 0 1 0-1.5l7.769-7.77 1.505 1.505zm11.872-11.872-2.688 2.689-1.5-1.505 2.689-2.688a1.063 1.063 0 1 1 1.5 1.5zm-10.825-6.961 1.55-.442.442-1.55a1.191 1.191 0 0 1 2.29 0l.442 1.55 1.55.442a1.191 1.191 0 0 1 0 2.29l-1.55.442-.442 1.55a1.191 1.191 0 0 1 -2.29 0l-.442-1.55-1.55-.442a1.191 1.191 0 0 1 0-2.29zm18.274 14.29-1.55.442-.442 1.55a1.191 1.191 0 0 1 -2.29 0l-.442-1.55-1.55-.442a1.191 1.191 0 0 1 0-2.29l1.55-.442.442-1.55a1.191 1.191 0 0 1 2.29 0l.442 1.55 1.55.442a1.191 1.191 0 0 1 0 2.29zm-5.382-14.645 1.356-.387.389-1.358a1.042 1.042 0 0 1 2 0l.387 1.356 1.356.387a1.042 1.042 0 0 1 0 2l-1.356.387-.387 1.359a1.042 1.042 0 0 1 -2 0l-.387-1.355-1.358-.389a1.042 1.042 0 0 1 0-2z"/></svg></div><div action="promote" class="`)
//line webapp/public/class/edit.qtpl:53
		if s.Role == student.RoleAdmin {
//line webapp/public/class/edit.qtpl:53
			qw422016.N().S(`inactive`)
//line webapp/public/class/edit.qtpl:53
		}
//line webapp/public/class/edit.qtpl:53
		qw422016.N().S(`"><svg class="mini-icon" viewBox="0 0 24 24"><path d="M.012,12a1,1,0,0,1,.293-.707L8.477,3.122a5,5,0,0,1,7.07,0l8.172,8.171a1,1,0,0,1-1.414,1.414L14.133,4.536a3,3,0,0,0-4.242,0L1.719,12.707A1,1,0,0,1,.012,12Z"/><path d="M.012,22a1,1,0,0,1,.293-.707l9.586-9.586a3,3,0,0,1,4.242,0l9.586,9.586a1,1,0,0,1-1.414,1.414l-9.586-9.586a1,1,0,0,0-1.414,0L1.719,22.707A1,1,0,0,1,.012,22Z"/></svg></div><div action="demote" class="`)
//line webapp/public/class/edit.qtpl:59
		if s.Role == student.RoleStudent {
//line webapp/public/class/edit.qtpl:59
			qw422016.N().S(`inactive`)
//line webapp/public/class/edit.qtpl:59
		}
//line webapp/public/class/edit.qtpl:59
		qw422016.N().S(`"><svg class="mini-icon" viewBox="0 0 24 24"><path d="M.305,12.293a1,1,0,0,1,1.414,0l8.172,8.171a3,3,0,0,0,4.242,0l8.172-8.171a1,1,0,0,1,1.414,1.414l-8.172,8.171a5,5,0,0,1-7.07,0L.305,13.707a1,1,0,0,1,0-1.414Z"/><path d="M.305,2.293a1,1,0,0,1,1.414,0l9.586,9.586a1,1,0,0,0,1.414,0l9.586-9.586a1,1,0,0,1,1.414,1.414l-9.586,9.586a3,3,0,0,1-4.242,0L.305,3.707a1,1,0,0,1,0-1.414Z"/></svg></div><div action="accept"><svg class="mini-icon" viewBox="0 0 24 24"><path d="M22.319,4.431,8.5,18.249a1,1,0,0,1-1.417,0L1.739,12.9a1,1,0,0,0-1.417,0h0a1,1,0,0,0,0,1.417l5.346,5.345a3.008,3.008,0,0,0,4.25,0L23.736,5.847a1,1,0,0,0,0-1.416h0A1,1,0,0,0,22.319,4.431Z"/></svg></div><div action="delete"><svg class="mini-icon" viewBox="0 0 24 24"><path d="M21,4H17.9A5.009,5.009,0,0,0,13,0H11A5.009,5.009,0,0,0,6.1,4H3A1,1,0,0,0,3,6H4V19a5.006,5.006,0,0,0,5,5h6a5.006,5.006,0,0,0,5-5V6h1a1,1,0,0,0,0-2ZM11,2h2a3.006,3.006,0,0,1,2.829,2H8.171A3.006,3.006,0,0,1,11,2Zm7,17a3,3,0,0,1-3,3H9a3,3,0,0,1-3-3V6H18Z"/><path d="M10,18a1,1,0,0,0,1-1V11a1,1,0,0,0-2,0v6A1,1,0,0,0,10,18Z"/><path d="M14,18a1,1,0,0,0,1-1V11a1,1,0,0,0-2,0v6A1,1,0,0,0,14,18Z"/></svg></div></div></div></li>`)
//line webapp/public/class/edit.qtpl:82
	}
//line webapp/public/class/edit.qtpl:82
	qw422016.N().S(`</ul></li><li id="students"><div class="option-button"><div class="center prefix"><svg class="icon" viewBox="0 0 24 24"><path d="m7.5 13a4.5 4.5 0 1 1 4.5-4.5 4.505 4.505 0 0 1 -4.5 4.5zm0-7a2.5 2.5 0 1 0 2.5 2.5 2.5 2.5 0 0 0 -2.5-2.5zm7.5 17v-.5a7.5 7.5 0 0 0 -15 0v.5a1 1 0 0 0 2 0v-.5a5.5 5.5 0 0 1 11 0v.5a1 1 0 0 0 2 0zm9-5a7 7 0 0 0 -11.667-5.217 1 1 0 1 0 1.334 1.49 5 5 0 0 1 8.333 3.727 1 1 0 0 0 2 0zm-6.5-9a4.5 4.5 0 1 1 4.5-4.5 4.505 4.505 0 0 1 -4.5 4.5zm0-7a2.5 2.5 0 1 0 2.5 2.5 2.5 2.5 0 0 0 -2.5-2.5z"/></svg></div><div class="center label"><span>Ученики</span></div></div><ul class="option-menu">`)
//line webapp/public/class/edit.qtpl:100
	for _, s := range active {
//line webapp/public/class/edit.qtpl:100
		qw422016.N().S(`<li class="option-button"><input type="number" name="id" value="`)
//line webapp/public/class/edit.qtpl:102
		qw422016.N().DL(s.ID)
//line webapp/public/class/edit.qtpl:102
		qw422016.N().S(`" hidden><div class="center prefix"><svg class="icon" viewBox="0 0 24 24"><path d="M12,12A6,6,0,1,0,6,6,6.006,6.006,0,0,0,12,12ZM12,2A4,4,0,1,1,8,6,4,4,0,0,1,12,2Z"/><path d="M12,14a9.01,9.01,0,0,0-9,9,1,1,0,0,0,2,0,7,7,0,0,1,14,0,1,1,0,0,0,2,0A9.01,9.01,0,0,0,12,14Z"/></svg></div><div class="column label"><div class="column"><input class="edit" name="last_name" type="text" value="`)
//line webapp/public/class/edit.qtpl:113
		qw422016.E().S(s.LastName)
//line webapp/public/class/edit.qtpl:113
		qw422016.N().S(`" required><input class="edit" name="first_name" type="text" value="`)
//line webapp/public/class/edit.qtpl:114
		qw422016.E().S(s.FirstName)
//line webapp/public/class/edit.qtpl:114
		qw422016.N().S(`" required></div><div class="row actions"><input type="number" name="role" min="`)
//line webapp/public/class/edit.qtpl:118
		qw422016.N().D(int(student.RoleStudent))
//line webapp/public/class/edit.qtpl:118
		qw422016.N().S(`" max="`)
//line webapp/public/class/edit.qtpl:118
		qw422016.N().D(int(student.RoleAdmin))
//line webapp/public/class/edit.qtpl:118
		qw422016.N().S(`" value="`)
//line webapp/public/class/edit.qtpl:118
		qw422016.N().D(int(s.Role))
//line webapp/public/class/edit.qtpl:118
		qw422016.N().S(`" hidden><div class="center" role="`)
//line webapp/public/class/edit.qtpl:120
		qw422016.N().D(int(student.RoleRedactor))
//line webapp/public/class/edit.qtpl:120
		qw422016.N().S(`"`)
//line webapp/public/class/edit.qtpl:120
		if s.Role != student.RoleRedactor {
//line webapp/public/class/edit.qtpl:120
			qw422016.N().S(`style="display: none;"`)
//line webapp/public/class/edit.qtpl:120
		}
//line webapp/public/class/edit.qtpl:120
		qw422016.N().S(`><svg class="medium-icon" viewBox="0 0 24 24"><path d="m17.994 2.286a9 9 0 0 0 -14.919 5.536 8.938 8.938 0 0 0 2.793 7.761 6.263 6.263 0 0 1 2.132 4.566v.161a3.694 3.694 0 0 0 3.69 3.69h.62a3.694 3.694 0 0 0 3.69-3.69v-.549a5.323 5.323 0 0 1 1.932-4 8.994 8.994 0 0 0 .062-13.477zm-5.684 19.714h-.62a1.692 1.692 0 0 1 -1.69-1.69s-.007-.26-.008-.31h4.008v.31a1.692 1.692 0 0 1 -1.69 1.69zm4.3-7.741a7.667 7.667 0 0 0 -2.364 3.741h-1.246v-7.184a3 3 0 0 0 2-2.816 1 1 0 0 0 -2 0 1 1 0 0 1 -2 0 1 1 0 0 0 -2 0 3 3 0 0 0 2 2.816v7.184h-1.322a8.634 8.634 0 0 0 -2.448-3.881 7 7 0 0 1 3.951-12.073 7.452 7.452 0 0 1 .828-.046 6.921 6.921 0 0 1 4.652 1.778 6.993 6.993 0 0 1 -.048 10.481z"/></svg></div><div class="center" role="`)
//line webapp/public/class/edit.qtpl:125
		qw422016.N().D(int(student.RoleAdmin))
//line webapp/public/class/edit.qtpl:125
		qw422016.N().S(`"`)
//line webapp/public/class/edit.qtpl:125
		if s.Role != student.RoleAdmin {
//line webapp/public/class/edit.qtpl:125
			qw422016.N().S(`style="display: none;"`)
//line webapp/public/class/edit.qtpl:125
		}
//line webapp/public/class/edit.qtpl:125
		qw422016.N().S(`><svg class="medium-icon" viewBox="0 0 24 24"><path d="m18 9.064a3.049 3.049 0 0 0 -.9-2.164 3.139 3.139 0 0 0 -4.334 0l-11.866 11.869a3.064 3.064 0 0 0 4.33 4.331l11.87-11.869a3.047 3.047 0 0 0 .9-2.167zm-14.184 12.624a1.087 1.087 0 0 1 -1.5 0 1.062 1.062 0 0 1 0-1.5l7.769-7.77 1.505 1.505zm11.872-11.872-2.688 2.689-1.5-1.505 2.689-2.688a1.063 1.063 0 1 1 1.5 1.5zm-10.825-6.961 1.55-.442.442-1.55a1.191 1.191 0 0 1 2.29 0l.442 1.55 1.55.442a1.191 1.191 0 0 1 0 2.29l-1.55.442-.442 1.55a1.191 1.191 0 0 1 -2.29 0l-.442-1.55-1.55-.442a1.191 1.191 0 0 1 0-2.29zm18.274 14.29-1.55.442-.442 1.55a1.191 1.191 0 0 1 -2.29 0l-.442-1.55-1.55-.442a1.191 1.191 0 0 1 0-2.29l1.55-.442.442-1.55a1.191 1.191 0 0 1 2.29 0l.442 1.55 1.55.442a1.191 1.191 0 0 1 0 2.29zm-5.382-14.645 1.356-.387.389-1.358a1.042 1.042 0 0 1 2 0l.387 1.356 1.356.387a1.042 1.042 0 0 1 0 2l-1.356.387-.387 1.359a1.042 1.042 0 0 1 -2 0l-.387-1.355-1.358-.389a1.042 1.042 0 0 1 0-2z"/></svg></div><div action="promote" class="`)
//line webapp/public/class/edit.qtpl:131
		if s.Role == student.RoleAdmin {
//line webapp/public/class/edit.qtpl:131
			qw422016.N().S(`inactive`)
//line webapp/public/class/edit.qtpl:131
		}
//line webapp/public/class/edit.qtpl:131
		qw422016.N().S(`"><svg class="mini-icon" viewBox="0 0 24 24"><path d="M.012,12a1,1,0,0,1,.293-.707L8.477,3.122a5,5,0,0,1,7.07,0l8.172,8.171a1,1,0,0,1-1.414,1.414L14.133,4.536a3,3,0,0,0-4.242,0L1.719,12.707A1,1,0,0,1,.012,12Z"/><path d="M.012,22a1,1,0,0,1,.293-.707l9.586-9.586a3,3,0,0,1,4.242,0l9.586,9.586a1,1,0,0,1-1.414,1.414l-9.586-9.586a1,1,0,0,0-1.414,0L1.719,22.707A1,1,0,0,1,.012,22Z"/></svg></div><div action="demote" class="`)
//line webapp/public/class/edit.qtpl:137
		if s.Role == student.RoleStudent {
//line webapp/public/class/edit.qtpl:137
			qw422016.N().S(`inactive`)
//line webapp/public/class/edit.qtpl:137
		}
//line webapp/public/class/edit.qtpl:137
		qw422016.N().S(`"><svg class="mini-icon" viewBox="0 0 24 24"><path d="M.305,12.293a1,1,0,0,1,1.414,0l8.172,8.171a3,3,0,0,0,4.242,0l8.172-8.171a1,1,0,0,1,1.414,1.414l-8.172,8.171a5,5,0,0,1-7.07,0L.305,13.707a1,1,0,0,1,0-1.414Z"/><path d="M.305,2.293a1,1,0,0,1,1.414,0l9.586,9.586a1,1,0,0,0,1.414,0l9.586-9.586a1,1,0,0,1,1.414,1.414l-9.586,9.586a3,3,0,0,1-4.242,0L.305,3.707a1,1,0,0,1,0-1.414Z"/></svg></div><div action="accept" style="display: none;"><svg class="mini-icon" viewBox="0 0 24 24"><path d="M22.319,4.431,8.5,18.249a1,1,0,0,1-1.417,0L1.739,12.9a1,1,0,0,0-1.417,0h0a1,1,0,0,0,0,1.417l5.346,5.345a3.008,3.008,0,0,0,4.25,0L23.736,5.847a1,1,0,0,0,0-1.416h0A1,1,0,0,0,22.319,4.431Z"/></svg></div><div action="delete"><svg class="mini-icon" viewBox="0 0 24 24"><path d="M21,4H17.9A5.009,5.009,0,0,0,13,0H11A5.009,5.009,0,0,0,6.1,4H3A1,1,0,0,0,3,6H4V19a5.006,5.006,0,0,0,5,5h6a5.006,5.006,0,0,0,5-5V6h1a1,1,0,0,0,0-2ZM11,2h2a3.006,3.006,0,0,1,2.829,2H8.171A3.006,3.006,0,0,1,11,2Zm7,17a3,3,0,0,1-3,3H9a3,3,0,0,1-3-3V6H18Z"/><path d="M10,18a1,1,0,0,0,1-1V11a1,1,0,0,0-2,0v6A1,1,0,0,0,10,18Z"/><path d="M14,18a1,1,0,0,0,1-1V11a1,1,0,0,0-2,0v6A1,1,0,0,0,14,18Z"/></svg></div></div></div></li>`)
//line webapp/public/class/edit.qtpl:160
	}
//line webapp/public/class/edit.qtpl:160
	qw422016.N().S(`</ul></li></ul></form></main>`)
//line webapp/public/class/edit.qtpl:166
}

//line webapp/public/class/edit.qtpl:166
func WriteEdit(qq422016 qtio422016.Writer, active []*student.Student, inactive []*student.Student) {
//line webapp/public/class/edit.qtpl:166
	qw422016 := qt422016.AcquireWriter(qq422016)
//line webapp/public/class/edit.qtpl:166
	StreamEdit(qw422016, active, inactive)
//line webapp/public/class/edit.qtpl:166
	qt422016.ReleaseWriter(qw422016)
//line webapp/public/class/edit.qtpl:166
}

//line webapp/public/class/edit.qtpl:166
func Edit(active []*student.Student, inactive []*student.Student) string {
//line webapp/public/class/edit.qtpl:166
	qb422016 := qt422016.AcquireByteBuffer()
//line webapp/public/class/edit.qtpl:166
	WriteEdit(qb422016, active, inactive)
//line webapp/public/class/edit.qtpl:166
	qs422016 := string(qb422016.B)
//line webapp/public/class/edit.qtpl:166
	qt422016.ReleaseByteBuffer(qb422016)
//line webapp/public/class/edit.qtpl:166
	return qs422016
//line webapp/public/class/edit.qtpl:166
}
