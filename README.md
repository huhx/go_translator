A command line tool that translate chinese to english and vice versa,
it also provides the feature of querying the translation history.

> Right now, it only supports the English and Chinese languages.

## Background

When I want to translate words between chinese and english,
using the Google translation web page I need to choose the source and target language. It is less convenient to me. So I
decide to develop a command line tool
that simplify the translation work since it can identify the language automatically.

## Implementation

* `cobra`: the core of the command line
* `viper`: store the translation api key and secret
* `http`: use to call translation api

## Features

* Translation:

```shell
trans "I love coding" # 我喜欢编码
trans 我喜欢游泳 # I like swimming
```

* History record

```shell
trans -l

#                                 Query|                      Result|CreateTime
#                                  ----|                         ---|-------
#   Chinese people are very intelligent|                     中国人非常聪明|2024-04-18 17:02:10
#                              我们的爱从未消失|   Our love never disappears|2024-04-18 17:04:12
#            make the world much better|                    让世界变得更美好|2024-04-18 17:04:30
```

* Clean up history record

```shell
trans -c # successfully cleared
```

## Deploy

### prepare the viper configuration file:  ~/.go_translator.yaml
```yaml
appId: "{baidu_app_id}"
appSecret: "{baidu_app_secret}"
fileName: "go_translator_record.csv
```
https://fanyi-api.baidu.com/manage/developer

### build the exec command
```shell
go build -o trans main.go
go install trans

# or
make deploy
```
