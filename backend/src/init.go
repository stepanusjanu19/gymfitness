package src

import (
	"backend/database"
	"backend/lib/config"
	"backend/src/routes"
	"flag"
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

func Run()  {

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
		route.Run(":5001")
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
