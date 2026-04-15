package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter() gin.HandlerFunc {
	// this can grow infinitely and each ip will live forever.
	// to clean up old limiters, a background goroutine could be used.
	var limiters sync.Map

	go func(){
		for range time.Tick(10 * time.Minute){
			limiters.Range(func(k,v any) bool {
				limiter := v.(*rate.Limiter)
				if limiter.Tokens() >= float64(limiter.Burst()){
					limiters.Delete(k)
				}
				return true
			})
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()

		var limiter *rate.Limiter
		// done this way becase each new ratelimiter is memory allocated
		// and it's just unnecesary to declare each time no matter what.
		if val, ok := limiters.Load(ip); ok {
			limiter = val.(*rate.Limiter)
		} else {
			ratelimit := rate.NewLimiter(
				rate.Every(time.Second * 5), //5 seconds
				2, //burst = tokens (requests)
			)
			val, _ := limiters.LoadOrStore(ip, ratelimit)
			limiter = val.(*rate.Limiter)
		}


		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"Error": "Limite exceeded",
			})
			return
		}
		c.Next()
	}
}
