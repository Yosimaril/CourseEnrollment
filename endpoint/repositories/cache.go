package repositories

import (
	"fmt"

	"yosimaril/CourseEnrollment/config"
)

func InvalidateCachePrefixes(prefixes ...string) {
	if config.Redis == nil {
		return
	}

	deleted := map[string]struct{}{}

	for _, prefix := range prefixes {
		if prefix == "" {
			continue
		}

		iter := config.Redis.Scan(config.Ctx, 0, prefix+"*", 0).Iterator()
		for iter.Next(config.Ctx) {
			key := iter.Val()
			if _, exists := deleted[key]; exists {
				continue
			}

			if err := config.Redis.Del(config.Ctx, key).Err(); err != nil {
				fmt.Println("[Redis] cache invalidation error:", err)
				break
			}

			deleted[key] = struct{}{}
		}

		if err := iter.Err(); err != nil {
			fmt.Println("[Redis] cache invalidation error:", err)
		}
	}
}