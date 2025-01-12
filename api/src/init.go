package src

import (
	"api/database"
	"api/lib/config"
	"api/src/routes"
	"flag"
	"log"
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run()  {

	viper.SetConfigFile("../.env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	migrateFlag := flag.NewFlagSet("migrate", flag.ExitOnError)
	seedFlag := flag.NewFlagSet("seed", flag.ExitOnError)

	if len(os.Args) < 2{
		fmt.Println("expected 'migrate' or 'seed' subcommands")
		os.Exit(1)
	}

	db := config.ConnectDB()
	if db == nil {
		fmt.Println("Database connection failed.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "migrate":
		migrateFlag.Parse(os.Args[2:])
		migrateRun(db)
		os.Exit(1)
	case "seed":
		seedFlag.Parse(os.Args[2:])
		seederRun(db)
		os.Exit(1)
	default:
		route := gin.Default()
		routes.Init(route)
		route.Run(":" + viper.GetString("PRODUCTION_PORT") )
	}
}

func migrateRun(db *gorm.DB)  {
	fmt.Println("Running migrations...")
	database.LoadMigrate(db)
	fmt.Println("Close migrations...")
}

func seederRun(db *gorm.DB)  {
	fmt.Println("Running seeders...")
	database.LoadSeeders(db)
	fmt.Println("Close seeders...")
}
