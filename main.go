package main

import (
	"fmt"
	"goweb/controller"
    "goweb/utils"
	"net/http"
    "strings"
    "text/template"
)

func main() {
    
   
    http.HandleFunc("/",forward)
    // http.HandleFunc("/favicon.ico",favicon)
    
    // fmt.Println("Server Start!")
	http.ListenAndServe("localhost:80", nil)
    
    first,_ := template.ParseFiles("view/first.tpl")
    
    // viewpath := utils.ViewPath()
    
    // fmt.Println(viewpath)
    
    fmt.Println(first)
    
    fmt.Println(utils.ViewPath)
    
     a := "abc" 
    b := "abc" 
    
     fmt.Println(a==b)

}

func forward(rw http.ResponseWriter, r *http.Request) {
    
    url := r.URL
    
    path := url.Path
    
    if strings.HasPrefix("/user",path) {
        controller.ForwardUser(rw, r)
    } else {
        controller.ForwardRoot(rw, r)
    }
   
}

func favicon(rw http.ResponseWriter, r *http.Request)  {
    // fmt.Println("Icon")
}

