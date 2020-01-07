#!/usr/bin/python3
import sys
from os import path
import re
import subprocess

if len(sys.argv) < 3:
    exit("Please add operating parameters!")

input_path = sys.argv[1]
output_path = sys.argv[2]

if path.exists(input_path) != True or output_path == None:
    exit("File path is incorrect or does not exist!")

try:
    with open(input_path, mode='r') as f:
        content = f.read()

    raws = []
    exists_interval = False
    exists_count_max = False
    for line_text in content.split('\n'):
        if re.search('(|\W+)#(|\W+)ClientAliveInterval\W+[0-9]+', line_text):
            continue
        if re.search('(|\W+)ClientAliveInterval\W+[0-9]+', line_text):
            raws.append("ClientAliveInterval 15")
            exists_interval = True
            continue
        if re.search('(|\W+)#(|\W+)ClientAliveCountMax\W+[0-9]+', line_text):
            continue
        if re.search('(|\W+)ClientAliveCountMax\W+[0-9]+', line_text):
            raws.append("ClientAliveCountMax 45")
            exists_count_max = True
            continue
        raws.append(line_text)

    if exists_interval == False:
        raws.append("ClientAliveInterval 15")
    if exists_count_max == False:
        raws.append("ClientAliveCountMax 45")

    with open(output_path, mode='w') as f:
        f.write('\n'.join(raws))

    subprocess.run(['systemctl', 'restart', 'sshd'])
    print('ssh_keeplive keeplive set successfully!')
except Exception as e:
    print(e)
