package formats

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"git.sr.ht/~poldi1405/glog"
	"github.com/valyala/fasthttp"
)

// Common log format as described in https://wikiless.org/wiki/Common_Log_Format?lang=en

func CommonLogFormat(lvl glog.Level, t time.Time, caller string, message string) string {
	return fmt.Sprintf(message, t.Format("02/Jan/2006:15:04:05 -0700"))
}

// CLFormatFromRequest allows passing a net/http.Request to generate the
// message needed for proper CLF
func CLFormatFromRequest(r *http.Request, bodySize int64, userIdentifier ...string) string {
	buf := &strings.Builder{}
	res := csv.NewWriter(buf)

	user, _, ok := r.BasicAuth()
	if !ok {
		user = "-"
	}

	userIdentifier = append(userIdentifier, "-")

	res.Comma = ' '
	res.Write([]string{
		r.RemoteAddr,
		userIdentifier[0],
		user,
		"[%s]",
		fmt.Sprintf("%s %s %s", r.Method, r.URL.Path, r.Proto),
		strconv.Itoa(r.Response.StatusCode),
		strconv.FormatInt(bodySize, 10),
	})

	return buf.String()
}

// CLFormatFromRequest allows passing a github.com/valyala/fasthttp.RequestCtx
// to generate the message needed for proper CLF
func CLFormatFromRequestCtx(ctx *fasthttp.RequestCtx, userIdentifier ...string) string {
	buf := bytes.NewBuffer([]byte{})
	res := csv.NewWriter(buf)

	var user string
	auth := ctx.Request.Header.Peek("Authorization")
	if len(auth) == 0 {
		user = "-"
	} else {
		// shamelessly stolen from net/http/request.go
		// dont look at it{{{
		// seriously, don't.{{{
		// okay, you've been warned.{{{
		u, _, _ := func(auth string) (username, password string, ok bool) {
			const prefix = "Basic "
			username = "-"
			lower := func(b byte) byte {
				if 'A' <= b && b <= 'Z' {
					return b + ('a' - 'A')
				}
				return b
			}

			// Case insensitive prefix match. See Issue 22736.
			if len(auth) < len(prefix) || !func(s, t string) bool {
				if len(s) != len(t) {
					return false
				}
				for i := 0; i < len(s); i++ {
					if lower(s[i]) != lower(t[i]) {
						return false
					}
				}
				return true
			}(auth[:len(prefix)], prefix) {
				return
			}
			c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
			if err != nil {
				return
			}
			cs := string(c)
			s := strings.IndexByte(cs, ':')
			if s < 0 {
				return
			}
			return cs[:s], cs[s+1:], true
		}(string(auth)) // }}}}}}}}}
		user = u
	}

	userIdentifier = append(userIdentifier, "-")

	res.Comma = ' '
	res.Write([]string{
		ctx.RemoteAddr().String(),
		userIdentifier[0],
		user,
		"[%s]",
		fmt.Sprintf("%s %s %s", string(ctx.Method()), string(ctx.Request.URI().Path()), ctx.Request.Header.Protocol()),
		strconv.Itoa(ctx.Response.StatusCode()),
		strconv.Itoa(len(ctx.Response.Body())),
	})

	return buf.String()
}

var _ glog.FormatFunction = CommonLogFormat
