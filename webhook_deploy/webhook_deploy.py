#!/usr/bin/python3
import subprocess
import configparser
import os
import json


class Config(configparser.ConfigParser):
    def optionxform(self, optionstr):
        return optionstr


try:
    subprocess.run(['apt', 'update'])
    subprocess.run(['apt', '-y', 'install', 'webhook'])
    subprocess.run(['systemctl', 'enable', 'webhook'])

    if not os.path.exists('/opt/webhook'):
        os.makedirs('/opt/webhook')

    if not os.path.exists('/opt/webhook/scripts'):
        os.makedirs('/opt/webhook/scripts')

    if not os.path.exists('/opt/webhook/hook.json'):
        with open('/opt/webhook/hook.json', mode='w') as f:
            f.write(json.dumps([]))

    file_service = '/lib/systemd/system/webhook.service'
    config = Config()
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
    print('webhook successfully deployed!')
except Exception as error:
    print(error)
