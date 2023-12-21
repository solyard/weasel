# ⚠⚠⚠ This project will freeze soon! Weasel will reborn and return in new skin ... ⚠⚠⚠



![Image](weasel.png)

Bot for Prometheus Alertmanager

---

**Weasel** - it's a simple Bot for Alertmanager that can recieve alerts and transfer it to telegram with templating feature and MARKDOWN support

---

**STATUS**:

![GO Compile](https://img.shields.io/github/actions/workflow/status/solyard/weasel/go.yml?style=flat-square&label=GO%20Compile)

![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/solyard/weasel/ci.yml?style=flat-square&label=Docker%20Build)

[![Go Report Card](https://goreportcard.com/badge/github.com/solyard/weasel)](https://goreportcard.com/report/github.com/solyard/weasel)

---

**SUPPORTED MESSENGERS**
- [x] Telegram
- [ ] Slack

---
<h2>HOW TO USE</h2>

Just build/pull image and set your (or default) template and add your Telegram Bot Token recieved from [@BotFather](https://t.me/botfather).

> Don't forget to setup the env variable (TELEGRAM_BOT_TOKEN)

Bot supports command to get your chatID and threadID if you are using Telegram Topics:

```bash
/chatID - will reply to your message with ChatID
/threadID - will reply to your message with ThreadID (Topic ID)
```

Run image with `docker` or on your Linux / Windows system and add config to your `Prometheus Alertmanager` to start recieve some messages

Example for **VictoriaMetrics Operator (VMAlertmanager)**

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMAlertmanager
metadata:
  name: vmalertmanager
  namespace: monitoring
spec:
  replicaCount: 1
  configSecret: alertmanager-config
  configRawYaml: |
        global:
          resolve_timeout: 5m
        route:
          group_wait: 5s
          group_interval: 1m
          repeat_interval: 15m
          receiver: 'telegram'
        receivers:
        - name: 'telegram'
          webhook_configs:
          - url: 'http://prometheus-weasel:8081/api/v1/alert/{chat_id}'
            send_resolved: true
```

If you are inspired by new Telegram Topics you can use this construction to send Alert in specified Topic:

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMAlertmanager
metadata:
  name: vmalertmanager
  namespace: monitoring
spec:
  replicaCount: 1
  configSecret: alertmanager-config
  configRawYaml: |
        global:
          resolve_timeout: 5m
        route:
          group_wait: 5s
          group_interval: 1m
          repeat_interval: 15m
          receiver: 'telegram'
        receivers:
        - name: 'telegram'
          webhook_configs:
          - url: 'http://prometheus-weasel:8081/api/v1/alert/{chat_id}/{topic_id}'
            send_resolved: true
```

You can `run` image with simple `docker run` command
```
docker run -p 8081:8081 solard/weasel -e TELEGRAM_BOT_TOKEN="<secret token>"
```
If you want to add additional config and (or) template just mount it into `docker image`
```
docker run -p 8081:8081 -v $(pwd)/my_custom_template.tmpl:/confing/default.tmpl -e TELEGRAM_BOT_TOKEN="<secret token>" solard/weasel
```

For testing your installation you simply can use `curl` with `POST` method:

```bash
curl -XPOST -H "Content-type: application/json" -d '{"alerts":[{"annotations":{"description":"SOME TEXT DATA","summary":"TEST ALERT"},"generatorURL":"http:\/\/alert:8080\/api\/v1\/15821810008956981301\/8832311346543396454\/status","labels":{"alertgroup":"rules","alertname":"CRITICAL TEST","instance":"my-test-instance","severity":"critical"},"startsAt":"2021-04-20T05:15:04.65109161Z"},{"annotations":{"description":"SOME TEXT DATA","summary":"TEST LERT"},"generatorURL":"http:\/\/alert:8080\/api\/v1\/15821810008956981301\/15760119835279596093\/status","labels":{"alertgroup":"rules","alertname":"WARNING TEST","instance":"my-test-instance","severity":"warning"},"startsAt":"2021-04-20T05:14:34.648183556Z"}],"commonAnnotations":{"summary":"SOME ANNOTATIONS"},"commonLabels":{"alertgroup":"rules","instance":"my-test-instance"},"externalURL":"http:\/\/alert:9093","groupKey":0,"groupLabels":{},"receiver":"telegram","status":"resolved","version":0}' 'localhost:8081/api/v1/alert/{chat_id}'
```

**How alert's looks**:

![image](https://user-images.githubusercontent.com/8751732/116091391-7062ac80-a6ad-11eb-8645-86f2750d1d21.png)

