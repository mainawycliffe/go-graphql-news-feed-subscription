//go:generate go run github.com/99designs/gqlgen
package graphql

import (
	"context"
	"sync"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func GraphQLServer() *Resolver {

	return &Resolver{
		posts: nil,
		mutex: sync.Mutex{},
	}
}

type Resolver struct {
	posts []*Post
	mutex sync.Mutex
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
	panic("not implemented!")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetPosts(ctx context.Context) ([]*Post, error) {
	return r.posts, nil
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) NewPostAdded(ctx context.Context) (<-chan *Post, error) {
	panic("not implemented")
}
