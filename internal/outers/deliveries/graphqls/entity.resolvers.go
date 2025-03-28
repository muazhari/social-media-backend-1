package graphqls

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.70

import (
	"context"
	"social-media-backend-1/internal/outers/deliveries/graphqls/model"
)

// FindPostByID is the resolver for the findPostByID field.
func (r *entityResolver) FindPostByID(ctx context.Context, id string) (*model.Post, error) {
	result := &model.Post{
		ID: id,
	}

	return result, nil
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
