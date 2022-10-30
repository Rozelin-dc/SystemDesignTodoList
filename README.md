
# System Design Todo List
## 起動方法
```bash
$ docker-compose up --build
```

[localhost:3050](http://localhost:3050)にアクセス

## ディレクトリ構成
```txt
SystemDesignTodoList
├── client: クライアントサイドの実装。詳細はclientディレクトリ内のREADME参照
├── docs: swagger置き場
├── domain: domain層
│   ├── model: model層。各種構造体の定義
│   └── repository: repository層。各種インターフェースの定義
├── handler: handler層。エンドポイントの中身の実装
├── infra: infra層。repository層で定義したインターフェースの実装
├── middleware: 各エンドポイントへのアクセス権をチェック
├── sql: DB関連ファイル
├── util: その他必要な実装(バリデーションの実装)
└── README.md: このファイル
```
