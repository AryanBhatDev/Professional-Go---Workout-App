package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/AryanBhatDev/Professional-Go---Workout-App/internal/app"
	"github.com/AryanBhatDev/Professional-Go---Workout-App/internal/routes"
)

func main() {

	var port int

	flag.IntVar(&port, "port", 8080, "go backend server port")

	flag.Parse()

	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	defer app.DB.Close()

	r := routes.HandleRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	app.Logger.Printf("We are running our app on port:%d\n", port)

	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}
}
