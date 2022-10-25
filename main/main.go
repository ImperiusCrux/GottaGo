package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/id/", calledIfIdUrl)
	log.Fatalln(http.ListenAndServe(":8080", nil))
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
