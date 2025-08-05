package src

import (
	"api/database"
	"api/lib/config"
	"api/src/routes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func Run() {

	envPath := os.Getenv("ENV_PATH")
	if envPath == "" {
		// envPath = "/go_project/src/gymfitness/api/.env" //read this env for path in prod server path
		envPath = "./.env"
	}
	viper.SetConfigFile(envPath)

	portDefault := viper.GetString("PRODUCTION_PORT")
	if portDefault == "0" || portDefault == "" {
		portDefault = "5001"
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	migrateFlag := flag.NewFlagSet("migrate", flag.ExitOnError)
	seedFlag := flag.NewFlagSet("seed", flag.ExitOnError)

	if len(os.Args) < 2 {
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
		route.Run(":" + portDefault)
	}
}

func migrateRun(db *gorm.DB) {
	fmt.Println("Running migrations...")
	database.LoadMigrate(db)
	fmt.Println("Close migrations...")
}

func seederRun(db *gorm.DB) {
	fmt.Println("Running seeders...")
	database.LoadSeeders(db)
	fmt.Println("Close seeders...")
}
