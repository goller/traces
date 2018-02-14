package traces

import "net/http"

func SendSpan(req *http.Request, span *Span) {
	req.Header.Set("X-InfluxDays-TraceId", span.TraceID)
	req.Header.Set("X-InfluxDays-ParentSpanId", span.ID)
}

func NewSpanFromRequest(req *http.Request, name string) *Span {
	span := NewSpan(name)
	tid := req.Header.Get("X-InfluxDays-TraceId")
	pid := req.Header.Get("X-InfluxDays-ParentSpanId")
	if tid != "" && pid != "" {
		span.TraceID = tid
		span.ParentID = pid
	}
	return span
}
