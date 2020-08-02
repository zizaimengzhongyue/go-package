package main

import (
	"fmt"
	"reflect"
)

type KeyItem interface {
	GetKey() string
}

type Item struct {
	Key   string
	Value string
}

type Node struct {
	Key   string
	Value string
	It    Item
}

func (this Node) GetKey() string {
	return this.Key
}

func (this Node) GetValue() string {
	return this.Value
}

func Int() {
	a := 9
	typ := reflect.TypeOf(a)
	fmt.Printf("Bits.Int: %d\n", typ.Bits())
	typ = reflect.TypeOf(uint8(a))
	fmt.Printf("Bits.Uint8: %d\n", typ.Bits())
	b := 1.23423
	typ = reflect.TypeOf(b)
	fmt.Printf("Bits.Float: %d\n", typ.Bits())
}

func Array() {
	arr := [5]int{1, 2, 3, 4, 5}
	typ := reflect.TypeOf(arr)
	fmt.Printf("Array.Elem: %+v\n", typ.Elem())
	fmt.Printf("Array.Len: %d\n", typ.Len())
}

func Chan() {
	ch := make(chan int)
	typ := reflect.TypeOf(ch)
	fmt.Printf("Chan.ChanDir: %+v\n", typ.ChanDir())
	fmt.Printf("Chan.Elem: %+v\n", typ.Elem())
	ch1 := make(chan<- int)
	typ = reflect.TypeOf(ch1)
	fmt.Printf("Chan<-.ChanDir: %+v\n", typ.ChanDir())
	fmt.Printf("Chan<-.ChanDir: %+v\n", typ.Elem())
}

func Func() {
	a := func(a int, b int) int {
		return a + b
	}
	typ := reflect.TypeOf(a)
	fmt.Printf("Func.In: %+v\n", typ.In(0))
	fmt.Printf("Func.NumIn: %d\n", typ.NumIn())
	fmt.Printf("Func,Out: %+v\n", typ.Out(0))
	fmt.Printf("Func.NumOut: %d\n", typ.NumOut())
	fmt.Printf("Func.IsVariadic: %t\n", typ.IsVariadic())
	b := func(a ...int) int {
		sum := 0
		for _, v := range a {
			sum += v
		}
		return sum
	}
	typ = reflect.TypeOf(b)
	fmt.Printf("Func.IsVariabdic: %t\n", typ.IsVariadic())
}

func Map() {
	a := map[string]int{"key1": 1, "key2": 2}
	typ := reflect.TypeOf(a)
	fmt.Printf("Map.Key: %+v\n", typ.Key())
	fmt.Printf("Map.Elem: %+v\n", typ.Elem())
}

func Ptr() {
	a := 9
	typ := reflect.TypeOf(&a)
	fmt.Printf("Ptr.Elem: %+v\n", typ.Elem())
}

func Slice() {
	a := []int{1, 2}
	typ := reflect.TypeOf(a)
	fmt.Printf("Slice.Elem: %+v\n", typ.Elem())
}

func Struct() {
	a := Node{Key: "key1", Value: "value1", It: Item{Key: "key2", Value: "value2"}}
	typ := reflect.TypeOf(a)
	fmt.Printf("Struct.NumField: %d\n", typ.NumField())
	fmt.Printf("Struct.Field: %+v\n", typ.Field(0))
	fmt.Printf("Struct.FieldByIndex: %+v\n", typ.FieldByIndex([]int{2}))
	sf, ok := typ.FieldByName("Key")
	fmt.Printf("Struct.FieldByName: %+v, %t\n", sf, ok)
}

func main() {
	n := Node{Key: "k1", Value: "v1"}
	typ := reflect.TypeOf(n)
	fmt.Printf("%+v\n", typ)
	fmt.Printf("Align: %d\n", typ.Align())
	fmt.Printf("Method 0: %+v\n", typ.Method(0))
	fmt.Printf("Method 1: %+v\n", typ.Method(1))
	m, bol := typ.MethodByName("GetKey")
	fmt.Printf("Method GetKey: %+v, %t\n", m, bol)
	fmt.Printf("Method Num: %d\n", typ.NumMethod())
	fmt.Printf("Type Name: %s\n", typ.Name())
	fmt.Printf("Pkg Path: %s\n", typ.PkgPath())
	fmt.Printf("Size: %d\n", typ.Size())
	fmt.Printf("String: %s\n", typ.String())
	fmt.Printf("    String 不保证唯一\n")
	fmt.Printf("Kind: %d\n", typ.Kind())
	fmt.Printf("    Kind 是枚举值，枚举了基础类型\n")
	tType := reflect.TypeOf((*KeyItem)(nil)).Elem()
	fmt.Printf("Implements type interface KeyItem: %t\n", typ.Implements(tType))
	fmt.Printf("AssignableTo type interface KeyItem: %t\n", typ.AssignableTo(tType))
	fmt.Printf("type Node struct ConvertibleTo type interface KeyItem: %t\n", typ.ConvertibleTo(tType))
	fmt.Printf("type interface KeyItem ConvertibleTo type Node struct: %t\n", tType.ConvertibleTo(typ))
	fmt.Printf("Comparable: %t\n", typ.Comparable())
	fmt.Printf("下面是针对反射类型的特定方法: \n")
	Int()
	Array()
	Chan()
	Func()
	Map()
	Ptr()
	Slice()
	Struct()
}
