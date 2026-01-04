package main

import (
	"context"
	"errors"
	"fortune/config"
	"fortune/db"
	"fortune/handler/color"
	"fortune/pkg/log"
	"fortune/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var cfg = pflag.StringP("config", "c", "", "apiserver config file path.")

func main() {
	pflag.Parse()

	//init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init DB
	mysqlDB, err := db.NewMysqlDB()
	if err != nil {
		log.Errorf("NewMysqlDB error:%v", err)
		panic(err)
	}
	err = color.NewColorHandler(context.Background(), mysqlDB)
	if err != nil {
		log.Errorf("NewColorHandler error:%v", err)
		panic(err)
	}

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine
	g := gin.Default()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	tlsAddr := viper.GetString("tls.addr")
	pem := viper.GetString("tls.pem")
	key := viper.GetString("tls.key")

	if pem != "" && key != "" {
		go func() {
			log.Infof(g.RunTLS(tlsAddr, pem, key).Error())
		}()
	}

	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
