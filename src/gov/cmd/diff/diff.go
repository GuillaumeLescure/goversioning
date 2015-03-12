package diff

import (
    "bufio"
    "goUpdater/flags"
    "goUpdater/worker"
    "io/ioutil"
    "log"
    "os"
)

type self struct {
    source []byte
    target []byte
    fPatch *os.File
    wPatch *bufio.Writer
}

var (
    this self
)

func main() {
    if flags.Check() == false {
        os.Exit(1)
    }

    openAll()
    worker.Work(this.source, this.target, this.wPatch)
    closeAll()
}

func openAll() {
    file, err := os.Open(flags.PathSource())
    if err != nil {
        log.Fatalln(err)
    }
    this.source, err = ioutil.ReadAll(bufio.NewReader(file))
    err = file.Close()
    if err != nil {
        log.Fatalln(err)
    }

    file, err = os.Open(flags.PathTarget())
    if err != nil {
        log.Fatalln(err)
    }
    this.target, err = ioutil.ReadAll(bufio.NewReader(file))
    err = file.Close()
    if err != nil {
        log.Fatalln(err)
    }

    this.fPatch, err = os.OpenFile(flags.PathPatch(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
    if err != nil {
        log.Fatalln(err)
    }
    this.wPatch = bufio.NewWriter(this.fPatch)
}

func closeAll() {
    err := this.fPatch.Close()
    if err != nil {
        log.Fatalln(err)
    }
}
