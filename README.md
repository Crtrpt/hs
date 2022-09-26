# hs

快速启动一个 http 静态文件服务器

如果需要复杂的功能 请去用 nginx apache 或许 caddy


### 下载

https://github.com/Crtrpt/http-server/releases

下载后
可执行文件之后放到系统环境目录


给与可执行权限 (liunx)
```
chmod +x hs
```

### 运行

liunx
```
hs --addr=:8082 --root=public
```
windows
```
hs.exe --addr=:8082 --root=public
```
