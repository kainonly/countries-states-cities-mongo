#!/usr/bin/python3
import subprocess
import configparser
import os

# subprocess.run(["apt", "update"])
#
# subprocess.run(["apt", "-y", "install", "webhook"])

file_service = '/lib/systemd/system/webhook.service'
config = configparser.ConfigParser()
config.read(file_service)

config.set(
    'Service',
    'ExecStart',
    '/usr/bin/webhook -nopanic -hotreload -port 5985 -urlprefix projects -hooks /opt/webhook/hook.json'
)

with open(file_service, mode='w') as f:
    config.write(f)

subprocess.run(['systemctl', 'daemon-reload'])
subprocess.run(['systemctl', 'restart', 'webhook'])
subprocess.run(['systemctl', 'enable', 'webhook'])
