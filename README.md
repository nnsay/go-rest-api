# Fork 自[go-rest-api](https://hub.docker.com/r/chentex/go-rest-api)

仅供个人测试使用

# ChangeLog

- 修改 Dockerfile 跑不通的地方

# 使用方式

## 编译和运行

- 本地直接使用

```bash
go build rest-api.go
./rest-api
```

- Docker

```bash
podman build . -t rest:latest
podman run -p 8080:8080 localhost/rest:latest
```

## 测试

```bash
curl http://localhost:8080/test

{"color":"yellow","message":"This is a Test","notify":"false","message_format":"text"}
```

```bash
curl http://localhost:8080/hola/world

{"color":"yellow","message":"Hola world","notify":"false","message_format":"text"}
```

## 镜像

```bash
podman login docker.io

podman build . -t rest:latest
podman tag localhost/rest docker.io/nnsay/rest:latest
podman push docker.io/nnsay/rest:latest
```
