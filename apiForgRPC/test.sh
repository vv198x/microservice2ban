#!/usr/bin/sh
grpcurl -plaintext -proto 2ban.proto \
    -d '{"ip": "192.168.1.55:1"}' \
    127.0.0.1:2048\
    IP2ban.IP