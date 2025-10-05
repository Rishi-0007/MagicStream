package middleware

import (
	"net"
	"sync"
	"time"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type visitor struct { limiter *rate.Limiter; lastSeen time.Time }
var ( visitors = map[string]*visitor{}; mu sync.Mutex )

func getVisitor(ip string, r rate.Limit, b int) *rate.Limiter {
	mu.Lock(); defer mu.Unlock()
	v, ok := visitors[ip]; if !ok { lim := rate.NewLimiter(r, b); visitors[ip] = &visitor{lim, time.Now()}; return lim }
	v.lastSeen = time.Now(); return v.limiter
}
func cleanupVisitors(){ for { time.Sleep(time.Minute); mu.Lock(); for ip,v := range visitors { if time.Since(v.lastSeen) > 3*time.Minute { delete(visitors, ip) } }; mu.Unlock() } }

func RateLimit(r rate.Limit, b int) gin.HandlerFunc {
	go cleanupVisitors()
	return func(c *gin.Context){
		ip,_,err := net.SplitHostPort(c.Request.RemoteAddr); if err != nil { ip = c.ClientIP() }
		if !getVisitor(ip,r,b).Allow(){ c.AbortWithStatus(429); return }
		c.Next()
	}
}
