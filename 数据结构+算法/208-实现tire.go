package main
type Trie struct {
	isWord   bool
	children map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{isWord: false, children: make(map[rune]*Trie)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	parent := this
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
		} else {
			newChild := &Trie{children: make(map[rune]*Trie)}
			parent.children[ch] = newChild
			parent = newChild
		}
	}
	parent.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	parent := this
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return parent.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, ch := range prefix {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return true
}
