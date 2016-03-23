package controller

import (
    "fmt"
    "net/http"
    _ "strings"
)

func ForwardRoot(rw http.ResponseWriter, r *http.Request)  {
    fmt.Println("Root")
}