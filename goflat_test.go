package goflat

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input  []interface{}
	output []int
}{
	// Empty
	{
		input:  []interface{}{},
		output: []int{},
	},
	// Examples
	{
		input:  []interface{}{0},
		output: []int{0},
	},
	{
		input:  []interface{}{0, []interface{}{1}},
		output: []int{0, 1},
	},
	{
		input:  []interface{}{[]interface{}{[]interface{}{5}}},
		output: []int{5},
	},
	{
		input:  []interface{}{0, 1, []interface{}{5, []interface{}{8, 9}}},
		output: []int{0, 1, 5, 8, 9},
	},
	// Bad input examples
	// Shows that it works only with integers
	{
		input:  []interface{}{[]interface{}{[]interface{}{5.6}}},
		output: []int{},
	},
}

func TestFlatten(t *testing.T) {

	for index, test := range tests {
		result := Flatten(test.input)

		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("Test %d output was incorrect. Expected %v, got %v.", index, test.output, result)
		}

	}

}

func benchmarkFlatten(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Flatten(tests[i].input)
	}
}

func BenchmarkFlatten0(b *testing.B) {	benchmarkFlatten(0, b) }
func BenchmarkFlatten1(b *testing.B) {	benchmarkFlatten(1, b) }
func BenchmarkFlatten2(b *testing.B) {	benchmarkFlatten(2, b) }
func BenchmarkFlatten3(b *testing.B) {	benchmarkFlatten(3, b) }
func BenchmarkFlatten4(b *testing.B) {	benchmarkFlatten(4, b) }
func BenchmarkFlatten5(b *testing.B) {	benchmarkFlatten(5, b) }
