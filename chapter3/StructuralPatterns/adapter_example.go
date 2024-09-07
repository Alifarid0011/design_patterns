package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyServer struct {
	Msg string
}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World ######")
}
func main() {
	server := &MyServer{
		Msg: "Hello, World hiiiii",
	}
	AdapterPrint(AMethod(func(a int, b string) (int, string) {
		fmt.Println("AMethod called")
		return a, b
	}))
	AdapterPrint(A{})
	AdapterPrint(B{})
	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type AdapterInterface interface {
	AdapterMethod(a int, b string) (int, string)
}
type AMethod func(a int, b string) (int, string)

func (A AMethod) AdapterMethod(a int, b string) (int, string) {
	return A(a, b)
}

type A struct {
}

func AdapterPrint(AdapterInterface AdapterInterface) {
	fmt.Println(AdapterInterface.AdapterMethod(112321312, "hello world"))
	fmt.Println("_______________________________")
}

func (A A) AdapterMethod(a int, b string) (int, string) {
	fmt.Println("from A struct")
	return a, b
}

type B struct {
}

func (B B) AdapterMethod(a int, b string) (int, string) {
	fmt.Println("from B struct")
	return a, b
}
