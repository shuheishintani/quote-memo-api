# QuoteMemo

## 概要

QuoteMemoは書籍の引用を管理・シェアできるWebサービスです。

## デモページ
[https://quote-memo-client.vercel.app](https://quote-memo-client.vercel.app)

※サーバー停止中



## 使用技術

バックエンド
  - Go v1.16
  - Gin v1.7.1
  - Gorm v1.21.9
  - PostgreSQL v13

フロントエンド
  - TypeScript v4.2.3
  - React v17.0.2
  - Next.js v10.2.2
  - Chakra UI (@chakra-ui/react v1.4.2)

認証
  - Firebase Authentication


開発環境
  - Docker/docker-compose

 CIツール
  - GitHub Actions


インフラ
  - Vercel
  - GCP(App Engine/Cloud SQL)

## システム構成図

<img src="./architecture.png" width="600px" height="700px">

## ER図

<img src="./data-model.png" width="800px" height="600px">

## 機能一覧
- Twitterログイン
- 書籍検索
- 引用の作成・更新・削除
- 公開設定
- タグ付け
- 複数のタグを指定して検索
- 書籍に関連する引用一覧を表示
- 他のユーザーが公開した引用一覧を表示
- 他のユーザーが公開した引用をお気に入りに追加
- ページネーション
- ダークモード
- 退会処理

## 作成動機

## アピールポイント

## 苦労した点

## 課題

