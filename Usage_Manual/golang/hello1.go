
package main

import "fmt"

// go build hello.go

func main(){
    fmt.Println("Welcome to use Golang")
    // test11()
    // test21()
    // test22()
    // test31()
    // testArray11()
    testArray21()
}

func test11(){
    fmt.Println("hello world")
    fmt.Println("gday")
}
func test21(){
    var str_welcome string = "Welcome to learn some math"
    var a int = 1
    var b int = 2
    var c,d int
    c = 3
    d = 4
    var result int = a + b

    fmt.Println(str_welcome)
    fmt.Println("a + b =", a+b)
    fmt.Println("a + b =", result)
    fmt.Println("c + d =", c+d)
}
func test22(){
    str_welcome := "Welcome to learn some math"
    a := 3
    b := 4
    var c,d int
    result := a+b

    fmt.Println(str_welcome)
    fmt.Println("a + b =", a+b)
    fmt.Println("a + b =", result)
    fmt.Println("c + d =", c+d)
}
func test31(){
    x := 10
    y := 11
    z := 12
    total := plusPlus(x,y,z)
    fmt.Println("x + y + z =", total)
}
func plusPlus(a int, b int, c int) int{
    return a + b + c
}
func testArray11(){
    var arr [4] int
    fmt.Println(arr)
    arr[0] = 1
    arr[1] = 2
    fmt.Println(arr)

    // arrN := [4] int {10,11,12,13}
    arrN := [...] int {10,11,12,13}
    fmt.Println(arrN)
    arrN[0] = 0
    arrN[1] = 1
    fmt.Println(arrN)

    for i := 0; i < len(arrN); i++ {
        arrN[i] = arrN[i] + 10
    }
    fmt.Println(arrN)

}

func testArray21(){
    var arr [4][3] int
    fmt.Println(arr)
    for i := 0; i < 4; i++ {
        for j := 0; j < 3; j++ {
            arr[i][j] = i * 10 + j
        }
    }
    fmt.Println(arr)
}


