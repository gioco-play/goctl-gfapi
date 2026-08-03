
 # 版號更新時務必使用 1.0.1 三個數字組成
 # 安裝
```shell
 go install github.com/gioco-play/goctl-gfapi@latest
 ```
 # 使用
 #### test 為資料夾
```shell
 goctl-gfapi go -api test/template.api -dir test --home template
```

# 遠端範本
```shell
 goctl-gfapi go -api template.api -dir . --remote https://github.com/gioco-play/gf-template
```
 
 # 安裝依賴 
```shell
 go mod tidy
```

# add-dep（注入依賴到既有服務）
```shell
 goctl-gfapi add-dep --name=transaction --remote https://github.com/gioco-play/gf-template
```
須在服務目錄（含 `internal/`、`etc/` 的那層）下執行。依賴定義見 [gf-template 的 deps/ 說明](https://github.com/gioco-play/gf-template#deps)。

