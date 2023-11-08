import 'package:flutter/material.dart';
import 'package:memegram/http.dart';
import 'package:memegram/model.dart';

class UsernameById extends StatefulWidget {
  final String id;
  const UsernameById(this.id, {Key? key}) : super(key: key);

  @override
  // ignore: no_logic_in_create_state
  UsernameByIdState createState() => UsernameByIdState(id: id);
}

class UsernameByIdState extends State<UsernameById> {
  final String id;
  UsernameByIdState({required this.id});
  User? user;

  @override
  void initState() {
    Future<void> getUser() async {
      User? response = await Http.getUserById(id);
      if (response != null) {
        setState(() {
          user = response;
        });
      }
    }

    getUser();
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Text(user?.user ?? " ");
  }
}

class FeedCard extends StatefulWidget {
  final Post post;
  const FeedCard(this.post, {Key? key}) : super(key: key);

  @override
  // ignore: no_logic_in_create_state
  FeedCardState createState() => FeedCardState(post: post);
}

class FeedCardState extends State<FeedCard> {
  final Post post;
  FeedCardState({required this.post});

  @override
  Widget build(BuildContext context) {
    return Padding(
        padding: const EdgeInsets.all(8),
        child: Container(
          decoration: BoxDecoration(
              border: Border.all(width: 1, color: Colors.grey),
              borderRadius: const BorderRadius.all(Radius.circular(4.0))),
          child: Column(
            children: [
              Padding(
                  padding: const EdgeInsets.all(4),
                  child: Row(
                    children: [
                      const Icon(Icons.account_circle_outlined),
                      UsernameById(post.user_id)
                    ],
                  )),
              Container(
                decoration: const BoxDecoration(
                    border: Border(
                        top: BorderSide(width: 1, color: Colors.grey),
                        bottom: BorderSide(width: 1, color: Colors.grey))),
                child: Image.network(post.post_refer),
              ),
              const Padding(
                  padding: EdgeInsets.all(4),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.end,
                    children: [Icon(Icons.rocket)],
                  ))
            ],
          ),
        ));
  }
}

class FeedWidget extends StatefulWidget {
  const FeedWidget({super.key});

  @override
  State<FeedWidget> createState() => FeedWidgetState();
}

class FeedWidgetState extends State<FeedWidget> {
  List<Post>? postsData = [];

  @override
  void initState() {
    Future<void> getPosts() async {
      List<Post>? response = await Http.getPosts();

      if (response != null) {
        setState(() {
          postsData = response;
        });
      }
    }

    getPosts();
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return postsData == null
        ? const LinearProgressIndicator()
        : ListView.builder(
            itemBuilder: (context, index) {
              return FeedCard(postsData![index]);
            },
            itemCount: postsData?.length,
          );
  }
}
