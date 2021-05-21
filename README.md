# QuoteMemo

## デモページ
[https://quote-memo-client.vercel.app](https://quote-memo-client.vercel.app)
※サーバー停止中

## 概要

QuoteMemoは書籍の引用を管理・シェアできるWebサービスです。

## 使用技術

- バックエンド
  - Go v1.16
  - Gin v1.7.1
  - Gorm v1.21.9
  - PostgreSQL v13
- フロントエンド
  - TypeScript v4.2.3
  - React v17.0.2
  - Next.js v10.2.2
  - Chakra UI (@chakra-ui/react v1.4.2)
- 開発環境
  - Docker/docker-compose
- CIツール
  - GitHub Actions
- インフラ
  - Vercel
  - GCP(App Engine/Cloud SQL)

## システム構成図

![architecture](./architecture.png)

## ER図

![data-model](./data-model.png)

