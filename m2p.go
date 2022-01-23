package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "github.com/izassan/m2p/convert"
)

type m2pIo struct{
    input string
    outputDir string
    outputName string
}

func checkFileType(input string) string{
    if filepath.Ext(input) == ".zip"{ return "zip" }
    if fInfo, _ := os.Stat(input); fInfo.IsDir() { return "media" }
    return "other"
}


func main(){
    m2pio := m2pIo{}

    flag.StringVar(&m2pio.outputName, "output-name", "", "output file name")
    flag.Parse()
    if flag.NArg() < 2{
        fmt.Println("not enough argument")
        return
    }else if flag.NArg() > 2{
        fmt.Println("too many argument")
        return
    }
    m2pio.input, m2pio.outputDir = flag.Arg(0), flag.Arg(1)

    filetype := checkFileType(m2pio.input)
    if filetype == "other"{
        fmt.Println("unsupport format")
    }
    // check exist output  directory
    if _, err := os.Stat(m2pio.outputDir); err != nil{
        // if not exist, create directory
        err := os.Mkdir(m2pio.outputDir, 0777)
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
        extDir := convert.Zip2dir(m2pio.input, tmpDir)
        convert.Dir2pdf(extDir, m2pio.outputDir)
    }else if filetype == "media"{
        convert.Dir2pdf(m2pio.input, tmpDir)
    }

    // remove tmp directory
    os.RemoveAll(tmpDir)
}
