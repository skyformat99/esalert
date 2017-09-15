# esalert
提供查询elasticsearch数据根据规则报警功能

# 超级简单的配置
```yaml
host: locahost      # es host
port: 9200          # es port 
username: elastic   # es username
password: changeme  # es password
rules:              # 检查规则
  - index: gateway-*  
    query:          # 查询条件，可以使任何符合 es规范的查询条件，我将转化为json字符串放入请求body想es发起请求
      exists:
        field: message.serviceException
    hits: 10        # 当符合条件大于多少条是进行报警
    interval: 60    # 隔多久发起一次请求
    alert:          # 报警
      type: http    # 报警规则
      url: http://baidu.com
```
