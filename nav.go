package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "time"
)

type (Layout [40][]byte
Position [2]int
)

var lastDir string

func readLayout(filename string, layoutP *Layout) {

    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var roomLayoutLine []byte
    scanner := bufio.NewScanner(file)
    for i := 0; scanner.Scan(); i++ {
        roomLayoutLine = []byte(scanner.Text())
        layoutP[i] = roomLayoutLine
    }
}
func printLayout (move int, layout Layout) {
    fmt.Println("This is Move ",move)
    for _,v := range layout {
        fmt.Println(string(v))
    }
}

func insertActor(layout *Layout, pos Position) {
    layout[pos[0]][pos[1]] = '@'
}

func clearBlock(layout Layout, cleared *Layout, actor *Position) {
    if true == moveNorth(layout,actor) {
        //If north is accessible, cover first block
        for i:=0;i<5;i++ {
            if false == moveNorth(layout,actor) {
                break
            }
        }
    } else if start == actor {
    } else {
    }
}

func checkBlock(cleared Layout, start Position, actor Position) bool {
    return true
}

func moveEast(layout *Layout, actor *Position) bool {
    if layout[actor[0]][actor[1]+1] == '=' {
        return false
    } else {
        curr := actor
        layout[curr[0]][curr[1]] = '.'
        actor[1] = actor[1]+1
        layout[actor[0]][actor[1]] = '@'
        return true
    }
}



func main() {

    var roomLayout Layout
    var cleared Layout

    readLayout(os.Args[1], &roomLayout)

    actorPos := Position{1,2}
    insertActor(&roomLayout, actorPos)

    done := false
    startPos := actorPos

    for blocks:=0; done == false; blocks++ {
        clearBlock(roomLayout,&cleared,&actorPos)
        done = checkDone(cleared,startPos,actorPos)
        printLayout(moves, roomLayout, cleared)
        //if (moves == 10) {
        //    done = true
        //}
        time.Sleep(200 * time.Millisecond)
    }
    //fmt.Print(roomLayoutLine)
}
