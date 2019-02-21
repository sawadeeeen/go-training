// go-training/method.go

package main

import(
  "fmt"
)

// 構造体（Calcという型）
type Calc struct {n1,n2 int}

// 普通の関数
func Add(q Calc) int {
  return q.n1 + q.n2
}

// メソッド
func(p Calc) Add() int{
  return p.n1 + p.n2
}

func main(){
  q := Calc{8,6}
  fmt.Println(Add(q))

  p := Calc{8,6}
  fmt.Println(p.Add())
}
