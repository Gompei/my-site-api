# my-site-api

[![ci/cd](https://github.com/Gompei/my-site-api/actions/workflows/cicd.yml/badge.svg)](https://github.com/Gompei/my-site-api/actions/workflows/cicd.yml)

## リポジトリ概要

自分のWEBサイト用の記事管理用APIです。

## 環境

- go 1.17
- lambda
- Dynamo DB

## API概要

|  エンドポイント  |  概要  | HTTPメソッド |
| ---- | ---- | ---- |
|  /article/list  | 見出し用記事データを件数分取得  | GET |
|  /article  |  記事データ取得 | GET |

## その他

### 関連リポジトリ

- フロント: [Gompei/my-site](https://github.com/Gompei/my-site)
- api: [Gompei/my-site-api](https://github.com/Gompei/my-site-api)
- インフラ: [Gompei/my-site-terraform](https://github.com/Gompei/my-site-terraform)