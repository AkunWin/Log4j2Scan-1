# Log4j2Scan

![](https://img.shields.io/badge/build-passing-brightgreen)
![](https://img.shields.io/badge/golang-1.17-blue)

## 介绍
 
这是什么？

一款无须借助`dnslog`且完全无害的log4j2检测工具，解析`RMI`和`LDAP`协议实现

特点：
- 由`Golang`编写，直接运行可执行文件进行检测
- 动态渲染数据到`web`页面，方便观察结果

![](https://github.com/KpLi0rn/Log4j2Scan/blob/master/img/001.png)

## 使用

命令：`./Log4j2Scan -p port`

```text
    __                __ __  _ _____
   / /   ____  ____ _/ // / (_) ___/_________ _____
  / /   / __ \/ __ `/ // /_/ /\__ \/ ___/ __ `/ __ \
 / /___/ /_/ / /_/ /__  __/ /___/ / /__/ /_/ / / / /
/_____/\____/\__, /  /_/_/ //____/\___/\__,_/_/ /_/
 /____/    /___/
    coded by 天下大木头 & 4ra1n
[+] [09:34:02] use port 8001
[+] [09:34:02] start result http server
[+] [09:34:02] start fake reverse server
|------------------------------------|
|--Payload: ldap/rmi://your-ip:port--|
|------------------------------------|
```

只需要将`Payload`设置为：`ldap://your-ip:port`或`rmi://your-ip:port`

然后访问：`localhost:8888`即可动态查看最新结果
