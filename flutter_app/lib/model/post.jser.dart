// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'post.dart';

// **************************************************************************
// JaguarSerializerGenerator
// **************************************************************************

abstract class _$PostJsonSerializer implements Serializer<Post> {
  @override
  Map<String, dynamic> toMap(Post model) {
    if (model == null) return null;
    Map<String, dynamic> ret = <String, dynamic>{};
    setMapValue(ret, 'id', model.id);
    setMapValue(ret, 'imageURL', model.imageURL);
    setMapValue(ret, 'title', model.title);
    setMapValue(ret, 'content', model.content);
    setMapValue(ret, 'link', model.link);
    setMapValue(
        ret, 'postedOn', dateTimeUtcProcessor.serialize(model.postedOn));
    return ret;
  }

  @override
  Post fromMap(Map map) {
    if (map == null) return null;
    final obj = Post();
    obj.id = map['id'] as String;
    obj.imageURL = map['imageURL'] as String;
    obj.title = map['title'] as String;
    obj.content = map['content'] as String;
    obj.link = map['link'] as String;
    obj.postedOn = dateTimeUtcProcessor.deserialize(map['postedOn'] as String);
    return obj;
  }
}
