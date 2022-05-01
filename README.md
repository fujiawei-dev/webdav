# WebDAV Server

> The simplest WebDAV server.

## Usage

### Raspberry Pi

```shell
make linux-armv8
```

```shell
ssh root@raspberry 'mkdir -p /root/bin'
```

```shell
scp bin/webdav-linux-armv8 root@raspberry:/root/bin/webdav
```

```shell
ps ax | grep webdav | grep -v grep | sed 's/^\s*//g' | sed -e 's/\s.*//g' | xargs kill  || echo OK
```

```shell
ssh root@raspberry 'chmod +x /root/bin/webdav; export PATH=$PATH:/root/bin; setsid webdav -P=7654 -D=/mnt > /dev/null 2>&1 &'
```
