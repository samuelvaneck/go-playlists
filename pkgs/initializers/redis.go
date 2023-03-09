package initializers

import (
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisPool is a redis connection pool
func ConnectToRedis() (redis.Conn, error) {
	redisServerAddr := os.Getenv("REDIS_URL")
	const healthCheckPeriod = time.Minute
	var err error

	c, err := redis.Dial("tcp", redisServerAddr,
		// Read timeout on server should be greater than ping period.
		redis.DialReadTimeout(healthCheckPeriod+10*time.Second),
		redis.DialWriteTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	psc := redis.PubSubConn{Conn: c}

	if err = psc.Ping(""); err != nil {
		return nil, err
	}

	return c, nil
}
