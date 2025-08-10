package gateways

import (
	"context"
	"os"
	"social-media-backend-1/internal/outers/repositories"

	"github.com/google/uuid"
	"github.com/hasura/go-graphql-client"
	"github.com/hasura/go-graphql-client/pkg/jsonutil"
)

type EDFGateway struct {
	AccountRepository *repositories.AccountRepository
	Client            *graphql.SubscriptionClient
}

func NewEDFGateway(accountRepository *repositories.AccountRepository) *EDFGateway {
	routerURL := "ws://" + os.Getenv("ROUTER_1_HOST") + os.Getenv("ROUTER_1_PORT") + "/graphql"
	client := graphql.NewSubscriptionClient(routerURL).WithProtocol(graphql.GraphQLWS)

	return &EDFGateway{
		AccountRepository: accountRepository,
		Client:            client,
	}
}

func (s *EDFGateway) Start(ctx context.Context) error {
	var postLikeIncremented struct {
		PostLikeIncremented struct {
			ID *uuid.UUID `graphql:"id"`
		} `graphql:"postLikeIncremented"`
	}
	_, err := s.Client.Subscribe(&postLikeIncremented, nil, func(message []byte, err error) error {
		data := postLikeIncremented
		err = jsonutil.UnmarshalGraphQL(message, &data)
		if err != nil {
			return err
		}
		err = s.AccountRepository.IncrementTotalPostLike(ctx, data.PostLikeIncremented.ID, 1)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	var postLikeDecremented struct {
		PostLikeDecremented struct {
			ID *uuid.UUID `graphql:"id"`
		} `graphql:"postLikeDecremented"`
	}
	_, err = s.Client.Subscribe(&postLikeDecremented, nil, func(message []byte, err error) error {
		data := postLikeDecremented
		err = jsonutil.UnmarshalGraphQL(message, &data)
		if err != nil {
			return err
		}
		err = s.AccountRepository.DecrementTotalPostLike(ctx, data.PostLikeDecremented.ID, 1)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	var chatMessageIncremented struct {
		ChatMessageIncremented struct {
			ID *uuid.UUID `graphql:"id"`
		} `graphql:"chatMessageIncremented"`
	}
	_, err = s.Client.Subscribe(&chatMessageIncremented, nil, func(message []byte, err error) error {
		data := chatMessageIncremented
		err = jsonutil.UnmarshalGraphQL(message, &data)
		if err != nil {
			return err
		}
		err = s.AccountRepository.IncrementTotalPostLike(ctx, data.ChatMessageIncremented.ID, 1)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = s.Client.Run()
	if err != nil {
		return err
	}

	return nil
}
