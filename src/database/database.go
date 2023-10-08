package database

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"

	"github.com/miraliumre/clarity-api/settings"
	"github.com/nitishm/go-rejson/v4"
	"github.com/redis/go-redis/v9"
)

type Connection struct {
	client *redis.Client
}

type Document struct {
	Title     string `json:"title" binding:"required"`
	Author    string `json:"author" binding:"required"`
	Recipient string `json:"recipient" binding:"required"`
	Timestamp string `json:"timestamp" binding:"required"`
	Comments  string `json:"comments"`
}

func Init(s *settings.Settings) (*Connection, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     s.RedisAddr,
		Password: s.RedisPassword,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &Connection{client}, nil
}

func GetDocument(conn *Connection, hash string) (*Document, error) {
	var buffer bytes.Buffer

	h := rejson.NewReJSONHandler()
	h.SetGoRedisClientWithContext(context.Background(), conn.client)

	val, err := h.JSONGet(hash, ".")
	if err != nil {
		return nil, err
	}

	err = binary.Write(&buffer, binary.LittleEndian, val)
	if err != nil {
		return nil, err
	}

	doc := Document{}
	err = json.Unmarshal(buffer.Bytes(), &doc)

	return &doc, nil
}
