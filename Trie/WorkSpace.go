package main

import "fmt"

type trieNode struct {
	word     bool
	children [26]*trieNode
}

func (t *trieNode) insertion(word string) {
	current := t
	for _, val := range word {
		index := val - 'a'
		if current.children[index] == nil {
			current.children[index] = &trieNode{}
		}
		current = current.children[index]
	}
	current.word = true
}

func (t *trieNode) serch(word string) bool {
	current := t
	for _, val := range word {
		index := val - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	if current.word {
		return true
	}
	return false
}

func (t *trieNode) delete(node *trieNode, word string, intex int) *trieNode {
	if node == nil {
		return nil
	}
	if len(word) == intex {
		if exist(node) {
			node.word = false
			return nil
		}
		return node
	}

	currentindtex := word[intex] - 'a'
	node.children[currentindtex] = t.delete(node.children[currentindtex], word, intex+1)
	if exist(node) {
		return nil
	}
	return node
}

func exist(node *trieNode) bool {
	for _, val := range node.children {
		if val != nil {
			return false
		}
	}
	return true
}

func (r *trieNode) collectWord(node *trieNode, prefix string, suggestion *[]string) *[]string {
	if node == nil {
		return nil
	}

	if node.word {
		*suggestion = append(*suggestion, prefix)
	}

	for key, val := range node.children {
		if val != nil {
			nextWord := string(key + 'a')
			r.collectWord(val, prefix+nextWord, suggestion)
		}
	}
	return suggestion
}

func (t *trieNode) autocomplete(prefix string) *[]string {
	node, exiist := t.prefixEnd(prefix)
	if !exiist {
		fmt.Println("*")
		return nil
	}
	var suggestion []string
	words := t.collectWord(node, prefix, &suggestion)
	return words
}

func (t *trieNode) prefixEnd(prefix string) (*trieNode, bool) {
	current := t
	for _, val := range prefix {
		intex := val - 'a'
		if current.children[intex] == nil {
			return nil, false
		}
		current = current.children[intex]
	}
	return current, true
}

func main() {
	trie := trieNode{children: [26]*trieNode{}}
	trie.insertion("vajid")
	trie.insertion("vajidhuss")
	trie.insertion("vajidh")
	trie.insertion("vajidhus")
	trie.insertion("vajidhussain")
	trie.delete(&trie ,"vajidhussain", 0)
	// trie.delete(&trie,"vajid",0)
	fmt.Println(trie.autocomplete("vaji"))
	// fmt.Println("search:", trie.serch("vajid"))
}
