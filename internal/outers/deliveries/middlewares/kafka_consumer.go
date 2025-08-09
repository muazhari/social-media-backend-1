package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"social-media-backend-1/internal/outers/configs"
	"social-media-backend-1/internal/outers/repositories"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	FiveDataStoreConfig *configs.FiveDatastoreConfig
	AccountRepository   *repositories.AccountRepository
}

type postLikeAddedEvent struct {
	AccountID string `json:"account_id"`
}

type chatMessageAddedEvent struct {
	AccountID string `json:"account_id"`
}

func NewKafkaConsumer(fiveDataStoreConfig *configs.FiveDatastoreConfig, accountRepository *repositories.AccountRepository) *KafkaConsumer {
	return &KafkaConsumer{
		FiveDataStoreConfig: fiveDataStoreConfig,
		AccountRepository:   accountRepository,
	}
}

func (c *KafkaConsumer) Start(ctx context.Context) error {
	go c.consumeTopic(ctx, "postLike.increment", c.handlePostLikeIncrement)
	go c.consumeTopic(ctx, "postLike.decrement", c.handlePostLikeDecrement)
	go c.consumeTopic(ctx, "chatMessage.increment", c.handleChatMessageIncrement)
	return nil
}

func (c *KafkaConsumer) consumeTopic(ctx context.Context, topic string, handler func(context.Context, []byte) error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{c.FiveDataStoreConfig.Broker},
		GroupID:  c.FiveDataStoreConfig.GroupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer r.Close()
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			log.Printf("kafka read error on %s: %v", topic, err)
			time.Sleep(time.Second)
			continue
		}
		if err := handler(ctx, m.Value); err != nil {
			log.Printf("kafka handler error on %s: %v", topic, err)
		}
	}
}

func (c *KafkaConsumer) handlePostLikeIncrement(ctx context.Context, payload []byte) error {
	var evt postLikeAddedEvent
	if err := json.Unmarshal(payload, &evt); err != nil {
		return fmt.Errorf("unmarshal postLike event: %w", err)
	}
	return c.AccountRepository.IncrementTotalPostLike(ctx, uuid.MustParse(evt.AccountID), 1)
}

func (c *KafkaConsumer) handlePostLikeDecrement(ctx context.Context, payload []byte) error {
	var evt postLikeAddedEvent
	if err := json.Unmarshal(payload, &evt); err != nil {
		return fmt.Errorf("unmarshal postLike event: %w", err)
	}
	return c.AccountRepository.DecrementTotalPostLike(ctx, uuid.MustParse(evt.AccountID), 1)
}

func (c *KafkaConsumer) handleChatMessageIncrement(ctx context.Context, payload []byte) error {
	var evt chatMessageAddedEvent
	if err := json.Unmarshal(payload, &evt); err != nil {
		return fmt.Errorf("unmarshal chatMessage event: %w", err)
	}
	return c.AccountRepository.IncrementTotalChatMessage(ctx, uuid.MustParse(evt.AccountID), 1)
}
