package main
import "fmt"
import "errors"

func main(){
    fmt.Println("Hello World!")
    var intUNum uint16 = 32767 + 32768
    var int16Num int16 = 32767
    var intNum int = 32767
    fmt.Println(intNum)
    fmt.Println(intUNum)
    fmt.Println(int16Num)

    var float64 float64 = 12345678.9
    var float32 float32 = 12345678.9
    fmt.Println(float32)
    fmt.Println(float64)

    printfn("Go")


    fmt.Println("-----------IF ELSE------------")

    var result, remainder, err = div(10,0)
    if err!=nil{
        fmt.Println(err.Error())
    } else if remainder == 0 {
        fmt.Printf("result is %v", result)
    } else {
        fmt.Printf("result is %v and the remainder is %v", result, remainder)
    }

    fmt.Println("-----------SWITCH------------")

    switch {
    case err != nil:
        fmt.Println(err.Error())
    case remainder == 0:
        fmt.Printf("result is %v", result)
    default:
        fmt.Printf("result is %v and the remainder is %v", result, remainder)
    }

    fmt.Println(result, remainder)


    fmt.Println("-----------ARRAY------------")

    var intArr [3]int32 = [3]int32{1,2,3}
    fmt.Println(intArr)

    var intArr2 = [3]int32{1,2,3}
    fmt.Println(intArr2)

    var intArr3 = [...]int32{1,2,3,4,5}
    fmt.Println(intArr3)

    intArr4 := [...]int32{1,2,3,4}
    fmt.Println(intArr4)
    
    fmt.Println("-----------SLICE------------")
    var intSlice []int32 = []int32{1,2}
    fmt.Printf("Before Append: Length --> %v && Capacity --> %v \n", len(intSlice), cap(intSlice))
    intSlice = append(intSlice, 4)
    fmt.Printf("After Append: Length --> %v && Capacity --> %v \n", len(intSlice), cap(intSlice))

    intSlice2 := []int32{5,6}
    intSlice = append(intSlice, intSlice2...)
    fmt.Println(intSlice)
    fmt.Printf("After Append another Slice: Length --> %v && Capacity --> %v \n", len(intSlice), cap(intSlice))

    var intSlice3 []int32 = make([]int32, 7,8)
    fmt.Println(intSlice3)
    fmt.Printf("Length --> %v && Capacity --> %v \n", len(intSlice3), cap(intSlice3))
}

func printfn (printVal string) {
    fmt.Println("Hello" , printVal)
}

func div (num int, den int) (int, int, error) {
    var err error
    if den == 0 {
        errors.New("Can't divide by 0")
        return 0, 0, err
    }

    var res int = num/den
    var rem int = num%den
    return res, rem, err
}
