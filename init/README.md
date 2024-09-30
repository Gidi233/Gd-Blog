# Systemd 配置、安装和启动

- [Systemd 配置、安装和启动](#systemd-配置安装和启动)
	- [1. 前置操作](#1-前置操作)
	- [2. 创建 Gd-Blog systemd unit 模板文件](#2-创建-gd-blog-systemd-unit-模板文件)
	- [3. 复制 systemd unit 模板文件到 sysmted 配置目录](#3-复制-systemd-unit-模板文件到-sysmted-配置目录)
	- [4. 启动 systemd 服务](#4-启动-systemd-服务)

## 1. 前置操作

1. 创建需要的目录 

```bash
sudo mkdir -p /data/Gd-Blog /opt/Gd-Blog/bin /etc/Gd-Blog /var/log/Gd-Blog
```

2. 编译构建 `Gd-Blog` 二进制文件

```bash
make build # 编译源码生成 Gd-Blog 二进制文件
```

3. 将 `Gd-Blog` 可执行文件安装在 `bin` 目录下

```bash
sudo cp _output/platforms/linux/amd64/Gd-Blog /opt/Gd-Blog/bin # 安装二进制文件
```

4. 安装 `Gd-Blog` 配置文件

```bash
sed 's/.\/_output/\/etc\/Gd-Blog/g' configs/Gd-Blog.yaml > Gd-Blog.sed.yaml # 替换 CA 文件路径
sudo cp Gd-Blog.sed.yaml /etc/Gd-Blog/ # 安装配置文件
```

5. 安装 CA 文件

```bash
make ca # 创建 CA 文件
sudo cp -a _output/cert/ /etc/Gd-Blog/ # 将 CA 文件复制到 Gd-Blog 配置文件目录
```

## 2. 创建 Gd-Blog systemd unit 模板文件

执行如下 shell 脚本生成 `Gd-Blog.service.template`

```bash
cat > Gd-Blog.service.template <<EOF
[Unit]
Description=APIServer for blog platform.
Documentation=https://github.com/Gidi233/Gd-Blog/blob/master/init/README.md

[Service]
WorkingDirectory=/data/Gd-Blog
ExecStartPre=/usr/bin/mkdir -p /data/Gd-Blog
ExecStartPre=/usr/bin/mkdir -p /var/log/Gd-Blog
ExecStart=/opt/Gd-Blog/bin/Gd-Blog --config=/etc/Gd-Blog/Gd-Blog.yaml
Restart=always
RestartSec=5
StartLimitInterval=0

[Install]
WantedBy=multi-user.target
EOF
```

## 3. 复制 systemd unit 模板文件到 sysmted 配置目录

```bash
sudo cp Gd-Blog.service.template /etc/systemd/system/Gd-Blog.service
```

## 4. 启动 systemd 服务

```bash
sudo systemctl daemon-reload && systemctl enable Gd-Blog && systemctl restart Gd-Blog
```
