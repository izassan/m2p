package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "github.com/izassan/m2p/convert"
)

func checkFileType(input string) string{
    imageExts := []string{".jpg", "jpeg", ".png"}
    if filepath.Ext(input) == ".zip"{
        return "zip"
    }
    for _, imageExt := range imageExts{
        if filepath.Ext(input) == imageExt{
            return "media"
        }
    }
    return "other"
}


func main(){
    var inputDir string
    var outputDir string
    var naming string

    flag.StringVar(&naming, "naming", "%s_%d.jpg", "naming")
    flag.Parse()

    if flag.NArg() < 2{
        fmt.Println("not enough argument")
        return
    }else if flag.NArg() > 2{
        fmt.Println("too many argument")
        return
    }

    inputDir, outputDir = flag.Arg(0), flag.Arg(1)

    // dummy output
    fmt.Println(inputDir, outputDir)
}
