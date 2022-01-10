package m2p

import(
    "log"
    "os"
    "io"
    "regexp"
    "path/filepath"
    "github.com/urfave/cli/v2"
)


func copy_to_dest(src_path, dest_path string){
    // check exist dest directory
    if _, err := os.Stat(dest_path); err != nil{
        // if not exist, create directory
        err := os.Mkdir(dest_path, 0777)
        if err != nil{
            panic(err)
        }
    }

    // open source file
    src_file, err := os.Open(src_path)
    if err != nil{
        panic(err)
    }
    defer src_file.Close()

    // open destination file
    path_to_dest := filepath.Join(dest_path, filepath.Base(src_path))
    dest_file, err := os.Create(path_to_dest)
    if err != nil{
        panic(err)
    }
    defer dest_file.Close()

    // execute copy process
    _, err = io.Copy(dest_file, src_file)
    if err != nil{
        panic(err)
    }
}

func generate_directories(src_path string) (string, string){
    // generate directory path
    pdf_dir := filepath.Join(src_path, "pdf")
    old_dir := filepath.Join(src_path, "old")

    // check exist pdf directory
    if _, err := os.Stat(pdf_dir); err != nil{
        // if not exist, create directory
        err := os.Mkdir(pdf_dir, 0777)
        if err != nil{
            panic(err)
        }
    }

    // check exist old directory
    if _, err := os.Stat(old_dir); err != nil{
        // if not exist, create directory
        err := os.Mkdir(old_dir, 0777)
        if err != nil{
            panic(err)
        }
    }

    return pdf_dir, old_dir
}

func convert_mediafiles(src_path, dest_path string){
    // define regex
    regex_move_only := regexp.MustCompile(`(?i)(.+\.(jpg|png|gif|mp4|pdf))`)
    regex_zip_pattern := regexp.MustCompile(`(?i)(.+\.(zip))`)

    // read files in src_path
    files, err := os.ReadDir(src_path)
    if err != nil{
        panic(err)
    }

    // predefine variable
    var path_to_file string
    var path_to_directory string
    var path_to_pdf string
    var path_copy_src string

    // generate pdf and old directory
    path_to_dir_pdf, path_to_old := generate_directories(src_path)

    // process all files in src_path
    for _, file := range files{
        // generate file path
        path_to_file = filepath.Join(src_path, file.Name())

        // process directory pattern
        if file.IsDir(){
            // skip old  and pdf directory
            if path_to_file == path_to_old || path_to_file == path_to_dir_pdf{
                continue
            }

            // output log
            print_process_pattern("directory", file.Name())

            // main process
            path_to_pdf = lib.Dir2pdf(path_to_file, path_to_dir_pdf, path_to_old)

            // set path_copy_src
            path_copy_src = path_to_pdf

        // process zip pattern
        }else if regex_zip_pattern.MatchString(file.Name()){
            // output log
            print_process_pattern("zip", file.Name())

            path_to_directory = lib.Zip2dir(path_to_file, path_to_old)
            path_to_pdf = lib.Dir2pdf(path_to_directory, path_to_dir_pdf, path_to_old)

            // set path_copy_src
            path_copy_src = path_to_pdf

        // process other media(jpg, png, gif, mp4) pattern
        }else if regex_move_only.MatchString(file.Name()){
            // output log
            print_process_pattern("other_media", file.Name())

            // set path_copy_src
            path_copy_src = path_to_file

        // process not media pattern(unsupported media pattern)
        }else{
            // output log
            print_process_pattern("not media", file.Name())

            // set path_copy_src
            path_copy_src = ""

        }

        // dest directory exist check
        if _, err := os.Stat(dest_path); err != nil{
            // if not exist, create directory
            err := os.Mkdir(dest_path, 0777)
            if err != nil{
                panic(err)
            }
        }

        // copy to dest_path
        if path_copy_src != ""{
            copy_to_dest(path_copy_src, dest_path)
        }
    }
}

func print_process_pattern(process_pattern, file_path string){
    log.Print(
        "file: ",
        filepath.Base(file_path),
        " ",
        "process_pattern: ",
        process_pattern)
}

func main(){
    // define cli
    var dest_path string
    var src_path string
    app := &cli.App{
        Name: "m2p",
        Usage: "media files convert to pdf. and media files send other directory",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:   "dest_path",
                Aliases: []string{"d"},
                Value:  "./dest",
                Usage:   "set destinaion path",
                Destination: &dest_path,
            },
            &cli.StringFlag{
                Name:   "src_path",
                Aliases: []string{"s"},
                Value:  "./src",
                Usage:   "set source path",
                Destination: &src_path,
            },
        },
        Action: func(c *cli.Context) error {
            convert_mediafiles(src_path, dest_path)
            return nil
        },
    }

    app.Version = "0.1.0"
    err := app.Run(os.Args)
    if err != nil{
        log.Fatal(err)
    }
}
