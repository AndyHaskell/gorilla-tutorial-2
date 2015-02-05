package main

import (
	"fmt"
	"time"
	"os"
	"net/http"

	"github.com/justinas/alice"
	"github.com/gorilla/handlers"
)

func serveHelloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world!")
}

type helloWorldHandler struct{}                                              //1
func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){//2
	fmt.Fprintf(w, "Hello world!")
}

func sloth(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `<body><img src="http://andyhaskell.github.io/Slothful-Soda/images/sloth.jpg" width="240px" height="300px" /></body>`)
}

func sleepConstructor(h http.Handler) http.Handler{                       //1
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){ //2
        fmt.Println("Sloth is sleeping, please wait")
        time.Sleep(3000 * time.Millisecond)
        h.ServeHTTP(w, r)                                                 //3
    })
}

func teaConstructor(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("*drinks hibiscus tea*")
		time.Sleep(500 * time.Millisecond)
		h.ServeHTTP(w, r)
	})
}

func main(){
	helloHandler := http.HandlerFunc(serveHelloWorld)
	http.Handle("/", helloHandler)



	slothHandler := http.HandlerFunc(sloth)

	sleepHandler := http.HandlerFunc(
	  func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Sloth is sleeping, please wait")
		time.Sleep(3000 * time.Millisecond)
		slothHandler.ServeHTTP(w, r)
	})

	http.Handle("/sloths", sleepHandler)
	http.Handle("/sloths2", sleepConstructor(slothHandler))
	http.Handle("/sloths3", alice.New(sleepConstructor, teaConstructor).Then(slothHandler))



	//sleepTeaSlothChain := sleepConstructor(teaConstructor(slothHandler))

	sleepTeaSlothChain :=
	  alice.New(sleepConstructor, teaConstructor).Then(slothHandler)

	teaSleepSlothChain :=
	  alice.New(teaConstructor, sleepConstructor).Then(slothHandler)

	teaTwiceChain :=
	  alice.New(teaConstructor, teaConstructor).Then(slothHandler)

	http.Handle("/sleepTeaSlothChain", sleepTeaSlothChain)
	http.Handle("/teaSleepSlothChain", teaSleepSlothChain)
	http.Handle("/teaTwiceChain", teaTwiceChain)


	muxWithLog := handlers.LoggingHandler(os.Stdout, http.DefaultServeMux) //1
	http.ListenAndServe(":1123", muxWithLog)                               //2
}
