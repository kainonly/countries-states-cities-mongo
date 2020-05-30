# Automatic Deploy

一些自动部署脚本测试与存档

- `utils_deploy` 服务部署常用工具集，适用于debian系统，脚本将更新linux内核，安装 `python2/3` 与 `git`
- `docker_deploy` 适用debian系统的docker服务部署安装
- `docker_mirror` Docker镜像设置脚本
- `ssh_keeplive` 保持SSH客户端心跳连接
- `webhook_deploy` 网络回调钩子自动部署
- `k8s_image` 从国内源拉取Kubernetes需要的镜像组件
- `ngx_deploy_for_oss` 将Angular构建资源上传至指定的阿里云对象存储
- `ngx_deploy_for_cos` 将Angular构建资源上传至指定的腾讯云对象存储