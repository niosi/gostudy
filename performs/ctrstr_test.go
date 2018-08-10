package main

import (
	"testing"
	"bytes"
	"fmt"
	"strings"
)

func TestStr(t *testing.T) {
    str := ""
    for i := 0; i < 1000000; i++ {
        str += "test"
    }
    print(len(str))
}

func TestStrUseBuffer(t *testing.T)  {
	strBuf := bytes.NewBufferString("")
	for i := 0; i < 1000000; i++ {
		strBuf.WriteString("test")
	}
	print(strBuf.Len())
}

func TestStrConcat(t *testing.T) {
    a, b := "Hello", "world"
    for i := 0; i < 1000000; i++ {
        fmt.Sprintf("%s%s", a, b)
        //strings.Join([]string{a, b}, "")
    }
}

func TestStrConcatJoin(t *testing.T) {
    a, b := "Hello", "world"
    for i := 0; i < 1000000; i++ {
        //fmt.Sprintf("%s%s", a, b)
        strings.Join([]string{a, b}, "")
    }
}
