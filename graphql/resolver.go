//go:generate go run github.com/99designs/gqlgen
package graphql

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

const imageDir = "images"

func GraphQLServer() *Resolver {

	return &Resolver{
		Posts: nil,
		mutex: sync.Mutex{},
	}
}

type Resolver struct {
	Posts         []*Post
	ListenToPosts map[string]struct {
		Post chan *Post
	}
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

	url, err := saveUploadedFile(post.Image)

	if err != nil {
		log.Printf("Error saving file: %v", err)
		return nil, fmt.Errorf("Error saving file")
	}

	newPost := &Post{
		ID:       randString(20),
		Title:    post.Title,
		Content:  post.Summary,
		Link:     post.Link,
		PostedOn: time.Now(),
		ImageURL: url,
	}

	r.mutex.Lock()
	r.Posts = append(r.Posts, newPost)

	// notify listeners
	for _, v := range r.ListenToPosts {
		v.Post <- newPost
	}

	r.mutex.Unlock()

	return newPost, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetPosts(ctx context.Context) ([]*Post, error) {
	r.mutex.Lock()
	posts := r.Posts
	r.mutex.Unlock()

	return posts, nil
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) NewPostAdded(ctx context.Context) (<-chan *Post, error) {

	id := randString(10)
	events := make(chan *Post, 1)

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.ListenToPosts, id)
		r.mutex.Unlock()
	}()

	r.mutex.Lock()
	r.ListenToPosts[id] = struct {
		Post chan *Post
	}{
		Post: events,
	}

	r.mutex.Unlock()

	return events, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func saveUploadedFile(image *graphql.Upload) (string, error) {

	filaname := fmt.Sprintf("%s%s", randString(20), filepath.Ext(image.Filename))

	f, err := os.Create(fmt.Sprintf("%s/%s", imageDir, filaname))

	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadAll(image.File)

	if err != nil {
		return "", err
	}

	_, err = f.Write(content)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://localhost:8080/%s/%s", imageDir, image.Filename), nil
}
