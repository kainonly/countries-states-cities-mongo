## ssh_keeplive

保持SSH客户端心跳连接

```shell script
# 添加执行权限
chmod +x ./ssh_keeplive.py
# 运行脚本
./ssh_keeplive.py <input_path> <output_path>
```

- `input_path` 输入文件路径，例如：`/etc/ssh/sshd_config`
- `output_path` 输出文件路径，一般与 `input_path` 相同