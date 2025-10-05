package utils
import "strings"
func Split2(s string, sep byte) []string { i:=strings.IndexByte(s,sep); if i<0 { return []string{s} }; return []string{s[:i], s[i+1:]} }
