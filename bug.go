func main() {

        x := 10

        fmt.Println(x)

        defer fmt.Println("This will be printed after the function exits.")

        if x > 5 {

                fmt.Println("x is greater than 5")
                return
        }

        fmt.Println("This will not be printed")
}