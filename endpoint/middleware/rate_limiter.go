package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"
	"yosimaril/CourseEnrollment/i18n"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Visitor struct {
	Limiter  *rate.Limiter
	LastSeen time.Time
}

var (
	visitors = make(map[string]*Visitor)
	mutex    sync.Mutex
)

const (
	RequestsPerSecond = 1
	BurstSize         = 10
)

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = "en"
		}

		ip := c.ClientIP()

		mutex.Lock()

		visitor := getVisitor(ip)

		visitor.LastSeen = time.Now()

		mutex.Unlock()

		allowed := visitor.Limiter.Allow()

		fmt.Printf(
			"[RateLimiter] IP=%-15s | Allowed=%-5t | Tokens=%5.2f | Time=%s\n",
			ip,
			allowed,
			visitor.Limiter.Tokens(),
			time.Now().Format("15:04:05"),
		)

		if !allowed {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": i18n.T(lang, "too_many_requests"),
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

func getVisitor(ip string) *Visitor {
	visitor, exists := visitors[ip]

	if !exists {
		visitor = &Visitor{
			Limiter: rate.NewLimiter(
				rate.Limit(RequestsPerSecond),
				BurstSize,
			),
		}

		visitors[ip] = visitor
	}
	return visitor
}

func CleanupVisitors() {
	ticker := time.NewTicker(time.Minute)

	for range ticker.C {
		mutex.Lock()

		for ip, visitor := range visitors {
			if time.Since(visitor.LastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}

		mutex.Unlock()
	}
}
