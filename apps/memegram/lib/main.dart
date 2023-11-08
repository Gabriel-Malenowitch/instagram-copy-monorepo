import 'package:flutter/material.dart';
import 'package:memegram/feed.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        title: 'Flutter Demo',
        theme: ThemeData(
          colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        ),
        home: const App());
  }
}

class App extends StatefulWidget {
  const App({super.key});

  @override
  State<App> createState() => AppState();
}

class AppState extends State<App> {
  int currentBodyIndex = 0;
  List<Widget> bodyOptions = <Widget>[const FeedWidget()];

  void onItemTap(int index) {
    setState(() {
      currentBodyIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Center(
          child: Text(
            "Memegram",
            // style: TextStyle(fontFamily: ""),
          ),
        ),
      ),
      body: bodyOptions.elementAt(currentBodyIndex),
      bottomNavigationBar: BottomNavigationBar(items: const [
        BottomNavigationBarItem(
            icon: Icon(Icons.feed_outlined),
            activeIcon: Icon(Icons.feed),
            label: "Feed"),
        BottomNavigationBarItem(
            icon: Icon(Icons.account_circle_outlined),
            activeIcon: Icon(Icons.account_circle),
            label: "Profile"),
      ], currentIndex: currentBodyIndex, onTap: onItemTap),
    );
  }
}
