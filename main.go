package ltsvExporter

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mkideal/cli"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	rev string
	ver string
)

type opt struct {
	Help    bool   `cli:"h,help" usage:"display help"`
	Version bool   `cli:"v,version" usage:"display version and revision"`
	Port    int    `cli:"p,port" usage:"set port number" dft:"6666"`
	LTSVURL string `cli:"*u,url" usage:"a URL of the LTSV"`
}

func Run(args []string) {
	cli.Run(&opt{}, func(ctx *cli.Context) error {
		argv := ctx.Argv().(*opt)
		if argv.Help {
			ctx.String(ctx.Usage())
			os.Exit(0)
		}

		if argv.Version {
			ctx.String(fmt.Sprintf("%s/%s\n", ver, rev))
			os.Exit(0)
		}

		port := argv.Port
		ltsvURL := argv.LTSVURL

		prometheus.MustRegister(newExporter(ltsvURL))

		http.Handle("/metrics", prometheus.Handler())

		addr := fmt.Sprintf(":%d", port)
		log.Print("Listening 127.0.0.1", addr)
		log.Fatal(http.ListenAndServe(addr, nil))

		return nil
	})
}
