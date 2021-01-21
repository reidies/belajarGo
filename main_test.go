package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}

func TestMySumAssert(t *testing.T) {
	if runtime.GOOS != "darwin"{
		t.Skip("can run only in macos")
	}
	x := mySum(2, 3)
	assert.Equal(t, 5, x, "result is not equal")
}

func TestMySumAssertEnvironment (t *testing.T) {

	//subtest
	t.Run("5", func(t *testing.T){
		x := mySum(2, 3)
		require.Equal(t, 5, x, "result must be 5")
	})
	t.Run("9", func(t *testing.T){
		x := mySum(2, 3,4)
		require.Equal(t, 9, x, "result must be 9")
	})
}

func TestMySum1(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
		test{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func BenchmarkSayHello(b *testing.B) {
	var eko Person
	eko.Name = "rudi"
	for i := 0; i < b.N; i++ {
		SayHello(eko)
	}
}