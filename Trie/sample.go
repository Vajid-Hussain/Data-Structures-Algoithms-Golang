// package main

// import "fmt"

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

// func (t *trie) exist(word string) bool {
// 	current := t.root

// 	for _, val := range word {
// 		index := val - 'a'
// 		if current.children[index] == nil {
// 			return false
// 		}
// 		current = current.children[index]
// 	}
// 	if current.word {
// 		return true
// 	}
// 	return false
// }

// func (t *trie)delete(node *trieNode, word string, index int) *trieNode{
// 	if node==nil{
// 		return nil
// 	}
// 	if index==len(word){
// 		node.word=false
// 		if nextExist(node){
// 			return nil
// 		}
// 		return node
// 	}

// 	currentIntex:=word[index]-'a'
// 	node.children[currentIntex]=t.delete(node.children[currentIntex], word, index+1)
// 	if nextExist(node){
// 		return nil
// 	}
// 	return node
// }

// func nextExist(node *trieNode)bool{
// 	for _,val:=range node.children{
// 		if val!=nil{
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	hp := trie{root: &trieNode{}}
// 	hp.insert("vajid")
// 	hp.insert("vajidhussain")
// 	hp.delete(hp.root,"vajidhussain",0)
// 	hp.delete(hp.root, "vajid",0)
// 	fmt.Println(hp.exist("vajidhussain"))
// 	fmt.Println(hp.exist("vajid"))
// }
