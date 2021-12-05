package mapper

import (
	"github.com/nikitasadok/controllermapper/trie"
)

const domainName = "https://domain.com/"

type Mapper struct {
	TrieRoot *trie.Node
}

func NewMapper(controllers []string) *Mapper {
	m := &Mapper{TrieRoot: &trie.Node{Children: make(map[rune]*trie.Node)}}

	for _, str := range controllers {
		m.TrieRoot.Insert(str)
	}
	return m
}

func (m *Mapper) IsValidController(key string) bool {
	return m.TrieRoot.Find(m.TrieRoot, key[len(domainName):])
}
