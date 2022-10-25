package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", calledIfIdUrl)
	log.Fatalln(http.ListenAndServeTLS(":4443", "\\src\\certfile.pm", "src\\keyfile.pm", nil))
}

func dummyRoutine() {
	go doSomethingFunky()
	go doSomethingElseFunky()
}

func fibbTesting() {
	fibb := Fibb{Id: 001}
	fibb.AnotherFibb()
	//fibb.CalcFibb()
	fmt.Println(fibb.GetFunky())
}

func crashtest() {
	var a int = 5
	b := "lul"
	a = 5
	print(b)
	print(a)
}
