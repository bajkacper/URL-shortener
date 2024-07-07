package utils

import (
		"math/rand"
			"time"
		)

		const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		const keyLength = 6

		func GenerateKey() string {
				rand.Seed(time.Now().UnixNano())
					key := make([]byte, keyLength)
						for i := range key {
									key[i] = charset[rand.Intn(len(charset))]
										}
											return string(key)
										}

