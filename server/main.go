package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"github.com/sharpvik/log-go/v2"
	"net/http"
	"os"
	"runtime/debug"
)

type flagStore struct {
	help    *bool
	version *bool
	config  *string
}

var flag flagStore
var config Config

func init() {
	log.SetLevel(log.LevelDebug)
	flag.help = getopt.BoolLong("help", 'h', "Display help message")
	flag.version = getopt.BoolLong(
		"version", 'v', "Display current version of luba server")
	flag.config = getopt.StringLong("config", 'c', "", "Specify path to custom config file")
	getopt.Parse()
	useFlags(flag)
}

func useFlags(flag flagStore) {
	if *flag.help {
		getopt.Usage()
		os.Exit(0)
	} else if *flag.version {
		prog, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("🤔 luba server version unknown")
		} else {
			fmt.Printf("👸 luba server %s @ %s\n",
				prog.Main.Version, prog.Main.Path)
		}
		os.Exit(0)
	}
}

func configure() {
	if *flag.config != "" {
		config = MustConfig(*flag.config)
	} else {
		log.Info("luba will use the default config (try --config to customize)")
		config = DefaultConfig
	}
}

func main() {
	configure()
	config.Log()
	if err := NewServer().ListenAndServe(); err != http.ErrServerClosed {
		log.Errorf("server shut with error: %s", err)
	}
}
