
# ScafGinAPI

https://github.com/kodaimura/scaf-gin の派生プロジェクトで、API開発に特化したテンプレートです。  
認証系のAPIをデフォルトで実装しています。

### 必要なツール
- **Docker**
- **make**

---

## 🚀 使い方

### インストール
[webscaf](https://github.com/kodaimura/webscaf) を使って、簡単にセットアップできます。  
Githubのテンプレート機能やcloneでも、そのまま利用できます。  

### 起動
以下のコマンドでデフォルトアプリを起動できます。

```bash
make up
```

ログイン・サインアップ機能付きの**Gin API**が立ち上がります。  
http://localhost:8000/api

---

## 🧰 コマンド一覧（Makefile）

```bash
make up        # コンテナの起動
make down      # コンテナの停止と破棄
make reup      # コンテナの停止、破棄、再起動
make build     # コンテナの再ビルド
make stop      # コンテナの停止のみ
make in        # appコンテナ内にbashで入る
make log       # コンテナのログを監視
make ps        # コンテナの状態を確認
```

### 環境切り替え

異なる環境で動作させたい場合、`ENV`変数を指定してください。
指定なしの場合は dev で起動します。
```bash
make up ENV=prod      # 本番環境で起動
make up ENV=dev       # 開発環境で起動
```

### 環境変数設定

コンテナ内の環境変数（例：データベース設定や認証設定など）を設定したい場合は、`.env`ファイルを利用します。  
環境ごとに異なる設定を`.env.prod`や`.env.dev`として管理できます。
