import 'dart:convert';

import 'package:memegram/model.dart';
import 'package:http/http.dart' as http;

// ignore: constant_identifier_names
const BASE_URL = "http://127.0.0.1:9000";

class Http {
  static Future<List<Post>?> getPosts() async {
    final response = await http.get(Uri.parse('$BASE_URL/posts'));
    List<Post> result = [];
    List<dynamic> decoded = jsonDecode(response.body);

    for (var decodedPost in decoded) {
      Post post = Post();
      post.created_at = decodedPost["Created_at"]!;
      post.id = decodedPost["Id"]!;
      post.likes = decodedPost["Likes"]!;
      post.post_refer = decodedPost["Post_refer"]!;
      post.updated_at = decodedPost["Updated_at"]!;
      post.user_id = decodedPost["User_id"]!;

      result.add(post);
    }

    return result;
  }

  static Future<User?> getUserById(String id) async {
    final response = await http.get(Uri.parse('$BASE_URL/users/$id'));
    dynamic decodedUser = jsonDecode(response.body);

    User user = User();
    user.id = decodedUser["Id"]!;
    user.user = decodedUser["User"]!;
    user.created_at = decodedUser["Created_at"]!;
    user.updated_at = decodedUser["Updated_at"]!;

    return user;
  }
}
