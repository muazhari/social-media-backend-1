package graphqls

import (
	"context"
	"social-media-backend-1/internal/outers/container"
	"social-media-backend-1/internal/outers/deliveries/graphqls/model"

	"github.com/google/uuid"
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

func (r *Resolver) GetAccountsByIDs(ctx context.Context, ids []string) ([]*model.Account, []error) {
	var convertedIDs []*uuid.UUID
	for _, id := range ids {
		convertedID, err := uuid.Parse(id)
		if err != nil {
			return nil, []error{err}
		}
		convertedIDs = append(convertedIDs, &convertedID)
	}

	foundAccounts, err := r.RootContainer.UseCaseContainer.AccountUseCase.GetAccountsByIDs(ctx, convertedIDs)
	if err != nil {
		return nil, []error{err}
	}

	var results []*model.Account
	for _, foundAccount := range foundAccounts {
		result := &model.Account{
			ID:               foundAccount.ID.String(),
			ImageURL:         foundAccount.ImageURL,
			Name:             *foundAccount.Name,
			Email:            *foundAccount.Email,
			Password:         *foundAccount.Password,
			Scopes:           foundAccount.Scopes,
			TotalPostLike:    *foundAccount.TotalPostLike,
			TotalChatMessage: *foundAccount.TotalChatMessage,
		}
		results = append(results, result)
	}

	return results, nil
}
