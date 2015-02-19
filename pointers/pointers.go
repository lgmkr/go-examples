package main

import "fmt"

func zeroval(ival int) {
  ival = 0
}

func zeroptr(iptr *int) {
  *iptr = 0
}

func zone(iptr *int) {
  *iptr= 2
}

func main(){
  i := 1
  fmt.Println("initial:", i)  /// -> 1

  zeroval(i)
  fmt.Println("zeroval:", i) /// -> 1

  zeroptr(&i)
  fmt.Println("zeroptr:", i) /// -> 0  this is pointer

  fmt.Println("pointer:", &i) ///  -> pounter in memory

  ni := new(int)
  zone(ni)
  fmt.Println("zone:", *ni) /// -> 2
}
