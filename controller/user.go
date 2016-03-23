package controller

import (
   "fmt"
   "net/http"
   "text/template"
)


func ForwardUser(rw http.ResponseWriter, r *http.Request) {
    
    url := r.URL
    fmt.Println(url)
    // fmt.Println("Hello World!")
    
    first,_ := template.ParseFiles("view/first.abc")
    
    fmt.Println(first)
    
    // first.ExecuteTemplate(rw, "Firstfield", "Hello World!")
    
    var data map[string]string  
    
    data = make(map[string]string)
    
    data["Firstfield"] = "Hellow WOrld2222"
    
    // json := {Firstfield:"Hello World123"}
    
    first.Execute(rw, data)
}



