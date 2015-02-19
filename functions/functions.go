package main

import "fmt"

func plus(a int, b int) int {

    return a + b
}

func vals() (int, int) {
  return 3, 7
}

func sum(nums ...int) { // arguments 
  fmt.Print(nums, " ")
  total := 0
  for _, num :=range nums {
    total += num
  }
  fmt.Println(total)
}

func intSeq() func() int { // closure
  i := 0
  return func() int {
    i += 1
    return i
  }
}

func fact(n int) int { //recursion
  if n==0 {
    return 1
  }
  return n * fact(n-1)
}

func main() {
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    fmt.Println(c)

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...)

    nextInt := intSeq()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())

    fmt.Println(fact(7))
}
