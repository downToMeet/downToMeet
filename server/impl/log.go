package impl

import log "github.com/sirupsen/logrus"

// RequestLogHook is a log.Hook implementation that adds to the log.Entry
// information gleaned from the entry context, such as the HTTP request (using
// RequestFromContext) and the user session (using SessionFromContext).
type RequestLogHook struct{}

var _ log.Hook = RequestLogHook{}

// Levels returns log.AllLevels and implements log.Hook.
func (r RequestLogHook) Levels() []log.Level { return log.AllLevels }

// Fire implements log.Hook.
func (r RequestLogHook) Fire(entry *log.Entry) error {
	if ctx := entry.Context; ctx != nil {
		if r := RequestFromContext(ctx); r != nil {
			tentativeAddField(entry.Data, "method", r.Method)
			tentativeAddField(entry.Data, "url", r.URL.String())
		}
		if s := SessionFromContext(ctx); s != nil {
			tentativeAddField(entry.Data, "is_logged_in", !s.IsNew)
			tentativeAddField(entry.Data, "session", s.Values)
		}
	}
	return nil
}

func tentativeAddField(existing log.Fields, k string, v interface{}) {
	if _, ok := existing[k]; !ok {
		existing[k] = v
	}
}
