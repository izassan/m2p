package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "github.com/izassan/m2p/convert"
)

func checkFileType(input string) string{
    if filepath.Ext(input) == ".zip"{ return "zip" }
    if fInfo, _ := os.Stat(input); fInfo.IsDir() { return "media" }
    return "other"
}


func main(){
    var inputFile string
    var outputDir string
    var naming string

    flag.StringVar(&naming, "naming", "%s_%d", "naming")
    flag.Parse()
    if flag.NArg() < 2{
        fmt.Println("not enough argument")
        return
    }else if flag.NArg() > 2{
        fmt.Println("too many argument")
        return
    }
    inputFile, outputDir = flag.Arg(0), flag.Arg(1)

    filetype := checkFileType(inputFile)
    if filetype == "other"{
        fmt.Println("unsupport format")
    }
    // check exist output  directory
    if _, err := os.Stat(outputDir); err != nil{
        // if not exist, create directory
        err := os.Mkdir(outputDir, 0777)
        if err != nil{
            panic(err)
        }
    }

    tmpDir := "./tmp"
    if _, err := os.Stat(tmpDir); err != nil{
        // if not exist, create directory
        err := os.Mkdir(tmpDir, 0777)
        if err != nil{
            panic(err)
        }
    }

    if filetype == "zip"{
        extDir := convert.Zip2dir(inputFile, tmpDir)
        convert.Dir2pdf(extDir, outputDir)
    }else if filetype == "media"{
        convert.Dir2pdf(inputFile, tmpDir)
    }
}
