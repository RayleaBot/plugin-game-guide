# RayleaBot Game Guide Plugin

`raylea.game-guide` 提供游戏角色攻略查询，并使用 RayleaBot 托管的渲染模板生成卡片。插件以独立 Go 模块和独立 GitHub Release 发布。

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
