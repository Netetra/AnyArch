#!/bin/bash

# moduleをインストール
go get github.com/rivo/tview
go get github.com/pelletier/go-toml/v2

#ビルド
go build main
