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

with open(input_path, mode='r') as f:
    content = f.read()

content = re.sub('\n(|\W+)#(|\W+)ClientAliveInterval\W+[0-9]+', '', content)
content = re.sub('\n(|\W+)#(|\W+)ClientAliveCountMax\W+[0-9]+', '', content)
content += '''
ClientAliveInterval 15
ClientAliveCountMax 45
'''

with open(output_path, mode='w') as f:
    f.write(content)

subprocess.run(["systemctl", "restart", "sshd"])
