## ngx_deploy_for_oss

将Angular构建资源上传至指定的阿里云对象存储

#### Setup

将文件复制到Angular项目下的`utils` 目录，修改 `package.json`，加入以下配置

```json
{
  "name": "<project_name>",
  "scripts": {
    "build": "ng build --prod",
    "deploy": "npm run build && python utils/deploy.py",
    "deploy:dependency": "pip install oss2 requests"
  },
  "deploy": {
    "prod": {
      "accessKeyId": "",
      "accessKeySecret": "",
      "endpoint": "",
      "bucket": ""
    }
  }
}
```

安装依赖

```shell script
npm run deploy:dependency
```

执行自动部署

```shell script
npm run deploy
```