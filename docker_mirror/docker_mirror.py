#!/usr/bin/python3
import sys
import subprocess
import os
import json

try:
    if len(sys.argv) < 2:
        exit("Please add operating parameters!")

    mirrors = sys.argv[1:]
    daemon_path = '/etc/docker/daemon.json'
    if not os.path.exists(daemon_path):
        with open(daemon_path, mode='w') as f:
            f.write(json.dumps({
                "registry-mirrors": mirrors
            }))
    else:
        with open(daemon_path, mode='r') as f:
            content = json.loads(f.read())
            content['registry-mirrors'] = mirrors

        with open(daemon_path, mode='w') as f:
            f.write(json.dumps(content))

    subprocess.run(['systemctl', 'restart', 'docker'])
    print('docker mirrors setup successfully!')
except Exception as error:
    print(error)
