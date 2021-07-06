package main

import (
  _ "embed"
  "github.com/wailsapp/wails"
  "os"
  "wallet/controllers"
  "wallet/database"
)

func basic() string {
  return "Hello World!"
}

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

  //currencies, _ :=currencyController.GetAll()
  //cur := currencies[ 0 ]
  //cur.USDChange = 10.2
  //currencyController.Update(map[string]interface{}{
  //  "usd_change": cur.USDChange,
  //  "balance": cur.USDChange,
  //  "ID": cur.ID,
  //})

  app := wails.CreateApp(&wails.AppConfig{
    Width: 1000,
    Title:  "wallet",
    JS:     js,
    CSS:    css,
    Resizable: true,
    Colour: "#131313",
    MinWidth: 500,
    MinHeight: 400,
  })
  app.Bind(basic)
  app.Bind(&currencyController)
  app.Bind(&transactionController)
  app.Run()
}
