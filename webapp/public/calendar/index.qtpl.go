// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line webapp/public/calendar/index.qtpl:2
package calendar

//line webapp/public/calendar/index.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line webapp/public/calendar/index.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line webapp/public/calendar/index.qtpl:2
func StreamIndex(qw422016 *qt422016.Writer) {
//line webapp/public/calendar/index.qtpl:2
	qw422016.N().S(`<main id="swup" class="animated transition-fade" namespace="calendar" no-return><section class="menu"><input id="datepicker" name="date" type="hidden"></section></main>`)
//line webapp/public/calendar/index.qtpl:8
}

//line webapp/public/calendar/index.qtpl:8
func WriteIndex(qq422016 qtio422016.Writer) {
//line webapp/public/calendar/index.qtpl:8
	qw422016 := qt422016.AcquireWriter(qq422016)
//line webapp/public/calendar/index.qtpl:8
	StreamIndex(qw422016)
//line webapp/public/calendar/index.qtpl:8
	qt422016.ReleaseWriter(qw422016)
//line webapp/public/calendar/index.qtpl:8
}

//line webapp/public/calendar/index.qtpl:8
func Index() string {
//line webapp/public/calendar/index.qtpl:8
	qb422016 := qt422016.AcquireByteBuffer()
//line webapp/public/calendar/index.qtpl:8
	WriteIndex(qb422016)
//line webapp/public/calendar/index.qtpl:8
	qs422016 := string(qb422016.B)
//line webapp/public/calendar/index.qtpl:8
	qt422016.ReleaseByteBuffer(qb422016)
//line webapp/public/calendar/index.qtpl:8
	return qs422016
//line webapp/public/calendar/index.qtpl:8
}
