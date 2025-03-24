# yadsh


Yadsh is an implementation of [YandexDisk REST API](https://yandex.ru/dev/disk-api/doc/ru/) in Golang.

For usage:
1) Create OAuth token: https://yandex.ru/dev/disk-api/doc/ru/concepts/quickstart#oauth
2) Export as env var: `export OAUTH_TOKEN="...."`
3) Run app / use package:
```
go build -o ./yadsharez ./cmd/yadsharez/main.go
./yadsharez [FULL_FILE_PATH]
```

Another option via `go install`:
```
go install github.com/ep4sh/yadsharez/cmd/yadsharez@latest
```

Please feel free to contibute :>



## Books
I'm in love with books. If you want to thank me, just help me to buy books from the list

[![buy-me-a-book](https://img.shields.io/badge/Amazon-Buy%20me%20a%20book-important)](https://www.amazon.com/hz/wishlist/ls/3NSSXQK5CTS8N?ref_=wl_share)
