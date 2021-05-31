package main

import (
	"fmt"
	"math"
	"reflect"
)

const (
	Integer1 = 1000
	Integer2 = math.MaxUint64 + 1
	Float1   = 1.23
	Float2   = 1e100
	Float3   = 1e400
)

func main() {
	fmt.Println("integer1=", Integer1, "type", reflect.TypeOf(Integer1).Name())
	// 编译错误
	// fmt.Println("integer2=", Integer2, "type", reflect.TypeOf(Integer2).Name())
	fmt.Println("integer2/10=", Integer2/10, "type", reflect.TypeOf(Integer2/10).Name())

	fmt.Println("float1=", Float1, "type", reflect.TypeOf(Float1).Name())
	fmt.Println("float2=", Float2, "type", reflect.TypeOf(Float2).Name())
	// 编译错误
	// fmt.Println("float3=", Float3, "type", reflect.TypeOf(Float3).Name())
	fmt.Println("float3/float2=", Float3/Float2, "type", reflect.TypeOf(Float3/Float2).Name())
}
