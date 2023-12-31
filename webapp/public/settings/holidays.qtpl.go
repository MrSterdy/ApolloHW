// Code generated by qtc from "holidays.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line webapp/public/settings/holidays.qtpl:2
package settings

//line webapp/public/settings/holidays.qtpl:2
import "github.com/MrSterdy/ApolloHW/model/holiday"

//line webapp/public/settings/holidays.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line webapp/public/settings/holidays.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line webapp/public/settings/holidays.qtpl:4
func StreamHolidays(qw422016 *qt422016.Writer, holidays []holiday.Holiday) {
//line webapp/public/settings/holidays.qtpl:4
	qw422016.N().S(`<main id="swup" class="animated transition-fade" namespace="control-holidays" main-button="СОХРАНИТЬ"><li id="template" style="display: none;"><div class="option-button"><div class="center column label"><div action="main"><svg class="medium-icon" viewBox="0 0 24 24"><path d="M21,4H17.9A5.009,5.009,0,0,0,13,0H11A5.009,5.009,0,0,0,6.1,4H3A1,1,0,0,0,3,6H4V19a5.006,5.006,0,0,0,5,5h6a5.006,5.006,0,0,0,5-5V6h1a1,1,0,0,0,0-2ZM11,2h2a3.006,3.006,0,0,1,2.829,2H8.171A3.006,3.006,0,0,1,11,2Zm7,17a3,3,0,0,1-3,3H9a3,3,0,0,1-3-3V6H18Z"/><path d="M10,18a1,1,0,0,0,1-1V11a1,1,0,0,0-2,0v6A1,1,0,0,0,10,18Z"/><path d="M14,18a1,1,0,0,0,1-1V11a1,1,0,0,0-2,0v6A1,1,0,0,0,14,18Z"/></svg></div><div class="center"><input type="text" name="holiday" class="medium center edit" readonly></div></div></div></li><form class="column menu"><ul class="column">`)
//line webapp/public/settings/holidays.qtpl:26
	for _, holiday := range holidays {
//line webapp/public/settings/holidays.qtpl:26
		qw422016.N().S(`<li><div class="option-button"><div class="center column label"><div action="main"><svg class="medium-icon" viewBox="0 0 24 24"><path d="M21,4H17.9A5.009,5.009,0,0,0,13,0H11A5.009,5.009,0,0,0,6.1,4H3A1,1,0,0,0,3,6H4V19a5.006,5.006,0,0,0,5,5h6a5.006,5.006,0,0,0,5-5V6h1a1,1,0,0,0,0-2ZM11,2h2a3.006,3.006,0,0,1,2.829,2H8.171A3.006,3.006,0,0,1,11,2Zm7,17a3,3,0,0,1-3,3H9a3,3,0,0,1-3-3V6H18Z"/><path d="M10,18a1,1,0,0,0,1-1V11a1,1,0,0,0-2,0v6A1,1,0,0,0,10,18Z"/><path d="M14,18a1,1,0,0,0,1-1V11a1,1,0,0,0-2,0v6A1,1,0,0,0,14,18Z"/></svg></div><div class="center"><input type="text" name="holiday" value="`)
//line webapp/public/settings/holidays.qtpl:39
		qw422016.E().S(holiday.String())
//line webapp/public/settings/holidays.qtpl:39
		qw422016.N().S(`" class="medium center edit" readonly></div></div></div></li>`)
//line webapp/public/settings/holidays.qtpl:44
	}
//line webapp/public/settings/holidays.qtpl:44
	qw422016.N().S(`</ul><button id="add" type="button"><svg class="icon" viewBox="0 0 24 24"><path d="M7,0H4A4,4,0,0,0,0,4V7a4,4,0,0,0,4,4H7a4,4,0,0,0,4-4V4A4,4,0,0,0,7,0ZM9,7A2,2,0,0,1,7,9H4A2,2,0,0,1,2,7V4A2,2,0,0,1,4,2H7A2,2,0,0,1,9,4Z"/><path d="M7,13H4a4,4,0,0,0-4,4v3a4,4,0,0,0,4,4H7a4,4,0,0,0,4-4V17A4,4,0,0,0,7,13Zm2,7a2,2,0,0,1-2,2H4a2,2,0,0,1-2-2V17a2,2,0,0,1,2-2H7a2,2,0,0,1,2,2Z"/><path d="M20,13H17a4,4,0,0,0-4,4v3a4,4,0,0,0,4,4h3a4,4,0,0,0,4-4V17A4,4,0,0,0,20,13Zm2,7a2,2,0,0,1-2,2H17a2,2,0,0,1-2-2V17a2,2,0,0,1,2-2h3a2,2,0,0,1,2,2Z"/><path d="M14,7h3v3a1,1,0,0,0,2,0V7h3a1,1,0,0,0,0-2H19V2a1,1,0,0,0-2,0V5H14a1,1,0,0,0,0,2Z"/></svg><input id="datepicker" type="hidden"></button></form></main>`)
//line webapp/public/settings/holidays.qtpl:59
}

//line webapp/public/settings/holidays.qtpl:59
func WriteHolidays(qq422016 qtio422016.Writer, holidays []holiday.Holiday) {
//line webapp/public/settings/holidays.qtpl:59
	qw422016 := qt422016.AcquireWriter(qq422016)
//line webapp/public/settings/holidays.qtpl:59
	StreamHolidays(qw422016, holidays)
//line webapp/public/settings/holidays.qtpl:59
	qt422016.ReleaseWriter(qw422016)
//line webapp/public/settings/holidays.qtpl:59
}

//line webapp/public/settings/holidays.qtpl:59
func Holidays(holidays []holiday.Holiday) string {
//line webapp/public/settings/holidays.qtpl:59
	qb422016 := qt422016.AcquireByteBuffer()
//line webapp/public/settings/holidays.qtpl:59
	WriteHolidays(qb422016, holidays)
//line webapp/public/settings/holidays.qtpl:59
	qs422016 := string(qb422016.B)
//line webapp/public/settings/holidays.qtpl:59
	qt422016.ReleaseByteBuffer(qb422016)
//line webapp/public/settings/holidays.qtpl:59
	return qs422016
//line webapp/public/settings/holidays.qtpl:59
}
