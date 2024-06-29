# 短縮URLアプリ

```bash
# redis起動
redis-server 

# サーバー起動
go run main.go

# テスト
go test -v <path>
```

# 動作確認
```bash
# 短縮URL作成・取得
curl -X POST http://localhost:9808/create-short-url -H "Content-Type: application/json" -d '{"long_url": "https://github.com/redis/go-redis","user_id": "e0dba740-fc4b-4977-872c-d360239e6b10"}'
```