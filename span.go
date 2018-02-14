package traces

import (
	"time"

	"github.com/rs/xid"
)

type Span struct {
	TraceID   string
	ParentID  string
	ID        string
	StartTime time.Time
	EndTime   time.Time
	Name      string
}

func NewSpan(name string) *Span {
	id := xid.New().String()
	return &Span{
		TraceID:   id,
		ParentID:  id,
		ID:        id,
		Name:      name,
		StartTime: time.Now(),
	}
}

func (s *Span) Finish() {
	s.EndTime = time.Now()
	RecordSpan(s)
}
