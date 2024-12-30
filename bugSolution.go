func main() {

        x := 10

        fmt.Println(x)

        if x > 5 {

                fmt.Println("x is greater than 5")
                defer fmt.Println("This will be printed after the function exits.")
                return
        }

        fmt.Println("This will not be printed")
        defer fmt.Println("This will be printed after the function exits.")
} 