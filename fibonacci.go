package main

func fib(x int) int {
  if x == 1 || x == 0 {
    return 1
  }
  return fib(x-1) + fib(x-2)
}
