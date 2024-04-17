A command line tool that translate chinese to english and vice versa,
it also provides the feature of querying the translation history.

## Background

When I want to translate word between chinese and english,
using the web page I need to choose the source and target language.
it is less convenient to me. So I want to develop a command line tool
that simplify the translation work. It can identify the language automatically.

## Implementation

* `cobra`: the core of the command line
* `viper`: store the translation api key and secret
* `http`: use to call translation api

## Deploy

```shell
go build -o trans main.go
go install trans
```

## Features
* Translation: 
```shell
trans "I love coding" # 我喜欢编码
trans 我喜欢游泳 # I like swimming

```
