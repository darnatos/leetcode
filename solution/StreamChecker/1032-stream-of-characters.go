package StreamChecker

type Node struct {
	next   [26]*Node
	isWord bool
}

type StreamChecker struct {
	root *Node
	sb   []byte
}

func Constructor(words []string) StreamChecker {
	root := new(Node)

	for _, w := range words {
		cur := root
		for i := len(w) - 1; i >= 0; i-- {
			ch := w[i] - 'a'
			if cur.next[ch] == nil {
				cur.next[ch] = new(Node)
			}
			cur = cur.next[ch]
		}
		cur.isWord = true
	}

	return StreamChecker{
		root: root,
		sb:   make([]byte, 0),
	}
}

func (sc *StreamChecker) Query(letter byte) bool {
	sc.sb = append(sc.sb, letter)
	cur := sc.root
	for i := len(sc.sb) - 1; i >= 0 && cur != nil; i-- {
		cur = cur.next[sc.sb[i]-'a']
		if cur != nil && cur.isWord {
			return true
		}
	}
	return false
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
