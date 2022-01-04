# git-get

Inspired by `go get`, git clone in your source folder.

## Usage

Set configure your `~/.gitconfig`

```ini
[get]
  root = ~/src
  lowercase = false
```

Then use `git get` instead of `git clone`

```bash
$ git get git@github.com:MattKetmo/git-get.git

# will be the same as

$ git clone git@github.com:MattKetmo/git-get.git ~/src/github.com/MattKetmo/git-get
```
