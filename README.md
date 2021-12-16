# Log4j2Scan

![](https://img.shields.io/badge/build-passing-brightgreen)
![](https://img.shields.io/badge/golang-1.17-blue)

## ⚠️免责声明
在原有的协议中追加以下内容：

本项目禁止进行未授权商业用途

本项目禁止二次开发后进行商业用途

本项目仅面向合法授权的企业安全建设行为，在使用本项目进行检测时，您应确保该行为符合当地的法律法规，并且已经取得了足够的授权。

如您在使用本项目的过程中存在任何非法行为，您需自行承担相应后果，我们将不承担任何法律及连带责任。

在使用本项目前，请您务必审慎阅读、充分理解各条款内容，限制、免责条款或者其他涉及您重大权益的条款可能会以加粗、加下划线等形式提示您重点注意。 除非您已充分阅读、完全理解并接受本协议所有条款，否则，请您不要使用本项目。您的使用行为或者您以其他任何明示或者默示方式表示接受本协议的，即视为您已阅读并同意本协议的约束。

## 介绍
该项目由 天下大木头 与 4ra1n师傅 共同开发，感谢4ra1n师傅的许多帮助
 
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

注意：`rmi`方式payload要为 `${jndi:rmi://127.0.0.1:8001/xxxx}` 类似这种形式


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
