package utils
import("crypto/sha256";"encoding/hex";"fmt")
func HashPassword(p,s string)string{h:=sha256.Sum256([]byte(s+":"+p));return hex.EncodeToString(h[:])}
func CheckPasswordHash(p,s,h string)bool{return HashPassword(p,s)==h}
func RandomSalt()string{return fmt.Sprintf("%x",RandomBytes(16))}
