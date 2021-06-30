package main

import (
	_ "embed"
	"github.com/wailsapp/wails"
	"os"
    "wallet/controllers"
    "wallet/database"
)


func main() {

    db := database.GetDB()

    if os.Getenv("MIGRATE" ) == "true"{
        err := InitMigrations( db )
        if err != nil{
            panic("Error creating migrations" )
        }
    }

    if os.Getenv("RUNAPP") == "false"{
        return
    }

    app := wails.CreateApp(&wails.AppConfig{
        Width:  1024,
        Height: 500,
        Title:  "wallet",
        Colour: "#131313",
    })

    currencyController := controllers.CurrencyController{Db: db}
    transactionController := controllers.TransactionController{Db: db}
    reasonController := controllers.ReasonController{Db: db}

    app.Bind(&currencyController)
    app.Bind(&transactionController)
    app.Bind(&reasonController)

    err := app.Run()
    if err != nil{
        panic("Error running application" )
    }
}
