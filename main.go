package ltsvExporter

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mkideal/cli"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	rev string
	ver string
)

type opt struct {
	cli.Helper
	Version bool   `cli:"!v,version" usage:"display version and revision"`
	Port    int    `cli:"p,port" usage:"set the port number to listen" dft:"6666"`
	LTSVURL string `cli:"*u,url" usage:"set a URL of the LTSV"`
}

func Run(args []string) {
	cli.Run(&opt{}, func(ctx *cli.Context) error {
		argv := ctx.Argv().(*opt)

		if argv.Version {
			ctx.String(fmt.Sprintf("%s/%s\n", ver, rev))
			return nil
		}

		ltsvScraper := newRemoteLTSVScraper(argv.LTSVURL)

		prometheus.MustRegister(newExporter(ltsvScraper))

		http.Handle("/metrics", prometheus.Handler())

		addr := fmt.Sprintf(":%d", argv.Port)
		log.Print("Listening 127.0.0.1", addr)
		log.Fatal(http.ListenAndServe(addr, nil))

		return nil
	})
}
