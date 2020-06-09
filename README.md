# traefik-custom-forward-auth

traefik + patched ForwardAuth middleware

## Patched ForwardAuth middleware

`traefik.http.middlewares.{コンテナ名}.forwardauth.forwardRequest=true` を指定することで、  
オリジナルのリクエストを forward できるようにパッチを当ててみたやつ。

## Getting Started

run:
```bash
docker-compose up -d

docker-compose logs -f
```

request: randauth によりランダムでリクエストが認証成功または失敗し、 `X-User-Id` ヘッダーにダミーの username が付与される。
```bash
curl http://localhost/ -H 'Host: whoami.example.com' -d '{"test_body": "hogehoge"}'

curl http://localhost/anything -H 'Host: httpbin.example.com' -d '{"test_body": "hogehoge"}'
```

dashboard:
```
http://localhost:8080/dashboard
```

traces:
```
http://localhost:16686/
```
