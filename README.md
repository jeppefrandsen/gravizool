# gravizool

Small [Gravizo](http://gravizo.com) tool for encoding and decoding GitHub Markdown files based on the new [GitHub Flavored Markdown Spec](https://github.github.com/gfm/#link-destination). Makes it possible to use and view UML diagrams on GitHub and directly in pull requests for architectural reviews, module documentation etc.

It has the advantage over the [Gravizo](https://github.com/TLmaK0/gravizo) approach that it can still be rendered locally and images can be seen directly in pull requests.

### Usage

```
$ ./gravizool -h
Usage of ./gravizool:
  -b	Backup the GitHub Markdown file (default true)
  -d string
    	Decode the GitHub Markdown file
  -e string
    	Encode the GitHub Markdown file
  -f string
    	Fix the GitHub Markdown file
  -v	Print the program version
```

### Workflow

* Decode the Markdown file: `gravizool -d README.md`
* Change the Markdown file. Both [VS Code](https://code.visualstudio.com) and [Atom](http://atom.io) editor supports on the fly Markdown rendering (right click on the file and select "Markdown Preview")
* Fix the Markdown file in case of rendering issues: `gravizool -f README.md`
* Encode the Markdown file again: `gravizool -e README.md`
* Commit to GitHub

README.md files can be viewed in pull requests by pressing the "View" button under the "Files changed" tab.

### Example

If we add the [Gravizo](http://gravizo.com) example below using [plantuml](http://plantuml.com) syntax to a README.md it will not render in e.g. [Atom](http://atom.io) due to [Gravizo incompatibilities](http://www.gravizo.com/#incompatibilities):

```
<img src='https://g.gravizo.com/svg?
@startuml
class FooBar {
  String foo
  void bar()
}
@enduml
'>
```

To fix these issues use the gravizool fix (-f) option which will convert the example to the following and make it render:

```
<img src='https://g.gravizo.com/svg?;
@startuml;
class FooBar {;
  String foo;
  void bar();
};
@enduml
'>
```

Before uploading the README.md to GitHub it needs to be encoded with the gravizool encode (-e) option which will convert the example to:

```
<img src='https://g.gravizo.com/svg?%3B%0A%40startuml%3B%0Aclass%20FooBar%20{%3B%0A%20%20String%20foo%3B%0A%20%20void%20bar%28%29%3B%0A}%3B%0A%40enduml
'>
```

The example will now be rendered on GitHub and can be seen directly in e.g. the pull request:

<img src='https://g.gravizo.com/svg?%3B%0A%40startuml%3B%0Aclass%20FooBar%20{%3B%0A%20%20String%20foo%3B%0A%20%20void%20bar%28%29%3B%0A}%3B%0A%40enduml
'>

### Limitations

* Currently only works with SVG (searches for `.gravizo.com/svg?` and `enduml`)
