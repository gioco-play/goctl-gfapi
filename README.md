
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

# menu（互動式選單）
```shell
 goctl-gfapi menu
```
在服務的 `api/` 目錄下執行，呈現互動式選單，可選擇產生 API 服務或注入依賴。
- 範本固定取自 `gioco-play/gf-template`，無需 `--home`、`--remote` 或 `--branch` 選項
- 執行前顯示等價命令供確認
- 給人用的互動方式；自動化請直接用 `go` 或 `add-dep` 子命令

# add-dep（注入依賴到既有服務）
```shell
 goctl-gfapi add-dep --name=transaction --remote https://github.com/gioco-play/gf-template
 goctl-gfapi add-dep --name=notify,grabber --remote https://github.com/gioco-play/gf-template
```
須在服務目錄（含 `internal/`、`etc/` 的那層）下執行。`--name` 可用逗號分隔一次注入多個依賴（依序執行，非交易性）。依賴定義見 [gf-template 的 deps/ 說明](https://github.com/gioco-play/gf-template#deps)。

服務目錄下的 `makefile` 已內建對應 target，也可以直接用：
```shell
 make add-dep NAME=transaction
```
`REMOTE` 預設為 `https://github.com/gioco-play/gf-template`，可用 `make add-dep NAME=xxx REMOTE=...` 覆寫。
> 註：此 target 只會出現在新產生的服務；已存在的舊 `makefile` 需自行手動加上。

