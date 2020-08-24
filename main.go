package main

import "log"

import "./src/router"

func main() {
  app := router.Routes()

  log.Fatal(app.Listen(3000))
}