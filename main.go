package main

import (
	"github.com/spf13/viper"
	"log"
	"mygram/database"
	"mygram/routers"
)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	//if err := viper.ReadInConfig(); err != nil {
	//	log.Fatalln("Error loading env file:", err)
	//}

	log.Println("Env successfully loaded")
}

func main() {
	database.StartDB()
	routers.StartApp().Run(":" + viper.GetString("PORT"))
}
