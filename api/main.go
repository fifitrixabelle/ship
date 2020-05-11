package main

import "github.com/gin-gonic/gin"
import "github.com/go-redis/redis/v7"
import "net/http"


func main() {
  a := gin.Default()
  r := redis.NewClient(&redis.Options{
    Addr:     "redis:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
  })

  a.GET("/status", func(c *gin.Context) {
    redisPong, redisErr := r.Ping().Result()
    c.JSON(http.StatusOK, gin.H{
      "redisPong": redisPong,
      "redisErr": redisErr,
    })
  })


 a.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
