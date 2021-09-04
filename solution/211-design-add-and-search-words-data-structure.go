package solution

type WordDictionary struct {
	children       [26]*WordDictionary
	isTerminalWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		children:       [26]*WordDictionary{},
		isTerminalWord: false,
	}
}

func (t *WordDictionary) AddWord(word string) {
	curr := t
	for i := range word {
		ch := word[i] - 'a'
		if curr.children[ch] == nil {
			trie := Constructor()
			curr.children[ch] = &trie
		}
		curr = curr.children[ch]
	}

	curr.isTerminalWord = true
}

func (t *WordDictionary) Search(word string) bool {
	curr := t
	for i := range word {
		if curr == nil {
			return false
		}
		ch := word[i] - 'a'
		if ch+'a' == '.' {
			for _, v := range curr.children {
				if v.Search(word[i+1:]) {
					return true
				}
			}
			return false
		}
		if curr.children[ch] == nil {
			return false
		}
		curr = curr.children[ch]
	}
	if curr == nil {
		return false
	}
	return curr.isTerminalWord
}
