package rps

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"go.k6.io/k6/event"
	"go.k6.io/k6/js/modules"
	"golang.org/x/time/rate"
)

func init() {
	fmt.Println("Register")
	modules.Register("k6/x/rps", New())
}

type RootModule struct {
	RPS *rate.Limiter
}

func New() modules.Module {
	rps, err := strconv.ParseFloat(strings.TrimSpace(os.Getenv("RPS")), 64)
	var limiter *rate.Limiter
	if err == nil {
		limiter = rate.NewLimiter(rate.Limit(rps), 1)
	}

	return &RootModule{
		RPS: limiter,
	}
}

func (root *RootModule) NewModuleInstance(vu modules.VU) modules.Instance { // nolint:varnamelen
	mod := &Module{
		vu: vu,
		c:  &Object{},
	}
	sub, ch := vu.Events().Global.Subscribe(event.TestStart)
	go func() {
		ev := <-ch
		if vu.State() != nil && root.RPS != nil {
			vu.State().RPSLimit = root.RPS
		}
		ev.Done()
		vu.Events().Global.Unsubscribe(sub)
	}()

	return mod
}

type Module struct {
	vu modules.VU
	c  *Object
}

type Object struct {
}

func (mod *Module) Exports() modules.Exports {
	return modules.Exports{
		Default: mod.c,
	}
}
