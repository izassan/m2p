package entity

import (
    "os"
    "path/filepath"
)

type M2pIo struct{
    Input string
    InputType string
    OutputDir string
    OutputName string
    TmpDir string
}

func (i *M2pIo) Init(in, out string){
    i.Input = in
    i.OutputDir = out
    i.InputType = checkFileType(i.Input)

    generateDirectory(i.OutputDir)
    generateDirectory(i.TmpDir)
}

func (i *M2pIo) Finish(){
    // remove tmp directory
    os.RemoveAll(i.TmpDir)
}

func generateDirectory(genDir string){
    if _, err := os.Stat(genDir); err != nil{
        // if not exist, create directory
        err := os.Mkdir(genDir, 0777)
        if err != nil{
            panic(err)
        }
    }
}

func checkFileType(input string) string{
    if filepath.Ext(input) == ".zip"{ return "zip" }
    if fInfo, _ := os.Stat(input); fInfo.IsDir() { return "media" }
    return "other"
}
