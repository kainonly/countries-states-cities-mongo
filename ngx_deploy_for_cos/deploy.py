import os
import json
from qcloud_cos import CosConfig, CosS3Client

with open('package.json') as f:
  package = json.loads(f.read())

Project = package['name']
Deploy = package['deploy']['prod']

config = CosConfig(
  Region=Deploy['region'],
  SecretId=Deploy['secretId'],
  SecretKey=Deploy['secretKey']
)
client = CosS3Client(config)

marker = ""
while True:
  response = client.list_objects(
    Bucket=Deploy['bucket'],
    Marker=marker
  )

  if 'Contents' in response.keys():
    objects = map(lambda v: {'Key': v['Key']}, response['Contents'])
    client.delete_objects(
      Bucket=Deploy['bucket'],
      Delete={
        'Object': list(objects),
        'Quiet': 'true'
      }
    )

  if response['IsTruncated'] == 'false':
    break

  marker = response['NextMarker']

print('Clear COS OK!')

for root, dir, files in os.walk('dist/' + Project):
  for file in files:
    local = os.path.join(root, file).replace('\\', '/')
    key = local.replace('dist/' + Project + '/', '')
    client.upload_file(
      Bucket=Deploy['bucket'],
      LocalFilePath=local,
      Key=key
    )
    print('Send <' + key + '> Success')

print('Sync COS OK!')
