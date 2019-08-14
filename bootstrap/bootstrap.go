package bootstrap

import (
	"context"
	"fmt"
	"golang-echo-layout/config"
	"golang-echo-layout/database"
	"golang-echo-layout/router"
	"golang-echo-layout/utils/log"
	"os"
	"os/signal"
	"time"
)

func StartServer(ctx context.Context) {
	fmt.Println(ctx.Value("env"))
	config.NewConfig(ctx.Value("env").(string))

	//New mysql
	database.NewMysql()
	if ctx.Value("env").(string) != "prod" {
		database.Mysql.LogMode(true)
	}
	defer database.Mysql.Close()

	// Logger init
	log.InitLog(config.Conf.App.LogLevel)

	s := router.Routers()

	//data, err := json.MarshalIndent(s.Routes(), "", "  ")
	//if err != nil {
	//	log.Error(err)
	//}
	//_ = ioutil.WriteFile("routes.json", data, 0644)

	go func() {
		log.Info("Start server at ", config.Conf.App.Addr)
		if err := s.Start(config.Conf.App.Addr); err != nil {
			log.Error("Shut down the server with error ", err)
		}
	}()

	quite := make(chan os.Signal)
	signal.Notify(quite, os.Interrupt)
	<-quite

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
}
