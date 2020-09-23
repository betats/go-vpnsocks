# go-vpnsocks

socks proxy via openvpn

## config
cp sample.env .env
vim .env

## profile
mkdir profile
copy openvpn profiles

## run
from profile dir, ramdom use one profile
docker-compose up

## check
ex: curl -4 -x socks5://127.0.0.1:1080 ifconfig.io
