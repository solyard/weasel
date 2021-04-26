![Image](weasel.png)

Telegram bot for Prometheus Alertmanager

---

**Weasel** - it's a simple Telegram Bot for Alertmanager that can recieve alerts and transfer it to telegram with templating feature and MARKDOWN support by Telegram

---

**STATUS**:

![Image](https://img.shields.io/github/workflow/status/solyard/weasel/Go?label=Go%20Compile%20&style=for-the-badge)

![Image](https://img.shields.io/github/workflow/status/solyard/weasel/ci?color=blue&label=Docker%20Build&style=for-the-badge)

[![Go Report Card](https://goreportcard.com/badge/github.com/solyard/weasel)](https://goreportcard.com/report/github.com/solyard/weasel)

---
<h2>HOW TO USE</h2>

Just build image with your (or default) template and add your Telegram Bot Token recieved from [@BotFather](https://t.me/botfather).

Run image with `docker` or on your Linux / Windows system and add config to your `Prometheus Alertmanager` to start recieve some messages

```yaml
          global:
          resolve_timeout: 5m
        route:
          group_wait: 30s
          group_interval: 5m
          repeat_interval: 3h
          receiver: 'telegram'
        receivers:
        - name: 'telegram'
          webhook_configs:
          - url: 'http://localhost:8081/api/v1/alerts/{chat_id}'
            send_resolved: true
```

You can `run` image with simple `docker run` command
```
docker run -p 8081:8081 solard/weasel
```
If you want to add additional config and (or) template just mount it into `docker image`
```
docker run -p 8081:8081 -v $(pwd)/config.yaml:/config/telegram.conf.yaml -v $(pwd)/my_custom_template.tmpl:/confing/default.tmpl solard/weasel
```

For testing your installation you simply can use `curl` with `POST` method:

```bash
curl -XPOST -H "Content-type: application/json" -d '{"alerts":[{"annotations":{"description":"SOME TEXT DATA","summary":"TEST ALERT"},"generatorURL":"http:\/\/alert:8080\/api\/v1\/15821810008956981301\/8832311346543396454\/status","labels":{"alertgroup":"rules","alertname":"CRITICAL TEST","instance":"my-test-instance","severity":"critical"},"startsAt":"2021-04-20T05:15:04.65109161Z"},{"annotations":{"description":"SOME TEXT DATA","summary":"TEST LERT"},"generatorURL":"http:\/\/alert:8080\/api\/v1\/15821810008956981301\/15760119835279596093\/status","labels":{"alertgroup":"rules","alertname":"WARNING TEST","instance":"my-test-instance","severity":"warning"},"startsAt":"2021-04-20T05:14:34.648183556Z"}],"commonAnnotations":{"summary":"SOME ANNOTATIONS"},"commonLabels":{"alertgroup":"rules","instance":"my-test-instance"},"externalURL":"http:\/\/alert:9093","groupKey":0,"groupLabels":{},"receiver":"telegram","status":"resolved","version":0}' 'localhost:8081/api/v1/alert/{chat_id}'
```

**How alert's looks**:

![image](https://user-images.githubusercontent.com/8751732/116091391-7062ac80-a6ad-11eb-8645-86f2750d1d21.png)

