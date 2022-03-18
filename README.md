# WebDAV Server

> The simplest WebDAV server.

## Usage

```shell
ssh root@master 'mkdir -p /root/bin'
scp bin/webdav-linux-armv8 root@master:/root/bin/webdav
ssh root@master 'chmod +x /root/bin/webdav; export PATH=$PATH:/root/bin; setsid webdav -P=7654 -D=/mnt > /dev/null 2>&1 &'
```
