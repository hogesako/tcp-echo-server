# command
```
echo "hoge" | nc localhost 2701
```

# build
```
docker buildx create --name multi-arch-builder --driver docker-container --use
docker buildx build --platform linux/amd64,linux/arm64 -t hogesako/tcp-echo-server --push .
```