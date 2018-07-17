package go_benchmark

import (
	"fmt"
	"testing"
	"strings"
	"bytes"
)

func BenchmarkStrJoinSprintf(b *testing.B) {
	v := "ni shuo wo shi bu shi tai wu liao le a?"
	var s string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("%s[%s]", s, v)
	}
	s = s + ""
}

func BenchmarkStrJoinStringsJoin(b *testing.B) {
	v := "ni shuo wo shi bu shi tai wu liao le a?"
	var s string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = strings.Join([]string{s, "[", v, "]"}, "")
	}
	s = s + ""
}

func BenchmarkStrJoinAdd(b *testing.B) {
	v := "ni shuo wo shi bu shi tai wu liao le a?"
	var s string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = s + "[" + v + "]"
	}
	s = s + ""
}

func BenchmarkStrBuf(b *testing.B) {
	v := "ni shuo wo shi bu shi tai wu liao le a?"
	var s string
	var buf bytes.Buffer
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.WriteString("[")
		buf.WriteString(v)
		buf.WriteString("]")
		s = buf.String()
	}
	s = s + ""
}

func BenchmarkStrChange(b *testing.B) {
	var s = "456"
	//var h = ""

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = "123"
		s = s + "234"
		//h = s + h + "123"
	}
}