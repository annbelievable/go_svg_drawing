package main

import (
	"log"
	"net/http"

	svg "github.com/ajstarks/svgo"
)

func main() {
	http.Handle("/circle", http.HandlerFunc(circle))
	http.Handle("/smileyFace", http.HandlerFunc(smileyFace))
	err := http.ListenAndServe(":1999", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:none;stroke:black")
	s.End()
}

func smileyFace(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	height := 1000
	width := 1000
	s.Start(width, height)
	//face
	s.Circle(width/2, height/2, 400, "fill:yellow;stroke:black")
	//left eye
	s.Ellipse(300, 400, 20, 80, "fill:black;stroke:black")
	//right eye
	s.Ellipse(700, 400, 20, 80, "fill:black;stroke:black")
	//smile
	s.Arc(300, 550, 50, 50, 50, true, false, 700, 550, "fill:none;stroke:black")
	s.End()
}
