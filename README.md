# gospace

To switch to a different workspace, specify the path to the workspace.

```
$ gospace /foo/bar
```

To view the current active workspace, leave the path out.

```
$ gospace
/foo/bar
```

To make it so that no workspace is active, use the `-clear` flag.

```
$ gospace -clear
```

To make a binary accessible from every workspace, use the `-global` flag and specify the path to the binary.

```
$ gospace -global /foo/bar/baz
```

To un-global a binary, use the `-unglobal` flag.

```
$ gospace -unglobal baz
```

## Install

First, install it into one of your workspaces.

```
$ go get github.com/rynlbrwn/gospace
```

Now make `gospace` accessible from every workspace.

```
$ gospace -global `which gospace`
```

Now modify your `$GOPATH` to point to `~/.gospace`. For example, you could put this in `~/.profile` and restart your terminal.

```
export GOPATH=~/.gospace
export PATH=$PATH:$GOPATH/bin
```

And finally set your current workspace:

```
$ gospace /foo/bar
```

## How Does It Work?

The file at `~/.gospace` is a symlink to the currently active workspace. When you switch workspaces, the symlink is overwritten to point to the workspace you specified.

The `-global` flag creates a symlink in `/usr/local/bin` to the binary in your current workspace (and `--unglobal` deletes it)
