# rego

## inspiration

there is one folder that seems to get hold all the files in our computer. from executables to plain text files to whole folders. it holds it all.

that's nice in the short run. easy to get stuff done that way. but over time you might start finding it harder to find files which would lead to over reliance on the search bar of the file manager.

that's not bad until a friend sees that folder and gives you the "do you really live like this?"
stare.

in that instance, you would start to think of ways to arrange the files into proper folders so that you would look saner. and feel saner too.

that's where `rego` comes in. it quickly re-arranges the files in that folder into sane sub-directories that eases access and the over all experience.

## usage

- download a relevant version for your operating system from the [releases](https://github.com/spobly/rego/releases/tag/v0.1.0) page

- install it
- run it with

```
$ rego -p path/to/folder
```

if you want the files to be moved to a folder in the base OS path e.g /home/(your-name) then use

```
$ rego -p path/to/folder -g
```

if you are in the folder that needs improvement simply use

```
$ rego -p .
```

or

```
$ rego -p . -g
```

depending on your preference.

**note**: you might have files that it is not familiar with so if you encounter that kind of error simply create a [new issue](https://github.com/spobly/rego/issues/new) so I can add.

feature requests and improvements are also welcomed especially those that improve the performance because it is really slow when dealing with large file.

thanks for reading.

happy organizing.
