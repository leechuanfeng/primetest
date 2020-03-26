package main

import(
  "bufio"
  "fmt"
  "os"
  "math"
  "strconv"
  "strings"
)

func isPrime(n int) bool{
  flag := true
  for i:=2; i <= int(math.Sqrt(float64(n))); i++ {
    if n % i == 0 {
      flag = false
      break
    }
  }
  
  if !flag {
    return false
  }
  return true
}

func main(){
  fmt.Println()
  reader := bufio.NewReader(os.Stdin)
  text := ""

  fmt.Print("Please input a number: ")
  text, _ = reader.ReadString('\n')
  text = strings.Replace(text,"\n", "", -1)
  
  
  n, _ := strconv.ParseInt(text,10,0)
  ans := isPrime(int(n))
  
  if ans {
    fmt.Println("It is a prime number")
  }else{
    fmt.Println("It is not a prime number")
  }
  
}