package worker

import (
    "bufio"
    "log"
    "strconv"
)

const (
    FOUND protoResult = iota
    NOT_FOUND_YET
    I_DO_IT
    YOU_DO_IT
    END
)

type (
    protoResult int

    workerPackage struct {
        result protoResult
        pos   int
    }
)

func Work(from []byte, to []byte, res *bufio.Writer) {
    chWorker := make(chan workerPackage)
    chSync := make(chan bool)
    go compare(from, to, "+", chWorker, chSync, res)
    go compare(to, from, "-", chWorker, chSync, res)
    <- chSync
    <- chSync
}

func compare(from []byte, to []byte, operation string, other chan workerPackage, sync chan bool, res *bufio.Writer) {
    hasToWork := true
    posFrom := 0
    posTo := 0
    posToBefore := 0
    for posTo < len(to) {
        select {
        case msg <- other :
            hasToWork = false;
            switch msg.result {
            case FOUND :
//              other <- workerPackage{NOT_FOUND_YET, posTo - posToBefore}
//          case NOT_FOUND_YET :
                if posTo - posToBefore < msg.pos {
//                  other <- workerPackage{I_DO_IT, posTo - posToBefore} §§§ FOUND ok mais reprise boulot
                    size = posToBefore + msg.pos
                    hasToWork = true
                } else {
                    other <- workerPackage{YOU_DO_IT}
                }
            case I_DO_IT :
                posFrom = posFrom + msg.pos
                posTo = posToBefore
                hasToWork = true
            case YOU_DO_IT :
                pos := 0
                if operation == "+" {
                    pos = posFrom
                } else {
                    pos = posTo
                }
                _, err := res.WriteString(operation + " " + strconv.Itoa(pos) + " ")
                if err != nil {
                    log.Fatalln(err)
                }
                _, err = res.Write() // write le text
                if err != nil {
                    log.Fatalln(err)
                }
            case END :
            }
        default :
            if hasToWork == true {
                posTo++
                if from[posFrom] == to[posTo] {
                    // search here
                }
            }
        }
    }
}
