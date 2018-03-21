package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"git.containerum.net/ch/kube-api/pkg/kubernetes"
	"git.containerum.net/ch/kube-api/pkg/router"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "CH_KUBE_API_DEBUG",
		Name:   "debug",
		Usage:  "start the server in debug mode",
	},
	cli.StringFlag{
		EnvVar: "CH_KUBE_API_KUBE_CONF",
		Name:   "kubeconf",
		Usage:  "config file for kubernetes apiserver client",
	},
	cli.BoolFlag{
		EnvVar: "CH_KUBE_API_LOG",
		Name:   "textlog",
		Usage:  "output log in text format",
	},
}

func server(c *cli.Context) error {
	if c.Bool("debug") {
		gin.SetMode(gin.DebugMode)
		log.SetLevel(log.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
		log.SetLevel(log.InfoLevel)
	}

	if c.Bool("textlog") {
		log.SetFormatter(&log.TextFormatter{})
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}

	kube := kubernetes.Kube{}
	kube.RegisterClient(c.String("kubeconf"))

	app := router.CreateRouter(&kube, c.Bool("debug"))

	// for graceful shutdown
	srv := &http.Server{
		Addr:    ":1212",
		Handler: app,
	}

	go exitOnErr(srv.ListenAndServe())

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) // subscribe on interrupt event
	<-quit                            // wait for event
	log.Infoln("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
