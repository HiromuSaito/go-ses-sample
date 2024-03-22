# go-ses-samplel

### ローカル環境でgoからSESのAPIを呼び出すサンプルコード


```
# 起動
docker-compose  up -d
go run main.go

# メール送信
curl -X POST localhost:8081/mail

# ブラウザから受信したメールを確認
http://localhost:8005/
```
