#!/bin/bash

# 同步本地 /root/dst 目录到远程服务器
rsync -avz --delete --exclude="mods" -e "ssh -p 58320" /root/dst/ root@117.72.100.4:/root/dst/