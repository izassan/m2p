package convert

import (
    "os"
    "strings"
    _ "image/jpeg"
    _ "image/png"
    "path/filepath"
    "github.com/izassan/m2p/entity"
)

func GeneratePDF(m2pio entity.M2pIo){
    if m2pio.InputType == "zip"{
        extDir := Zip2dir(m2pio.Input, m2pio.TmpDir)
        Dir2pdf(extDir, m2pio.OutputDir)
    }else if m2pio.InputType == "media"{
        Dir2pdf(m2pio.Input, m2pio.TmpDir)
    }
}

func Dir2pdf(dir_path, path_to_dir_pdf string){

    // generate path to pdf
    pdf_path := filepath.Join(
        path_to_dir_pdf,filepath.Base(dir_path) + ".pdf")

    // read files and sort
    files, err := os.ReadDir(dir_path)
    files = sortdir(files)
    if err != nil {
        panic(err)
    }

    generate_pdf(dir_path, pdf_path, files)
}

func Zip2dir(zip_path, tmp_dir string) string{
    dir_name := strings.Replace(zip_path, ".zip", "", -1)
    unzip(zip_path, dir_name)
    path_to_new := filepath.Join(tmp_dir, filepath.Base(dir_name))
    err := os.Rename(dir_name, path_to_new)
    if err != nil{
        panic(err)
    }
    return path_to_new
}
