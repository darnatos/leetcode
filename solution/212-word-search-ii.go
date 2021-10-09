package solution

type Trie struct {
	next [26]*Trie
	word string
}

func FindWords(board [][]byte, words []string) []string {
	ret := make([]string, 0, len(words))
	trie := BuildTrie(words)
	for i := range board {
		for j := range board[0] {
			fwDfs(board, i, j, trie, &ret)
		}
	}
	return ret
}

func BuildTrie(words []string) *Trie {
	r := &Trie{}
	for _, word := range words {
		p := r
		for _, ch := range word {
			idx := ch - 'a'
			if p.next[idx] == nil {
				p.next[idx] = &Trie{}
			}
			p = p.next[idx]
		}
		p.word = word
	}
	return r
}

func fwDfs(board [][]byte, i, j int, p *Trie, ret *[]string) {
	ch := board[i][j]
	if ch == '#' || p.next[ch-'a'] == nil {
		return
	}
	p = p.next[ch-'a']
	if len(p.word) > 0 {
		*ret = append(*ret, p.word)
		p.word = ""
	}

	board[i][j] = '#'
	if i > 0 {
		fwDfs(board, i-1, j, p, ret)
	}
	if i < len(board)-1 {
		fwDfs(board, i+1, j, p, ret)
	}
	if j > 0 {
		fwDfs(board, i, j-1, p, ret)
	}
	if j < len(board[0])-1 {
		fwDfs(board, i, j+1, p, ret)
	}

	board[i][j] = ch
}
