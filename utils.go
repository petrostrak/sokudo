package sokudo

import (
	"fmt"
	"regexp"
	"runtime"
	"time"
)

// LoadTime calculates function execution time. To use, add
// defer c.LoadTime(time.Now()) to the function body
func (s *Sokudo) LoadTime(start time.Time) {
	elapsed := time.Since(start)
	pc, _, _, _ := runtime.Caller(1)
	funcObj := runtime.FuncForPC(pc)
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	s.InfoLog.Println(fmt.Sprintf("Load Time: %s took %s", name, elapsed))
}
