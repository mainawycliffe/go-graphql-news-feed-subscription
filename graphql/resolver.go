//go:generate go run github.com/99designs/gqlgen
package graphql

import (
	"context"

	badger "github.com/dgraph-io/badger"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func GetResolver(db *badger.DB) *Resolver {
	return &Resolver{
		db: db,
	}
}

type Resolver struct {
	db *badger.DB
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Share(ctx context.Context, post NewPost) (*Post, error) {

	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) PostsByCategory(ctx context.Context, postID string) ([]*Post, error) {
	panic("not implemented")
}

func (r *queryResolver) AllPosts(ctx context.Context) ([]*Post, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) NewPostsByCategory(ctx context.Context, category string) (<-chan []*Post, error) {
	panic("not implemented")
}

func (r *subscriptionResolver) NewPosts(ctx context.Context) (<-chan []*Post, error) {
	panic("not implemented")
}
