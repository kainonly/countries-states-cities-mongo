#!/bin/sh

# 更新源

sudo apt update

# 更新已安装的包与内核

sudo DEBIAN_FRONTEND=noninteractive apt -y upgrade

# 安装 python

sudo apt -y install python python-pip python3 python3-pip
