package utils

import (
    "os"
    "path/filepath"
)

func RootPath()  string  {
    rootpath, _ := filepath.Abs(os.Args[0])
    return rootpath
}

// func ViewPath() string {
//     viewpath := RootPath() + "/view"
//     return viewpath
// }