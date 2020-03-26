package main

import(
  "testing"
  "fmt"
)

func TestForPrime(t *testing.T){
  fmt.Println("testForPrime running...")
  testCases := []struct{
    n int
  }{
    {
      n: 7,
    },
    {
      n: 11,
    },
    {
      n: 83,
    },
    {
      n: -3,
    },
  }
  t.Run("prime", func(t *testing.T) {
    for _, tt := range testCases{
      ans := isPrime(tt.n)
      if !ans {
        t.Errorf("%d should be a prime!", tt.n)
      }
    }
  })
}

func TestForNonPrime(t *testing.T){
  fmt.Println("TestForNonPrime running...")
  testCases := []struct{
    n int
  }{
    {
      n: 6,
    },
    {
      n: -8,
    },
    {
      n: 55,
    },
  }
  t.Run("nonprime", func(t *testing.T) {
    for _, tt := range testCases{
      ans := isPrime(tt.n)
      if ans {
        t.Errorf("%d should be a prime!", tt.n)
      }
    }
  })
}