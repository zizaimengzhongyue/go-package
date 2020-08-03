package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	Key    string
	Value  string
	GetKey func() string
}

func (this Node) GetValue() string {
	return this.Value
}

func Info() {
	str := `Value 只有三个字段，typ(*rtype) 包含类型信息；ptr(unsafe.Pointer) 包含数据信息；flag 包含 meta 信息"`
	fmt.Println(str)
}

func Zero() {
	var val reflect.Value
	fmt.Printf("与 reflect.Type 不同，reflect.Type 是个 interface，reflect.Value 是个 struct\n")
	fmt.Printf("Zero.IsValid: %t\n", val.IsValid())
	fmt.Printf("Zero.Kind: %s\n", val.Kind().String())
}

func Addr() {
	fmt.Printf("Addr 似乎可以理解为通过 Elem() 方法返回的元素都是可寻址的\n")
	node := Node{Key: "key1", Value: "value1"}
	val := reflect.ValueOf(&node)
	fmt.Printf("Struct.Element.CanAddr: %t\n", val.Elem().CanAddr())
	fmt.Printf("Struct.Element.Addr: %+v\n", val.Elem().Addr())
	slice := []int{1, 2, 3}
	val = reflect.ValueOf(&slice)
	fmt.Printf("Slice.Element.CanAddr: %t\n", val.Elem().CanAddr())
	fmt.Printf("Slice.Element.Addr: %+v\n", val.Elem().Addr())
}

func Bool() {
	fmt.Printf("只有 bool 类型的值才可以调用 Bool() 方法\n")
	a := true
	val := reflect.ValueOf(a)
	fmt.Printf("Bool.Bool: %t\n", val.Bool())
}

func Bytes() {
	fmt.Printf("只有 byte 类型的 slice 才可以调用 Bytes() 方法\n")
	a := []byte{1, 2, 3}
	val := reflect.ValueOf(a)
	fmt.Printf("Slice.Bytes: %+v\n", val.Bytes())
}

func CanSet() {
	node := &Node{Key: "key1", Value: "value1"}
	val := reflect.ValueOf(&node.Key).Elem()
	fmt.Printf("Ponter.CanSet: %t\n", val.CanSet())
	val.SetString("key2")
	fmt.Printf("Field.SetString: %+v\n", node)
}

func Call() {
	fmt.Printf("只有 func 类型的 reflect.Value 才能通过 Call 调用，入参为 []reflect.Value，返回值是 []reflect.Value\n")
	a := func(a int, b int) int {
		return a + b
	}
	p1 := 1
	p2 := 2
	val := reflect.ValueOf(a)
	ans := val.Call([]reflect.Value{reflect.ValueOf(p1), reflect.ValueOf(p2)})
	for _, v := range ans {
		fmt.Printf("Call.Response: %v\n", v)
	}
	b := func(n ...int) int {
		ans := 0
		for _, v := range n {
			ans += v
		}
		return ans
	}
	fmt.Printf("只有可变参函数才可以通过 CallSlice 调用\n")
	val = reflect.ValueOf(b)
	params := []int{1, 2, 3}
	slice := val.CallSlice([]reflect.Value{reflect.ValueOf(params)})
	for _, v := range slice {
		fmt.Printf("CallSlice.Response: %v\n", v)
	}
}

func Cap() {
	fmt.Printf("只有 slice、array、chan 的 reflect.Value 才可以通过 Cap() 获取容量大小\n")
	list := []int{1, 2, 3}
	val := reflect.ValueOf(list)
	fmt.Printf("Slice.Cap: %d\n", val.Cap())
	arr := make([]int, 5)
	val = reflect.ValueOf(arr)
	fmt.Printf("Array.Cap: %d\n", val.Cap())
	ch := make(chan int, 8)
	val = reflect.ValueOf(ch)
	fmt.Printf("Chan.Cap: %d\n", val.Cap())
}

func Close() {
	ch := make(chan int)
	val := reflect.ValueOf(ch)
	val.Close()
	fmt.Printf("Chan.Close\n")
}

func Elem() {
	fmt.Printf("reflect.ValueOf(&a).Elem() 和 reflect.ValueOf(a) 两个值不同，两者的差异是前者 CanSet() 为 true, 后者 CanSet() 为 false\n")
	a := 10
	ptr := reflect.ValueOf(&a)
	fmt.Printf("Ptr.Elem: %+v\n", ptr.Elem())
	val := reflect.ValueOf(a)
	ptrVal := ptr.Elem()
	fmt.Printf("Ptr.Elem()==val: %t\n", reflect.DeepEqual(ptrVal, val))
}

func Field() {
	n := Node{Key: "key1", Value: "val1", GetKey: func() string { return "hello" }}
	val := reflect.ValueOf(n)
	fmt.Printf("Struct.Field: %+v\n", val.Field(0))
	fmt.Printf("Struct.FieldbyIndex: %+v\n", val.FieldByIndex([]int{0}))
	fmt.Printf("Struct.FieldByName: %+v\n", val.FieldByName("Key"))
	fmt.Printf("Struct.NumField: %d\n", val.NumField())
	a := func(key string) bool { return key == "GetKey" }
	fmt.Printf("Struct.FieldByNameFunc: %+v\n", val.FieldByNameFunc(a))
}

func Float() {
	a := 1.27
	val := reflect.ValueOf(a)
	fmt.Printf("Float.Float: %f\n", val.Float())
}

func Index() {
	a := []int{1, 2, 3}
	val := reflect.ValueOf(a)
	fmt.Printf("Slice.Index: %d\n", val.Index(1))
}

func Int() {
	a := 9
	val := reflect.ValueOf(a)
	fmt.Printf("Int.Int: %d\n", val.Int())
}

func Interface() {
	n := &Node{Key: "key1", Value: "val1"}
	val := reflect.ValueOf(n)
	fmt.Printf("Struct.CanInterface: %t\n", val.CanInterface())
	fmt.Printf("Struct.Interface: %+v\n", val.Interface())
	v, _ := val.Interface().(*Node)
	fmt.Printf("Struct.Interface.Assert: %+v\n", v)
}

func IsNil() {
	var n *Node
	n = nil
	val := reflect.ValueOf(n)
	fmt.Printf("Struct.IsNil: %t\n", val.IsNil())
}

func IsValid() {
	var n *Node
	n = nil
	val := reflect.ValueOf(n)
	fmt.Printf("Struct.IsValid: %t\n", val.IsValid())
}

func IsZero() {
	var n *Node
	val := reflect.ValueOf(n)
	fmt.Printf("Struct.IsZero: %t\n", val.IsZero())
}

func Kind() {
	a := 7
	val := reflect.ValueOf(a)
	fmt.Printf("Int.Kind.String: %s\n", val.Kind().String())
}

func Len() {
	a := []int{1, 2, 3}
	val := reflect.ValueOf(a)
	fmt.Printf("[]Int.Len: %d\n", val.Len())
}

func Map() {
	list := map[string]string{"key1": "value1", "key2": "value2"}
	val := reflect.ValueOf(list)
	key := "key2"
	k := reflect.ValueOf(key)
	fmt.Printf("Map.MapIndex: %+v\n", val.MapIndex(k))
	fmt.Printf("Map.MapKeys: %+v\n", val.MapKeys())
}

func Method() {
	node := Node{Key: "key", Value: "val"}
	val := reflect.ValueOf(node)
	fmt.Printf("Struct.NumMethod: %d\n", val.NumMethod())
	fmt.Printf("Struct.MethodByName: %+v\n", val.MethodByName("GetValue"))
}

func Overflow() {
	var a int32 = 1
	val := reflect.ValueOf(a)
	fmt.Printf("Int.OverflowInt: %t\n", val.OverflowInt(int64(11231231231231238)))
}

func Pointer() {
	a := []int{1, 2, 3}
	val := reflect.ValueOf(a)
	fmt.Printf("Slice.Pointer: %+v\n", val.Pointer())
}

func Chan() {
	ch := make(chan int, 2)
	ch <- 8
	val := reflect.ValueOf(ch)
	fmt.Printf("Chan.Recv: ")
	fmt.Println(val.Recv())
	val.Send(reflect.ValueOf(10))
	fmt.Printf("Chan.Send: ")
	fmt.Println(val.Recv())
}

func Set() {
	a := 0
	val := reflect.ValueOf(&a)
	val.Elem().Set(reflect.ValueOf(10))
	fmt.Printf("Int.Set: %d\n", a)
	m := map[string]string{"key1": "val1", "key2": "val2"}
	val = reflect.ValueOf(&m)
	val.Elem().SetMapIndex(reflect.ValueOf("key1"), reflect.ValueOf("key3"))
	fmt.Printf("Map.SetMapIndex: %+v\n", m)
}

func Slice() {
	fmt.Printf("SetCap 方法设置的值必须大于等于当前 Slice.len 且小于等于 Slice.cap\n")
	a := []int{1, 2, 3}
	val := reflect.ValueOf(&a)
	fmt.Printf("Slice.Slice: %+v\n", val.Elem().Slice(0, 2))
	val.Elem().SetCap(3)
	val.Elem().SetLen(1)
	fmt.Printf("Slice.Set: %+v\n", a)
}

func String() {
	a := 9
	val := reflect.ValueOf(a)
	fmt.Printf("Int.String: %s\n", val.String())
}

func Type() {
	a := 9
	val := reflect.ValueOf(a)
	fmt.Printf("Int.Type: %+v\n", val.Type())
}

func Append() {
	a := []int{1, 2, 3}
	b := 4
	val := reflect.Append(reflect.ValueOf(a), reflect.ValueOf(b))
	fmt.Printf("reflect.Append: %+v\n", val)
}

func New() {
	node := &Node{Key: "key", Value: "value"}
	typ := reflect.TypeOf(node)
	val := reflect.New(typ)
	fmt.Printf("reflect.New: %+v\n", val)
}

func Convert() {
	a := 9
	var b int64 = 10
	val := reflect.ValueOf(a)
	typ := reflect.TypeOf(b)
	fmt.Printf("Int.Convert: %+v\n", val.Convert(typ))
}

func main() {
	Info()
	Zero()
	Addr()
	Bool()
	Bytes()
	CanSet()
	Call()
	Cap()
	Close()
	Elem()
	Field()
	Float()
	Index()
	Int()
	Interface()
	IsNil()
	IsValid()
	Kind()
	Len()
	Map()
	Method()
	Overflow()
	Pointer()
	Chan()
	Set()
	Slice()
	String()
	Type()
	Append()
	New()
	Convert()
}
