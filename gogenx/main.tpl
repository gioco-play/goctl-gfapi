package main

import (
	"flag"
	"fmt"
    "log"
	"github.com/joho/godotenv"
    "github.com/spf13/viper"

	{{.importPackages}}
)

var (
    configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")
    envFile    = flag.String("env", "etc/.env", "the env file")
)

func main() {
	flag.Parse()

    err := godotenv.Load(*envFile)
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
    // 配置額外參數
    viper.SetConfigName(".env")
    viper.SetConfigType("env")
    viper.SetConfigFile(*envFile)
    err = viper.ReadInConfig() // Find and read the config file
    if err != nil {            // Handle errors reading the config file
        panic(fmt.Errorf("Fatal error config file: %w \n", err))
    }

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
