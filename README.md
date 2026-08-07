# RayleaBot Game Guide Plugin

`raylea.game-guide` 提供游戏角色攻略查询，并使用 RayleaBot 托管的渲染模板生成卡片。插件以独立 Go 模块和独立 GitHub Release 发布。

## 目录结构

- `cmd/game-guide/`：只负责启动进程。
- `internal/plugin/`：角色检索、事件处理和测试。
- `internal/assets/`：由 Go 嵌入的角色数据；构建时映射为 artifact 的 `data/characters.json`。
- `templates/`：宿主渲染模板。
- `tools/build/`：统一组装后端、数据与模板。

## 本地联调

将本仓库路径写入 RayleaBot 根目录下被忽略的 `plugin-workspace.local.json`，并运行：

```powershell
$env:RAYLEA_PLUGIN_DEV = 'watch'
$env:RAYLEA_SERVER_RELOAD = 'watch'
node scripts/start-dev.mjs
```

启动脚本会连接本地 Go SDK、构建当前平台 artifact，并通过离线 `plugin dev-sync` 原子同步到 `plugins/installed/`。模板与数据文件由同一 artifact 清单校验。

## 发布

推送 `v*` 标签后，工作流使用固定 SDK 引用构建 Windows x64、Linux x64 和 macOS arm64 ZIP，并创建 GitHub Release。插件目录仓库随后记录各平台产物摘要并发布签名目录。

License: MIT
