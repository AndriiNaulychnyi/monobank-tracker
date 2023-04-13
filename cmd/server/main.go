package main

import (
	"MonobankTracker/pkg/configs"
	"MonobankTracker/pkg/currency"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {
	configs.NewServerConfig() // todo: return error instead of panic
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	dbConnection := connectionToDB()

	currency.RegisterEndpoints(router, dbConnection)
	configs.SetWebhook()
	err := router.Run(":" + viper.GetString("server.port"))
	log.Fatal(err)

}

func connectionToDB() *sql.DB {
	postgresqlDbInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		viper.GetString("postgresql.host"),
		viper.GetString("postgresql.port"),
		viper.GetString("postgresql.user"),
		viper.GetString("postgresql.password"),
		viper.GetString("postgresql.dbname"),
	)
	dbConnection, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}
	//defer dbConnection.Close()
	err = dbConnection.Ping()
	if err != nil {
		panic(err)
	}
	log.Print("Established a successful connection!")
	return dbConnection
}
