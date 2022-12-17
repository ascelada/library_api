package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	//fs := http.FileServer(http.Dir("./static"))
	//http.Handle("/", fs)

	//log.Printf("Starting the application, listening on :8080...")
	//log.Fatal(http.ListenAndServe(":8080", nil))

	log.Println("Application running in environment: ",
		os.Getenv("RUNTIME_SETUP"), "and on port: ", os.Getenv("PORT"))
	var router *gin.Engine
	router = gin.Default()
	router.Static("/", "./static")
	router.Run()

}
