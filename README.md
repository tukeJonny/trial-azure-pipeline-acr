# Azure Pipeline & Azure Container Registry Trial

[![Build Status](https://dev.azure.com/ne250143/test/_apis/build/status/tukeJonny.trial-azure-pipeline-acr?branchName=master)](https://dev.azure.com/ne250143/test/_build/latest?definitionId=3&branchName=master)

## 概要

Azure DevOps + Azure ACRの試験リポジトリです。

pushするイメージは、大きくサーバとして動くものと、静的ファイルとして配信されるものの２つがあります。

前者は `prepare-ci.sh` を実行してazure-pipeline.ymlを更新する必要があります

一方で後者は `put-staticfile.sh` を実行して、静的ファイルを正しい場所に設置する必要があります

## サーバとして動作するコンテナイメージの追加

```
$ mkdir problem_dir # 問題関連のファイルを設置するディレクトリ作成
$ vim problem_dir/Dockerfile # Dockerfile作成
$ vim ...
$ ./scripts/prepare-ci.sh # pushする前に、azure-pipelines.ymlを更新
```

## 静的ファイルの追加

```
$ ./scripts/put-staticfile.sh /path/to/staticfile
```

