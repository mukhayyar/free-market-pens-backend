package main

import "backend/db"

func main() {
	db.Init()

	e := Init()
	
	e.Logger.Fatal(e.Start(":1323"))
}
