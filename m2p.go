package main

import (
	"flag"
	"fmt"
    "github.com/izassan/m2p/convert"
    "github.com/izassan/m2p/entity"
)

func main(){
    m2pio := entity.M2pIo{TmpDir: "./tmp"}

    flag.StringVar(&m2pio.OutputName, "output-name", "", "output file name")
    flag.Parse()
    if flag.NArg() < 2{
        fmt.Println("not enough argument")
        return
    }else if flag.NArg() > 2{
        fmt.Println("too many argument")
        return
    }
    m2pio.Init(flag.Arg(0), flag.Arg(1))

    convert.GeneratePDF(m2pio)

    m2pio.Finish()
}
