package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("/etc/app/share/settings")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	log.Println(viper.GetString("APP_NAME"))
	log.Println(viper.GetString("APP_VERSION"))

	// if err := godotenv.Load(".env", ".env2"); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// appName := os.Getenv("APP_NAME")
	// log.Println(appName)
	// appVersion := os.Getenv("APP_VERSION")
	// log.Println(appVersion)
}
