import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_app/model/post.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

String get server => Platform.isAndroid ? '10.0.2.2' : 'localhost';

final String newPostAddedSubsciption = r'''
  subscription NewPostAdded{
    NewPostAdded {
      id
      title
      content
      link
      imageURL
      postedOn
    }
  }
  ''';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: GraphQLProvider(
        client: _client(),
        child: MyHomePage(
          title: 'GraphQL Subscription Demo',
        ),
      ),
    );
  }

  ValueNotifier<GraphQLClient> _client() {
    final httpLink = HttpLink(
      uri: "http://$server:8080/query",
    );

    final websocketLink = WebSocketLink(
      url: "ws://$server:8080/query",
      config: SocketClientConfig(
        autoReconnect: true,
      ),
    );

    final link = httpLink.concat(websocketLink);

    return ValueNotifier<GraphQLClient>(
      GraphQLClient(
        cache: OptimisticCache(
          dataIdFromObject: typenameDataIdFromObject,
        ),
        link: link,
      ),
    );
  }
}

class MyHomePage extends StatefulWidget {
  final String title;

  MyHomePage({Key key, this.title}) : super(key: key);

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  var _posts = new List<Post>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Subscription<Map<String, dynamic>>(
        'NewPostAdded',
        newPostAddedSubsciption,
        builder: ({dynamic loading, dynamic payload, dynamic error}) {
          if (error != null) {
            return Center(
              child: Text(error.toString()),
            );
          }

          if (loading == true) {
            return Center(
              child: const CircularProgressIndicator(),
            );
          }
          final serialize = new PostJsonSerializer();
          final post = serialize.fromMap(payload["NewPostAdded"]);

          _posts.add(post);

          return PostsList(_posts);
        },
      ),
    );
  }
}

class PostsList extends StatelessWidget {
  final List<Post> _posts;

  const PostsList(this._posts);

  @override
  Widget build(BuildContext context) {
    return ListView(children: [
      for (var post in _posts)
        ListTile(
          leading: _showImageUrl(post),
          title: Text("${post.title}"),
          subtitle: Text("${post.content}"),
        ),
    ]);
  }

  Widget _showImageUrl(Post post) {
    return post.imageURL != ""
        ? Image.network("http://$server:8080${post.imageURL}")
        : null;
  }
}
