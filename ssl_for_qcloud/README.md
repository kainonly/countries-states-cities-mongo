## ssl for qcloud

腾讯云SSL证书管理简化

## 使用

首先需要安装腾讯云开发者工具套件

```shell script
pip install --upgrade tencentcloud-sdk-python
```

将该包复制到实际项目中使用

- **search()** 搜索相关SSL证书
    - **domain** `str` 域名名称
- **delete()** 删除指定的SSL证书
    - **id** `str` 证书ID
- **upload()** 上传SSL证书
    - **pub_key** `str` 公钥内容
    - **pri_key** `str` 私钥内容