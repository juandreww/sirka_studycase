package main

import (
	"fmt"
	"net/http"
	"github.com/juandreww/sirka_studycase/controller"
	// "sirka/model"
)

func main() {
	mux := controller.RegisterApi()
	// db := model.ConnectDB()
	// defer db.Close()

	fmt.Println("Serving...");
	// for heroku
	// port := os.Getenv("PORT")
	// log.Fatal(http.ListenAndServe(":"+port, mux))

	// for Local
	http.ListenAndServe(":80", mux)
}
