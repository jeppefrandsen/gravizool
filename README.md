# gravizool

Small [Gravizo](http://gravizo.com) tool for encoding and decoding GitHub Markdown files based on the new [GitHub Flavored Markdown Spec](https://github.github.com/gfm/#link-destination). Makes it possible to use and view UML diagrams in pull requests for architectural reviews, library/module documentation etc.

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

Workflow:

* Decode the Markdown file: ```gravizool -d README.md```
*  Change the Markdown file (e.g. [Atom](http://atom.io) supports on the fly Markdown rendering)
* Encode the Markdown file again: ```gravizool -e README.md```
* Commit to GitHub

### Limitations

* Currently only works with SVG (searches for ```https://g.gravizo.com/svg?```)
