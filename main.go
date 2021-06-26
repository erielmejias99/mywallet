package main

import (
  _ "embed"
  "github.com/wailsapp/wails"
  "wallet/currency"
  "wallet/database"
    "wallet/util"
)

func basic() string {
  return "Hello World!"
}

func main() {

  db := database.GetDB()
  err := db.AutoMigrate( &currency.Currency{} )

  currencyRep := currency.NewRep( db )

  app := wails.CreateApp(&wails.AppConfig{
   Width:  1024,
   Height: 768,
   Title:  "wallet",
   Colour: "#131313",
  })

  app.Bind(basic)
  app.Bind(currencyRep)
  app.Bind(util.Login)

  err = app.Run()
  if err != nil{
   panic("Error running application" )
  }
}
