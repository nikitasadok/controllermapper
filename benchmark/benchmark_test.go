package benchmark_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/nikitasadok/controllermapper/bruteforcemapper"
	"github.com/nikitasadok/controllermapper/mapper"
)

const domainName = "https://domain.com/"

var isValid bool

func BenchmarkIsValidController(b *testing.B) {
	algos := []struct {
		name string
		f    func(b *testing.B, numControllers, longestController int)
	}{
		{name: "trie_mapper", f: benchMarkTrieMapper},
		{name: "bruteforce_mapper", f: benchMarkBruteforceMapper},
	}

	sizes := []struct {
		numControllers int
		longest        int
	}{
		{10, 100},
		{100, 100},
		{1000, 100},
		{1000, 10000},
		//{1000, 10000},
		//{1000, 1000000},
	}

	for _, a := range algos {
		for _, y := range sizes {
			n := a.name + " " + fmt.Sprintf("num_controllers: %d, longest: %d", y.numControllers, y.longest)
			b.Run(n, func(b *testing.B) {
				a.f(b, y.numControllers, y.longest)
			})
		}
	}
}

func benchMarkTrieMapper(b *testing.B, numControllers, longestController int) {
	b.StopTimer()
	controllers := generateRandomStrings(numControllers, longestController-len(domainName))
	m := mapper.NewMapper(controllers)
	inputs := generateRandomStrings(b.N, longestController-len(domainName))
	b.StartTimer()

	var valid bool
	for i := 0; i < b.N; i++ {
		valid = m.IsValidController(inputs[i])
	}

	isValid = valid
}

func benchMarkBruteforceMapper(b *testing.B, numControllers, longestController int) {
	b.StopTimer()
	controllers := generateRandomStrings(numControllers, longestController-len(domainName))
	m := bruteforcemapper.NewMapper(controllers)
	inputs := generateRandomStrings(b.N, longestController-len(domainName))
	b.StartTimer()

	var valid bool
	for i := 0; i < b.N; i++ {
		valid = m.IsValidController(inputs[i])
	}

	isValid = valid
}
func generateRandomStrings(n, size int) []string {
	var letterRunes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var res []string
	for i := 0; i < n; i++ {
		tmp := make([]byte, size)
		for j := 0; j < size; j++ {
			tmp[j] = letterRunes[rand.Intn(len(letterRunes))%size]
		}
		res = append(res, string(tmp))
	}

	return res
}
