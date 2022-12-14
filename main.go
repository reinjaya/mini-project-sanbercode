package main

import (
	"database/sql"
	"fmt"
	"os"
	"rein/tugas16/controllers"
	"rein/tugas16/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "password"
// 	dbname   = "practice"
// )

var (
	DB  *sql.DB
	err error
)

func main() {
	//psqlInfo := fmt.Sprintf("host=#{host} port=#{port} dbname=#{dbname} user=#{user} password=#{password} sslmode=disable")
	//psqlInfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)

	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed to load config")
	} else {
		fmt.Println("loaded config")
	}
	//psqlInfo := fmt.Sprintf("host=#{os.Getenv("DB_HOST")} port=#{os.Getenv("DB_PORT")} dbname=#{os.Getenv("DB_NAME")} user=#{os.Getenv("DB_USER")} password=#{os.Getenv("DB_PASSWORD")} sslmode=disable")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection established")
	}
	database.DbMigrate(DB)
	defer DB.Close()

	router := gin.Default()
	router.GET("/person", controllers.GetAllPerson)
	router.POST("/person", controllers.InsertPerson)
	router.PUT("/person/:id", controllers.UpdatePerson)
	router.DELETE("/person/:id", controllers.DeletePerson)

	router.Run(":" + os.Getenv("PORT"))
}
