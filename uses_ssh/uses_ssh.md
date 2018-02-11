# Some uses of Secure Shell (SSH)

## Introduction
Secure Shell (SSH)[1](https://en.wikipedia.org/wiki/Secure_Shell) is a network protocol originally designed by Tatu Ylönen, a researcher at Helsinki University of Technology, Finland, as a replacement for earlier protocols such as TELNET, ftp and rlogin. SSH version 2(SSH-2) is a revised version of the protocol by the Internet Engineering Task Force (IETF). SSH is primarily favoured over other protocols since it uses public-key cryptography to provide a secured connection. SSH uses the client-server model. While SSH has numerous uses,in this article we are going to illustrate some other common ones such as Remote Login,  Tunneling and File Transfer. 

## Remote Login
This is by far the most popular use for SSH. Essentially, it allows users to connect remotely to a machine. Commands typed in the terminal are executed on the remote server. 

In this case, the remote server runs an SSH Server and the local machine connects to it via an SSH client. This is built into Unix-like operating systems while Windows and other platforms require third-party clients such as PuTTY.

SSH allows clients to be authenticated using password or SSH keys. SSH-key based logins are more secure.

Password logins are quite straightforward. You need to know the remote server's IP address.

In a terminal:
```shell
ssh <username>@<IP>
```

## Tunneling [2](https://en.wikipedia.org/wiki/Tunneling_protocol#Secure_Shell_tunneling) 

SSH Tunneling sets up a secure channel between the client and a server using SSH protocol. Services can be forwarded through this secured, encrypted channel. This is achieved by SSH port forwarding. SSH tunneling is usually used to transmit unencrypted traffic through the encrypted SSH tunnel. SSH tunnels can also be used to bypass firewall rules. There are three types of SSH port forwarding:
- Local Port Forwarding
- Remote Port Forwarding
- Dynamic Port Forwarding

### Local Port Forwarding
Local Port Forwarding is tne most common type of tunneling used. It allows you to connect to another server from your local machine. This is commonly used to bypass certain firewall rules. 

For example, let's assume that `supinfo.com` is blocked on our local network. We could create a tunnel through another server on another network then access it locally.

In a terminal:


```shell
ssh -L 8080:www.supinfo.com:80 user@vps
```
Pointing browser at `http://localhost:8080/` would load the `supinfo.com` website through a tunnel to the remote vps server.

Another use of local port forwarding is to remotely access services that are accessible only when using `127.0.0.1` as the hostname. For example, phpMyAdmin which is an administration tool for MySQL and MariaDB. For instance, we could access phpMyAdmin remotely to manage a MySQL database hosted on amazon AWS. 

```shell
ssh -N -L 8888:127.0.0.1:80 <username>@<AWS instance public IP or DNS>
```
Then, phpMyAdmin would be available on local machine at `http://127.0.0.1:8888/phpmyadmin/`

### Remote Port Forwarding
Remote Port Forwarding, also called Reverse SSH, is used to provide access to a service running locally on the machine from a remote network. A simple example would be sharing access to web app running locally to a colleague with deploying the app. Typically, we would need to know our public IP address and configure NAT and firewall rules to allow someone to connect remotely to services running locally on our machine. 

Remote Port Forwarding allows us to give remote access to local services via a third publicly accessible server. To enable it, the `/etc/ssh/sshd_config` file needs to be modified by adding :

```shell
GatewayPorts yes
```

For example, let's assume we want to share a webapp running locally on `port 8889` to the outside world via a publicly available server(illustrated using a local server at 192.168.100.5). 

```shell
ssh -R 8080:localhost:8889 user@192.168.100.5
```

Loading `192.168.100.5:8080`  in the browser would load the web app running locally at `localhost:8889`

### Dynamic Port Forwarding

Dynamic Port Forwarding turns the SSH client into a SOCKS 4/5 proxy.Socket Secure (SOCKS)[3](https://en.wikipedia.org/wiki/SOCKS) is a protocol that allows information exchange between a client and a server through a proxy server. It allows traffic to bypass Internet filtering to access blocked content. The Onion Router (Tor) offers a SOCKS server interface to its clients.[4](https://www.torproject.org/docs/faq.html.en#TBBSocksPort)

