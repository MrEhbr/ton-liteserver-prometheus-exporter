package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"

	"github.com/MrEhbr/ton-liteserver-prometheus-exporter/collector"
)

func main() {
	version, date := "unknown", "unknown"
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		for _, v := range buildInfo.Settings {
			if v.Value == "" {
				continue
			}
			switch v.Key {
			case "vcs.revision":
				version = v.Value
			case "vcs.time":
				date = v.Value
			}
		}
	}

	//nolint: forbidigo // We need to preatty print the version
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s - version %s\n", c.App.Name, version)
		fmt.Printf("  build date: \t%s\n", date)
		fmt.Printf("  go version: \t%s\n", buildInfo.GoVersion)
	}

	app := &cli.App{
		Name:    "lightserver-prometheus-exporter",
		Usage:   "Prometheus exporter for TON LightServer",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "port",
				Usage:   "Port to listen on for Prometheus metrics",
				EnvVars: []string{"TON_LITESERVER_PROMETHEUS_EXPORTER_PORT"},
				Value:   "9100",
			},
		},
		Action: func(c *cli.Context) error {
			collector := collector.NewMytonCollector(collector.NewParser())
			if err := prometheus.Register(collector); err != nil {
				return fmt.Errorf("error registering collector: %w", err)
			}

			cancelInterrupt := make(chan struct{})
			var g run.Group
			{
				prometheusListener, err := net.Listen("tcp", net.JoinHostPort("", c.String("port")))
				if err != nil {
					return fmt.Errorf("failed to listen: %w", err)
				}
				g.Add(func() error {
					log.Printf("Starting server on %s", prometheusListener.Addr())
					return http.Serve(prometheusListener, promhttp.Handler())
				}, func(error) {
					_ = prometheusListener.Close()
				})
			}
			{
				// This function just sits and waits for ctrl-C.
				g.Add(func() error {
					c := make(chan os.Signal, 1)
					signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
					select {
					case sig := <-c:
						return fmt.Errorf("received signal: %v", sig)
					case <-cancelInterrupt:
						return nil
					}
				}, func(error) {
					close(cancelInterrupt)
				})
			}

			return g.Run()
		},
		Commands: []*cli.Command{
			{
				Name: "print",
				Action: func(c *cli.Context) error {
					parser := collector.NewParser()
					metrics, err := parser.Parse()
					if err != nil {
						return fmt.Errorf("error collecting metrics: %w", err)
					}

					enc := json.NewEncoder(os.Stdout)
					enc.SetIndent("", "  ")

					if err := enc.Encode(metrics); err != nil {
						return fmt.Errorf("error encoding metrics: %w", err)
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("app run error: %s\n", err.Error())
	}
}
