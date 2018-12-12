package main

import (
	"github.comcast.com/lfordy200/k8_test/webservice"
	"log"
	"net/http"
)

func main() {

	//rectangle := shapeutil.Rectangle{Width: 10.0, Height: 20.0}
	//fmt.Printf("Rectangle perimeter: %.2f", shapeutil.Perimeter(rectangle))

	router := webservice.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
