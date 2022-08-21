package models

import (
	"fmt"
	"testing"
	"unsafe"
)

// Go 内存对齐，推荐将相同类型的字段放在一起，使得占用内存更小

type s1 struct {
	a int8
	b string
	c int8
}

type s2 struct {
	a int8
	c int8
	b string
}

func TestStruct(t *testing.T) {
	v1 := s1{
		a: 1,
		b: "Franky",
		c: 2,
	}

	v2 := s2{
		a: 1,
		b: "Franky",
		c: 2,
	}

	fmt.Println(unsafe.Sizeof(v1), unsafe.Sizeof(v2)) // 32,24
}
