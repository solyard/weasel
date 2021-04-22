![Image](weasel.png)

Telegram bot for Prometheus Alertmanager

---

**Weasel** - it's a simple Telegram Bot for Alertmanager that can recieve alerts and transfer it to telegram with templating feature and MARKDOWN support by Telegram

---
<h2>HOW TO USE</h2>

Just build image with your (or default) template and add your Telegram Bot Token recieved from @BotFather.

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
