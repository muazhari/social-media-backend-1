package graphqls

import (
	"context"
	"github.com/google/uuid"
	"social-media-backend-1/internal/outers/container"
	"social-media-backend-1/internal/outers/deliveries/graphqls/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	RootContainer *container.RootContainer
	Dataloader    *Dataloader
}

func NewResolver(rootContainer *container.RootContainer) *Resolver {
	resolver := &Resolver{
		RootContainer: rootContainer,
	}
	resolver.Dataloader = NewDataloader(resolver)
	return resolver
}

func (r *Resolver) GetAccountsByIds(ctx context.Context, ids []string) ([]*model.Account, []error) {
	var convertedIds []*uuid.UUID
	for _, id := range ids {
		convertedId, err := uuid.Parse(id)
		if err != nil {
			return nil, []error{err}
		}
		convertedIds = append(convertedIds, &convertedId)
	}

	accounts, err := r.RootContainer.RepositoryContainer.AccountRepository.GetAccountsByIds(convertedIds)
	if err != nil {
		return nil, err
	}

	var results []*model.Account
	for _, account := range accounts {
		result := &model.Account{
			ID:               account.ID.String(),
			Name:             account.Name,
			Email:            account.Email,
			Password:         account.Password,
			TotalPostLike:    account.TotalPostLike,
			TotalChatMessage: account.TotalChatMessage,
		}
		results = append(results, result)
	}

	return results, nil
}
