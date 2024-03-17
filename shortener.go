
package main

import (
    "crypto/sha1"
    "encoding/base64"
    "github.com/go-redis/redis/v8"
)

const keyPrefix = "urlshortener:"


func redirectToURL(rdb *redis.Client, shortURL string) (string, error) {
    url, err := rdb.Get(ctx, keyPrefix+shortURL).Result()
    if err != nil {
        return "", err
    }

    return url, nil
}


func shortenURL(rdb *redis.Client, url string) (string, error) {
    hash := sha1.New()
    hash.Write([]byte(url))
    shortURL := base64.URLEncoding.EncodeToString(hash.Sum(nil))[:8]

    err := rdb.Set(ctx, keyPrefix+shortURL, url, 0).Err()
    if err != nil {
        return "", err
    }

    return shortURL, nil
}

