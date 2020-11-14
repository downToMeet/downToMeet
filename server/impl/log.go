package impl

import log "github.com/sirupsen/logrus"

type RequestFormatter struct{ log.Formatter }

func (f RequestFormatter) Format(entry *log.Entry) ([]byte, error) {
	if ctx := entry.Context; ctx != nil {
		if r := RequestFromSession(ctx); r != nil {
			tentativeAddField(entry.Data, "method", r.Method)
			tentativeAddField(entry.Data, "url", r.URL.String())
		}
		if s := SessionFromContext(ctx); s != nil {
			tentativeAddField(entry.Data, "is_logged_in", !s.IsNew)
			tentativeAddField(entry.Data, "session", s.Values)
		}
	}

	return f.Formatter.Format(entry)
}

func tentativeAddField(existing log.Fields, k string, v interface{}) {
	if _, ok := existing[k]; !ok {
		existing[k] = v
	}
}
