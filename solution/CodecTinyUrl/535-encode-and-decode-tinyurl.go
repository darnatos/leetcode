package CodecTinyUrl

type Codec struct {
	root  string
	m     string
	cache map[int]string
}

func Constructor() Codec {
	return Codec{
		root:  "http://tinyurl.com/",
		m:     "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		cache: make(map[int]string),
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	var sb = make([]byte, 0, 16)
	var tiny string
	for {
		sb = sb[:0]
		id := len(this.cache)
		for id > 0 {
			sb = append(sb, this.m[id%62])
			id /= 62
		}
		for i, j := 0, len(sb)-1; i < j; i, j = i+1, j-1 {
			sb[i], sb[j] = sb[j], sb[i]
		}
		tiny = string(sb)
		if _, ok := this.cache[id]; !ok {
			this.cache[id] = longUrl
			break
		}
	}

	return this.root + tiny
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	id := 0
	for i := len(shortUrl) - 1; i >= 0 && shortUrl[i] != '/'; i-- {
		id *= 62
		if shortUrl[i] <= 'z' {
			id += int(shortUrl[i]) - 'a'
		} else if shortUrl[i] <= 'Z' {
			id += int(shortUrl[i]) - 'A' + 26
		} else {
			id += int(shortUrl[i]) - '0' + 52
		}
	}
	return this.cache[id]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
