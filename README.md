# transfer

数据传输：客户端→转发器→服务端

```
注：password 长度必须为16, 24或者32

依次执行下述命令：
./server -l 127.0.0.1:5000 -k "password87654321"
./forward -l 127.0.0.1:4000 -t 127.0.0.1:5000 -k1 "password12345678" -k2 "password87654321"
./client -t 127.0.0.1:4000 -k "password12345678" -c "Hello world"
```