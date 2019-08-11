import 'package:jaguar_serializer/jaguar_serializer.dart';

part 'post.jser.dart';

class Post {
  String id;
  String imageURL;
  String title;
  String content;
  String link;
  DateTime postedOn;
}

@GenSerializer()
class PostJsonSerializer extends Serializer<Post> with _$PostJsonSerializer {}
