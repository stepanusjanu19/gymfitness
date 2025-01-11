package src

import (
	"backend/database"
	"backend/lib/config"
	"backend/src/routes"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Run()  {

	migrateFlag := flag.NewFlagSet("migrate", flag.ExitOnError)
	seedFlag := flag.NewFlagSet("seed", flag.ExitOnError)

	if len(os.Args) < 2{
		fmt.Println("expected 'migrate' or 'seed' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "migrate":
		migrateFlag.Parse(os.Args[2:])
		migrateRun()
		os.Exit(1)
	case "seed":
		seedFlag.Parse(os.Args[2:])
		seederRun()
		os.Exit(1)
	default:
		config.ConnectDB()
		route := gin.Default()
		routes.Init(route)
		route.Run(":5001")
	}
}

func migrateRun()  {
	fmt.Println("Running migrations...")
	db := config.ConnectDB()
	database.LoadMigrate(db)
	fmt.Println("Close migrations...")
}

func seederRun()  {
	fmt.Println("Running seeders...")
	db := config.ConnectDB()
	database.LoadSeeders(db)
	fmt.Println("Close seeders...")
}
