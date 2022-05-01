APP_NAME=webdav

#pgrep "$APP_NAME" | xargs kill
pkill "$APP_NAME"

cp bin/$APP_NAME-linux-amd64 /root/bin/$APP_NAME

chmod +x /root/bin/$APP_NAME

export PATH=$PATH:/root/bin

setsid $APP_NAME -P=7654 -D=/ > /dev/null 2>&1 &
