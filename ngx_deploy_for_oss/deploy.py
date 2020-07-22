import os
import json
from oss2 import Auth, Bucket, ObjectIterator

with open('package.json') as f:
    package = json.loads(f.read())

Project = package['name']
Deploy = package['deploy']['prod']

auth = Auth(Deploy['accessKeyId'], Deploy['accessKeySecret'])
bucket = Bucket(auth, Deploy['endpoint'], Deploy['bucket'])

for obj in ObjectIterator(bucket):
    bucket.delete_object(obj.key)
    print('Delete <' + obj.key + '> Success')

print('Clear OSS OK!')

for root, _, files in os.walk('dist/' + Project):
    for file in files:
        local = os.path.join(root, file).replace('\\', '/')
        key = local.replace('dist/' + Project + '/', '')
        bucket.put_object_from_file(key, local)
        print('Send <' + key + '> Success')

print('Sync OSS OK!')
