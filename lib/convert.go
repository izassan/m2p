package lib

import (
	"os"
    "strings"
    _ "image/jpeg"
    _ "image/png"
	"path/filepath"
)

func Dir2pdf(dir_path, path_to_dir_pdf, path_to_dir_old string) string{

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

    // move to old directory
    path_to_old := filepath.Join(path_to_dir_old, filepath.Base(dir_path))
    err = os.Rename(dir_path, path_to_old)
    if err != nil{
        panic(err)
    }

    // return pdf path string
    return pdf_path
}

func Zip2dir(zip_path, path_to_dir_old string) string{
    dir_name := strings.Replace(zip_path, ".zip", "", -1)
    unzip(zip_path, dir_name)
    path_to_old := filepath.Join(path_to_dir_old, filepath.Base(zip_path))
    err := os.Rename(zip_path, path_to_old)
    if err != nil{
        panic(err)
    }
    return dir_name
}
