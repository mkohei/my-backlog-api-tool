1. Set $GOPATH to your workspace.

1. Run `go get`.
```
$ go get github.com/mkohei/my-backlog-api-tool
$ cd $GOPATH/src/github.com/mkohei/my-backlog-api-tool/
```

1. Make your `conf.json`.

```
{
  "space_url": "https://xxx.backlog.jp",
  "apikey": "xxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

1. Go!
```
$ go run checkCompleted.go YOUR_BACKLOG_ISSUE_KEY
```
