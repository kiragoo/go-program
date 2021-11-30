package difstruct

import (
	"fmt"
	"reflect"
	"strings"
)

type Inf interface {
	GetName() string
	GetValue() int
}

type SA struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (sa *SA) GetName() string {
	return sa.Name
}

func (sa *SA) GetValue() int {
	return sa.Value
}

type SB struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
	Age   int    `json:"age"`
}

func (sb *SB) GetName() string {
	return sb.Name
}

func (sb *SB) GetValue() int {
	return sb.Value
}

func (sb *SB) GetAge() int {
	return sb.Age
}

func InfFn(inf Inf) {
	arr := strings.Split(reflect.TypeOf(inf).String(), ".")
	length := len(arr)
	if arr[length-1] == "SB" {
		fmt.Println("xx")
		instance := inf.(*SB)
		fmt.Println(instance.GetAge())
	}

}
