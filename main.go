package main

import (
	"errors"
	"fortune/config"
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
	//model.DB.Init()
	//defer model.DB.Close()
	//
	////init redis
	//utils.Init()

	//Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	//Create the Gin engine
	g := gin.New()

	//g.Static("/static", "./dist")

	//g.LoadHTMLGlob("tpl/*")

	//g.LoadHTMLFiles("static/dist/index.html")

	middlewares := []gin.HandlerFunc{}

	router.Load(
		//Cores,
		g,

		// Middlewares,
		middlewares...,
	)

	//Ping the server to make sure the router is working.
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
		/*go func() {
			log.Info(http.ListenAndServeTLS(tlsAddr,pem,key,g).Error())
			log.Infof("Start to listening the incoming requests on https address: %s",tlsAddr)
		}()*/
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
