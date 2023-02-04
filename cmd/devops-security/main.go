package main

import (
	"github.com/alexflint/go-arg"
	"github.com/denysvitali/devops-security-cli/pkg/client"
	"github.com/sirupsen/logrus"
	"strings"
)

const AppName = "devops-security"

var logger = logrus.New()
var config *Config = nil

var args struct {
	LogLevel string `arg:"--log-level" default:"warn"`
	MagicLinks *MagicLinksCmd `arg:"subcommand:magic-links"`
}


func main(){
	arg.MustParse(&args)
	setLogLevel(args.LogLevel)
	var err error
	config, err = ParseConfig()
	if err != nil {
		logger.Fatalf("unable to parse config: %v", err)
	}
	logger.Debugf("got config = %+v", config)

	if config.Token == "" {
		logger.Warnf("no token specified")
	}

	c, err := client.New(config.Token)
	if err != nil {
		logger.Fatalf("unable to get client: %v", err)
	}

	if args.MagicLinks != nil {
		doMagicLinks(c)
	} else {
		logger.Fatalf("please specify a subcommand")
	}

}

func setLogLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	}
}
