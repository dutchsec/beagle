package profiler

import (
	"bytes"
	"runtime"
	"strings"
	"time"

	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("go.dutchsec.com/beagle/profiler")

type profiler struct {
	now  time.Time
	prev time.Time

	s string
}

func New() profiler {
	p := profiler{
		prev: time.Now(),
		s:    "start",
		now:  time.Now(),
	}

	log.Infof("Profiler started: %s", p.findMethod())

	return p
}

func (p profiler) findMethod() string {
	trace := make([]byte, 1024)

	count := runtime.Stack(trace, true)
	trace = trace[:count]

	parts := bytes.Split(trace, []byte("\n"))

	return strings.TrimSpace(string(parts[6]))
}

func (p profiler) Done() {
	now := time.Now()

	if now.Sub(p.prev) > time.Second {
		log.Warningf("Profiler %s: %s took %v (%v)", p.findMethod(), p.s, now.Sub(p.now), now.Sub(p.prev))
	} else {
		log.Debugf("Profiler %s: %s took %v (%v)", p.findMethod(), p.s, now.Sub(p.prev), now.Sub(p.now))
	}
}

func (p *profiler) Report(s string) {
	now := time.Now()

	if now.Sub(p.prev) > time.Second {
		log.Warningf("Profiler %s: %s took %v (%v)", p.findMethod(), p.s, now.Sub(p.now), now.Sub(p.prev))
	} else {
		log.Debugf("Profiler %s: %s took %v (%v)", p.findMethod(), p.s, now.Sub(p.now), now.Sub(p.prev))
	}

	p.s = s
	p.prev = now
}
