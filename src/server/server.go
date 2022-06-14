package server

import (
	"context"
	"fmt"
	"github.com/mensylisir/kmpp-middleware/src/config"
	"github.com/mensylisir/kmpp-middleware/src/cron"
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/encrypt"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/migrate"
	"github.com/mensylisir/kmpp-middleware/src/router"
	"github.com/mensylisir/kmpp-middleware/src/util/httpx"
	"github.com/spf13/viper"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"syscall"
)
import "os/signal"

type Server struct {
	ConfigFile string
	Version    string
}

type Phase interface {
	Init() error
	PhaseName() string
}

type ServerOption func(*Server)

func SetConfigFile(file string) ServerOption {
	return func(server *Server) {
		server.ConfigFile = file
	}
}

func SetVersion(version string) ServerOption {
	return func(server *Server) {
		server.Version = version
	}
}

func Run(options ...ServerOption) {
	code := 1
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	server := Server{
		ConfigFile: filepath.Join("src", "conf", "app.yaml"),
		Version:    "not specified",
	}

	for _, option := range options {
		option(&server)
	}

	cleanFunc, err := server.initialize()
	if err != nil {
		fmt.Println("server init fail")
		os.Exit(code)
	}
EXIT:
	for {
		sgn := <-signalChannel
		fmt.Println("received signal:", sgn.String())
		switch sgn {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			code = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}
	cleanFunc()
	fmt.Println("server exit")
	os.Exit(code)
}

func Phases() []Phase {
	return []Phase{
		&encrypt.InitEncryptPhase{
			Multilevel: viper.GetStringMap("encrypt.multilevel"),
		},
		&db.InitDBPhase{
			Host:         viper.GetString("db.host"),
			Port:         viper.GetInt("db.port"),
			Name:         viper.GetString("db.name"),
			User:         viper.GetString("db.user"),
			Password:     viper.GetString("db.password"),
			MaxOpenConns: viper.GetInt("db.max_open_conns"),
			MaxIdleConns: viper.GetInt("db.max_idle_conns"),
		},
		&migrate.InitMigrateDBPhase{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetInt("db.port"),
			Name:     viper.GetString("db.name"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
		},
		&cron.InitCronPhase{
			Enable: viper.GetBool("cron.enable"),
		},
	}
}

func (server Server) initialize() (func(), error) {
	fns := Functions{}
	_, cancel := context.WithCancel(context.Background())
	fns.Add(cancel)
	config.Init()
	logger.Init()
	phases := Phases()
	for _, phase := range phases {
		if err := phase.Init(); err != nil {
			logger.Log.Errorf("start phase [%v] failed reason: %s", phase, err.Error())
			return nil, err
		}
		logger.Log.Infof("start phase [%s] success", phase.PhaseName())
	}

	route := router.New(server.Version)
	go func() {
		err := http.ListenAndServe(":6060", nil)
		if err != nil {
			return
		}
	}()
	httpClean := httpx.Init(route)
	fns.Add(httpClean)
	return fns.Ret(), nil
}

type Functions struct {
	List []func()
}

func (fs *Functions) Add(f func()) {
	fs.List = append(fs.List, f)
}

func (fs *Functions) Ret() func() {
	return func() {
		for i := 0; i < len(fs.List); i++ {
			fs.List[i]()
		}
	}
}
