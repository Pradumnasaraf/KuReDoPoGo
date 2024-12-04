package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	limiter     *redis_rate.Limiter
)

func RedisInit() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URI"),
	})
	limiter = redis_rate.NewLimiter(redisClient)
}

func RedisClose() {
	err := redisClient.Close()
	if err != nil {
		log.Fatal("Failed to close the redis client connection")
	}
}

func RedisRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := context.Background()
		limitInt, _ := strconv.Atoi(os.Getenv("REDIS_RATE_LIMIT"))

		res, err := limiter.Allow(ctx, c.ClientIP(), redis_rate.PerHour(limitInt))
		if err != nil {
			log.Fatal("unable to connect to redis instance")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
			c.Abort()
			return
		}
		if res.Remaining == 0 {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
