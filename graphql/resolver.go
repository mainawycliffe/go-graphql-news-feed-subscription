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

	"github.com/h2non/filetype"

	"github.com/99designs/gqlgen/graphql"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

const imageDir = "images"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GraphQLServer() *Resolver {

	return &Resolver{
		Posts:         nil,
		ListenToPosts: map[string]struct{ Post chan *Post }{},
		mutex:         sync.Mutex{},
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

	url := ""
	var err error

	// image is optional
	if post.Image != nil {
		url, err = saveUploadedFile(post.Image)

		if err != nil {
			log.Printf("Error saving file: %v", err)

			if err.Error() == "invalid image" {
				return nil, fmt.Errorf("The file you uploaded is not an image")
			}

			return nil, fmt.Errorf("Error saving file")
		}
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

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func saveUploadedFile(image *graphql.Upload) (string, error) {

	filename := fmt.Sprintf("%s%s", randString(20), filepath.Ext(image.Filename))

	f, err := os.Create(fmt.Sprintf("%s/%s", imageDir, filename))

	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadAll(image.File)

	if err != nil {
		return "", err
	}

	// check if file is supported
	if !filetype.IsImage(content) {
		return "", fmt.Errorf("invalid image")
	}

	_, err = f.Write(content)

	if err != nil {
		return "", err
	}

	// 	return fmt.Sprintf("http://localhost:8080/%s/%s", imageDir, filename), nil
	// return the image path, due to android and ios emulator having different localhost paths
	// under norm circumstance, this would be upload to an image server and return the CDN url
	return fmt.Sprintf("/%s/%s", imageDir, filename), nil
}
