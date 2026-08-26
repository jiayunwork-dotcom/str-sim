# str-sim：Go 字符串相似度核算，POST /api/compare 与子命令按选定度量比对两串

给定两串与度量（Levenshtein、Jaro-Winkler、N-gram、Soundex 等），算出距离或相似度；空串、长度不配的 Hamming 或未知度量返回错误。同一对输入在 HTTP 与子命令下分数必须一致，阈值过滤与加权合成不得 silently 改用另一套算法。

## 构建与启动

```bash
go build -o str-sim .
./str-sim                              # 启动 HTTP 服务 :8080
./str-sim jaro-winkler "Stephen" "Steven"
```

## 评测镜像

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh str-sim
```
