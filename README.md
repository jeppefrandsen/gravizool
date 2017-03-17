# gravizool

Small [gravizo.com](http://gravizo.com) tool for encoding and decoding GitHub Markdown files based on the new [GitHub Flavored Markdown Spec](https://github.github.com/gfm/#link-destination).

Makes it possible to use and view UML diagrams in pull requests for architectural reviews, library/module documentation etc.

### Usage

```
$ ./gravizool -h
Usage of ./gravizool:
  -b	Backup GitHub Markdown file before encode/decode (default true)
  -d string
    	Decode the given GitHub Markdown file
  -e string
    	Encode the given GitHub Markdown file
```

Typical workflow:

* Decode the Markdown file: ```gravizool -d README.md```
*  Change the Markdown file (e.g. the [Atom](http://atom.io) editor supports interactive Markdown rendering)
* Encode the Markdown file again: ```gravizool -e README.md```
* Push to GitHub

### Limitations

* Currently only works with SVG (searches for ```https://g.gravizo.com/svg```)
* End of URL must be ```enduml)``` (no newline, semicolon etc.)
