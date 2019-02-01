package utils

import (
	// "fmt"
	// "time"

	"samh_common_lib/base"
	"utils/config"
	"utils/log"

	"github.com/go-redis/redis"
)

func InitRedisClient(rc config.Redis) (client *redis.Client) {
	log.Debug(base.NowFunc())

	client = redis.NewClient(&redis.Options{
		Network:    rc.Network,
		Addr:       rc.Addr,
		Password:   rc.Password,
		MaxRetries: rc.Max_retries,
		PoolSize:   rc.Pool_size,
	})
	err := client.Ping().Err()
	if err != nil {
		log.Panic(err)
	}

	return
}

func InitRedisCluster(rc config.RedisCluster) (clusterClient *redis.ClusterClient) {
	log.Debug(base.NowFunc())

	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Password:   rc.Password,
		MaxRetries: rc.Max_retries,
		PoolSize:   rc.Pool_size,
		ClusterSlots: func() (slots []redis.ClusterSlot, err error) {
			len := len(rc.Master_addr_arr)
			gap := 16384 / len
			for i := 0; i < len; i++ {
				slots = append(slots, redis.ClusterSlot{
					Start: i * gap,
					End:   (i+1)*gap - 1,
					Nodes: []redis.ClusterNode{{
						Addr: rc.Master_addr_arr[i],
					}, {
						Addr: rc.Slave_addr_arr[i],
					}},
				})
			}
			return
		},
	})
	var err error
	err = clusterClient.Ping().Err()
	if err != nil {
		log.Panic(err)
	}
	err = clusterClient.ReloadState()
	if err != nil {
		log.Panic(err)
	}

	return
}
