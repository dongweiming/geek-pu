# 微信小程序 - 极客电玩小铺

## Introduction

功能
- [x] 游戏名称搜索
- [x] 游戏区域筛选
- [x] 微信通知订阅

## Getting start

##### 1. Check out the repository

```bash
❯ git clone https://github.com/dongweiming/geek-pu
```

##### 2. 使用微信开发者工具打开该项目


##### 3. 启动服务

```bash
❯ cd geek-pu
❯ go run manage.go initdb  # 初始化数据库
❯ cp config.toml.tmpl config.toml # 根据配置模板修改数据库配置
# 导入数据
❯ mysql -u USER -p DATABASE < database/data.sql
# 如果是自己的数据，可以通过manage.go创建，还可以加可选的 -v 增加商品说明
# go run manage.go add -t 马里奥派对 -c pd.jpg -r 2018-10-05 -s 9.1 -d 30245974 -a 美版 -l 简体中文 -p Switch
# go run manage.go update -p 2900 -i 14 -t 4
❯ go run server.go  # 启动服务, 当然也可以使用只对外用Nginx
```

## Experience it

扫描下方小程序码可以进行体验

<p align="center">
<img src="https://github.com/dongweiming/geek-pu/blob/master/assets/code.jpg">
</p>
