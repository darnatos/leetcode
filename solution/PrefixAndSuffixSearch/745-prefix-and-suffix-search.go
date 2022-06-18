package PrefixAndSuffixSearch

type TrieNode struct {
	children [27]*TrieNode
	weight   int
}

type WordFilter struct {
	root *TrieNode
}

func Constructor(words []string) WordFilter {
	root := new(TrieNode)
	for weight := range words {
		word := words[weight] + "{"
		for i := range word {
			cur := root
			cur.weight = weight
			for j := i; j < 2*len(word)-1; j++ {
				k := word[j%len(word)] - 'a'
				if cur.children[k] == nil {
					cur.children[k] = new(TrieNode)
				}
				cur = cur.children[k]
				cur.weight = weight
			}
		}
	}
	return WordFilter{root}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	cur := this.root
	for _, c := range suffix + "{" + prefix {
		if cur.children[c-'a'] == nil {
			return -1
		}
		cur = cur.children[c-'a']
	}
	return cur.weight
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
