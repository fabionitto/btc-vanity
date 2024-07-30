package btcvanity

import "strings"

// isMatch checks if the btc wallet matches the prefix of pattern
func isMatch(pattern string, addr string) bool {
	addr = addr[1 : len(addr)-1]
	
	return strings.HasPrefix(addr, pattern)
}
