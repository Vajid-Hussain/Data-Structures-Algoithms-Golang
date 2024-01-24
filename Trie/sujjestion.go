// package main

// import (
// 	"fmt"
// )

// type trieNode struct {
// 	word     bool
// 	children [26]*trieNode
// }

// type trie struct {
// 	root *trieNode
// }

// func (t *trie) insert(word string) {
// 	current := t.root
// 	for _, val := range word {
// 		index := val - 'a'
// 		if current.children[index] == nil {
// 			current.children[index] = &trieNode{}
// 		}
// 		current = current.children[index]
// 	}
// 	current.word = true
// }

// func (t *trie) serch(word string) bool {
// 	current := t.root

// 	for _, val := range word {
// 		index := val - 'a'
// 		if current.children[index] == nil {
// 			return false
// 		}
// 		current = current.children[index]
// 	}

// 	return current.word
// }

// func (t *trie) delete(node *trieNode, word string, index int) *trieNode {
// 	if node == nil {
// 		return nil
// 	}
// 	if len(word) == index {
// 		node.word = false
// 		if exist(node) {
// 			return nil
// 		}
// 		return node
// 	}

// 	currentIndex := word[index]-'a'
// 	node.children[currentIndex] = t.delete(node.children[currentIndex], word, index+1)
// 	if exist(node) {
// 		return nil
// 	}
// 	return node
// }

// func exist(node *trieNode) bool {
// 	for _, val := range node.children {
// 		if val != nil {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	trie := trie{root: &trieNode{}}
// 	trie.insert("vajid")
// 	trie.insert("vajidhussain")
// 	fmt.Println(trie.serch("vajidhussain"))
// 	trie.delete(trie.root, "vajid", 0)
// 	fmt.Println(trie.serch("vajid"))
// }
