# gravizool

Small [Gravizo](http://gravizo.com) tool for encoding and decoding GitHub Markdown files based on the new [GitHub Flavored Markdown Spec](https://github.github.com/gfm/#link-destination). Makes it possible to use and view UML diagrams in pull requests for architectural reviews, module documentation etc.

### Usage

```
$ ./gravizool -h
Usage of ./gravizool:
  -b	Backup GitHub Markdown file before encode/decode (default true)
  -d string
    	Decode the GitHub Markdown file
  -e string
    	Encode the GitHub Markdown file
  -f string
    	Fix the GitHub Markdown file
```

Workflow:

* Decode the Markdown file: `gravizool -d README.md`
*  Change the Markdown file. The [Atom](http://atom.io) editor supports on the fly Markdown rendering (right click on the file and select *Markdown Preview*)
* Fix the Markdown file in case of rendering issues: `gravizool -f README.md`
* Encode the Markdown file again: `gravizool -e README.md`
* Commit to GitHub

README.md files can be viewed in pull requests by pressing the *View* button under the *Files changed* tab)

### Limitations

* Currently only works with SVG (searches for `https://g.gravizo.com/svg?`)
