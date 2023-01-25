package utils

// GenerateRandomKey - generates a random key for the hill cipher
//
//	@param n - the size of the key
//	@return key - the key
func GenerateRandomKey(n int) string {
	// create a slice of bytes
	b := make([]byte, n)

	// loop through the key and add it to the key matrix
	for i, cache, remain := n-1, SeedSrc.Int63(), LetterIndexMask; i >= 0; {
		// if the cache is exhausted, get a new one
		if remain == 0 {
			cache, remain = SeedSrc.Int63(), LetterIndexMask
		} else {
			// get a random index from the key
			if idx := int(cache) & LetterIndexMask; idx < len(SourceKey) {
				// add the letter to the key
				b[i] = SourceKey[idx]
				i--
			}
			// shift the cache and decrement the remain count
			cache >>= LetterIndexBits
			remain--
		}

	}

	return string(b)
}
