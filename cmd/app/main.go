package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	e "github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"jet.com/infrared/db"
	"jet.com/infrared/logger"
	"jet.com/infrared/middleware"
	"jet.com/infrared/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var (
	configFile string
	port       int
)

func init() {
	flag.StringVar(&configFile, "config", "", "-config ./app.yaml")
	flag.IntVar(&port, "port", 0, "-port 8080")
	flag.IntVar(&port, "profile", 0, "-profile dev | prod")
}

type HttpServerConfig struct {
	Port     uint16
	Profile  string
	Database *db.Config
	Log      *logger.Config
}

type HttpServer struct {
	dB     *sqlx.DB
	serv   *http.Server
	config *HttpServerConfig
	ctx    *context.Context
	cancel context.CancelFunc
	close  *sync.WaitGroup
}

func newHttpServer() *HttpServer {
	return &HttpServer{}
}

func (server *HttpServer) ServerInit() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("DB", server.dB)
		c.Next()
	})
	middleware.Middleware(r)
	router.Router(r)

	server.serv = &http.Server{
		Addr:    fmt.Sprintf(":%d", server.config.Port),
		Handler: r,
	}

	ctx, cancel := context.WithCancel(context.Background())
	server.ctx = &ctx
	server.cancel = cancel
	w := &sync.WaitGroup{}
	w.Add(1)
	server.close = w
}

func (server *HttpServer) Close() {
	if server.dB != nil {
		_ = server.dB.Close()
	}
	if server.serv != nil {
		server.cancel()
	}
	server.close.Wait()
}

func (server *HttpServer) Open() {
	go func() {
		logger.Logger.Printf("server listen on %d", server.config.Port)
		err := server.serv.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				logger.Logger.Printf("server stop")
			} else {
				logger.Logger.Fatal(e.Wrap(err, "start fail"))
			}
		}
	}()
	go func() {
		<-(*server.ctx).Done()
		_ = server.serv.Shutdown(context.Background())
		server.close.Done()
	}()
}

func configInit() *HttpServerConfig {
	config := HttpServerConfig{}

	flag.Parse()
	if configFile == "" {
		log.Printf("[warm] config-file no set, use default file ./app.yaml")
		configFile = "./app.yaml"
	}

	f, err := os.Open(configFile)
	if err != nil {
		flag.PrintDefaults()
		log.Fatalf("config file not exist")
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	checkErr(err, "file decode fail")

	return &config
}

func main() {
	server := newHttpServer()

	config := configInit()
	server.config = config
	logger.Init(config.Log, config.Profile)

	newDb, err := db.NewDb(config.Database)
	checkErr(err, "数据库打开失败")
	server.dB = newDb

	server.ServerInit()
	server.Open()
	defer server.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)
	<-stop
}

func checkErr(err error, args ...string) {
	if err != nil {
		panic(e.Wrapf(err, "%s", args))
	}
}
