package main

import (
	"fmt"

	"github.com/nikitasadok/controllermapper/mapper"
	"github.com/nikitasadok/controllermapper/trie"
)

func main() {
	tr := trie.Node{Children: make(map[rune]*trie.Node)}
	node := &tr
	tr.Insert("nikita")
	fmt.Println(tr.Find(node, "anikita"))

	m := mapper.NewMapper([]string{"jopa", "nikita", "sadok", "authorize"})
	fmt.Println(m.IsValidController("https://domain.com/sadok?hello=1"))
	fmt.Println(m.IsValidController("https://domain.com/jopajopa?hello=1"))
	fmt.Println(m.IsValidController("https://domain.com/autthorize?hello=1"))

}
