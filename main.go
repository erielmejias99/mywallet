package main

import (
	_ "embed"
	"github.com/wailsapp/wails"
	"os"
    "wallet/controllers"
    "wallet/database"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

    db := database.GetDB()

    //if os.Getenv("MIGRATE" ) == "true"{
    err := InitMigrations( db )
    if err != nil{
        panic("Error creating migrations" )
    }
    //}

    if os.Getenv("RUNAPP") == "false"{
        return
    }

    currencyController := controllers.CurrencyController{Db: db}
    transactionController := controllers.TransactionController{Db: db}

    app := wails.CreateApp(&wails.AppConfig{
        Width:  1024,
        Height: 500,
        JS: js,
        CSS: css,
        Title:  "wallet",
        Colour: "#131313",
    })


    app.Bind(&currencyController)
    app.Bind(&transactionController)

    err = app.Run()
    if err != nil{
        panic("Error running application" )
    }
}
