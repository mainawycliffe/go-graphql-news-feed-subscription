# GraphQL Subscription Example

This is a demo to demonstrate [GraphQL Subscriptions](https://graphql.org/blog/subscriptions-in-graphql-and-relay/). The backend is build using Golang while the frontend is build using Flutter.

To demo GraphQL Subscriptions, I went with a very simple news feed example. As new posts are being added (shared), you can subscribe and get them in real time. With such a simple example, it is easy to add new post using GraphiQL and watch them show up in the app in real time.

## Backend

To run the backend, you can use [docker compose](https://docs.docker.com/compose/), simply run `docker-compose up -d --build` at the root directory.

### Adding New Post

To see subscription in action, you will need to add some posts while the flutter app is open. You can do this using the built in GraphiQL, whose URL is [http://localhost:8080](http://localhost:8080). 

Use the share mutation to add new posts:

```graphql
mutation share($title: String!, $summary: String!, $link: String!){
  share(post: { title: $title, summary: $summary,  link: $link}) {
    id
    title
    content
    link
    postedOn
  }
}
```

### Query Existing Posts

On top of that, you can query for existing posts:

```graphql
query getPosts {
  getPosts {
    id
    title
    content
    link
    postedOn
    imageURL
  }
}
```

## Frontend

To run the frontend, change directory to the `flutter_app` directory, and run `flutter run`. Make sure you have flutter [installed](https://flutter.dev/docs/get-started/install).

A follow up post is coming soon.
