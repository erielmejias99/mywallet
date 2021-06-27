package main

import (
    _ "embed"
    "github.com/wailsapp/wails"
    "os"
    "wallet/currency"
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

    currencyRep := currency.NewRep( db )

    app.Bind(currencyRep)

    err := app.Run()
    if err != nil{
        panic("Error running application" )
    }
}
