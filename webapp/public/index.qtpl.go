// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line webapp/public/index.qtpl:2
package public

//line webapp/public/index.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line webapp/public/index.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line webapp/public/index.qtpl:2
func StreamIndex(qw422016 *qt422016.Writer) {
//line webapp/public/index.qtpl:2
	qw422016.N().S(`<!DOCTYPE html><html lang="ru"><head><meta charset="UTF-8"><title>Apollo</title><script src="https://telegram.org/js/telegram-web-app.js"></script></head><body><script>document.cookie = "apollo_initdata=" + Telegram.WebApp.initData + "; SameSite=None; Secure";location.replace(`)
//line webapp/public/index.qtpl:2
	qw422016.N().S("`")
//line webapp/public/index.qtpl:2
	qw422016.N().S(`/?auth`)
//line webapp/public/index.qtpl:2
	qw422016.N().S("`")
//line webapp/public/index.qtpl:2
	qw422016.N().S(`);</script></body></html>`)
//line webapp/public/index.qtpl:20
}

//line webapp/public/index.qtpl:20
func WriteIndex(qq422016 qtio422016.Writer) {
//line webapp/public/index.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line webapp/public/index.qtpl:20
	StreamIndex(qw422016)
//line webapp/public/index.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line webapp/public/index.qtpl:20
}

//line webapp/public/index.qtpl:20
func Index() string {
//line webapp/public/index.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
//line webapp/public/index.qtpl:20
	WriteIndex(qb422016)
//line webapp/public/index.qtpl:20
	qs422016 := string(qb422016.B)
//line webapp/public/index.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
//line webapp/public/index.qtpl:20
	return qs422016
//line webapp/public/index.qtpl:20
}