# Some uses of Secure Shell (SSH)

## Introduction
Secure Shell (SSH)[1](https://en.wikipedia.org/wiki/Secure_Shell) is a network protocol originally designed by Tatu Ylönen, a researcher at Helsinki University of Technology, Finland, as a replacement for earlier protocols such as TELNET, ftp and rlogin. SSH version 2(SSH-2) is a revised version of the protocol by the Internet Engineering Task Force (IETF). SSH is primarily favoured over other protocols since it uses public-key cryptography to provide a secured connection. SSH uses the client-server model. While SSH has numerous uses,in this article we are going to illustrate some other common ones such as Remote Login,  Tunneling, Port Forwarding and File Transfer. 

## Remote Login
This is by far the most popular use for SSH. Essentially, it allows users to connect remotely to a machine. Commands typed in the terminal are executed on the remote server. 

In this case, the remote server runs an SSH Server and the local machine connects to it via an SSH client. This is built into Unix-like operating systems while Windows and other platforms require third-party clients such as PuTTY.

SSH allows clients to be authenticated using password or SSH keys. SSH-key based logins are more secure.

Password logins are quite straightforward. You need to know the remote server's IP address.

In a terminal:
```shell
ssh <username>@<IP>
```
