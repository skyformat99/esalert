# esalert                   
[![Build Status](https://travis-ci.org/23mf/esalert.svg?branch=master)](https://travis-ci.org/23mf/esalert)
[![Go Report Card](https://goreportcard.com/badge/github.com/23mf/esalert)](https://goreportcard.com/report/github.com/23mf/esalert)

提供查询elasticsearch数据根据规则报警功能

# 超级简单的配置
```yaml
host: locahost      # es host
port: 9200          # es port 
username: elastic   # es username
password: changeme  # es password
mail:
  usernmar: "xxx"
  password: "xxx"
  smtp_host: "xxx"
  smtp_port: "345"
  send_to:
    - xxx@xx.com
    - xxx@xx.com
  from_addr: xxx@xx.com             # 显示发送出去的用户是谁
  reply_to: xxx@xx.com              # 发送出去的邮件回复给谁
  tpl_file: "/xx/xx/xx.tpl"         # go template模板文件     tpl_file与content必须存在一个
  content: "xxx{{total}}xxxx"       # go template模板字符串
  subject: "xxxx"                   # 邮件主题
rules:              # 检查规则
  - name: "xxxxx"   # 没有规则必须有一个唯一的name
    index: gateway-*  
    query:          # 查询条件，可以使任何符合 es规范的查询条件，我将转化为json字符串放入请求body想es发起请求
      exists:
        field: message.serviceException
    hits: 10        # 当符合条件大于多少条是进行报警
    interval:       # 隔多久发起一次请求，该字段会根据里面的语义信息转换时间
      day: 1
      second: 30
    time:           # 查询当前时间之前多久的数据，该字段会根据里面的语义信息转换时间
      hour: 1
      minute: 2
    alert:                                  # 报警
      - type: http                          # http报警规则
        url: http://baidu.com
      - type: mail                          # mail报警规则
        mail:                               # 该配置项参数与外层mail参数一致，该配置优先级高于外层mail配置
          tpl_file: "/xx/xx/xx.tpl"         # go template模板文件     tpl_file与content必须存在一个
          content: "xxx{{total}}xxxx"       # go template模板字符串
          subject: "xxxx"                   # 邮件主题
```

# Futures
* 完善各种规则
* 使每个运行的rule可管理并可灵活扩充
* 提供web界面
