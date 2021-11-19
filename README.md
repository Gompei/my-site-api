# my-site-api

[![ci](https://github.com/Gompei/my-site-api/actions/workflows/ci.yml/badge.svg)](https://github.com/Gompei/my-site-api/actions/workflows/ci.yml)

## リポジトリ概要

自サイト用の記事管理用API

## 環境

- go 1.17
- lambda
- Dynamo DB

## API概要

|  エンドポイント  |  概要  | HTTPメソッド |
| ---- | ---- | ---- |
|  /api/article/search  |  全文検索APIにリクエスト(未実装)  | GET |
|  /api/article/list  | 見出し用記事データを件数分取得  | GET |
|  /api/article/physical-delete  |  記事データの物理削除  | DELETE |
|  /api/article  |  記事データのCRUD処理  | GET,POST,PUT,DELETE |
