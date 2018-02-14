package traces

import (
	"fmt"
	"net/http"
	"strings"
)

func LineProtocol(s *Span) string {
	return fmt.Sprintf(`traces,trace_id=%s,parent_id=%s,id=%s,name=%s duration_ns=%di %d`,
		s.TraceID,
		s.ParentID,
		s.ID,
		s.Name,
		s.EndTime.Sub(s.StartTime),
		s.StartTime.UnixNano())
}

func RecordSpan(span *Span) {
	http.Post(
		"http://localhost:8086/write?db=influxdays",
		"text/plain; charset=utf-8",
		strings.NewReader(LineProtocol(span)),
	)
}
