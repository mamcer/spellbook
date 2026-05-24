# grafana

## install

[https://grafana.com/docs/grafana/latest/setup-grafana/installation/debian/](https://grafana.com/docs/grafana/latest/setup-grafana/installation/debian/)

signing key

```bash
sudo apt-get install -y apt-transport-https
sudo apt-get install -y software-properties-common wget
sudo wget -q -O /usr/share/keyrings/grafana.key https://apt.grafana.com/gpg.key
```

stable repository 

```bash
echo "deb [signed-by=/usr/share/keyrings/grafana.key] https://apt.grafana.com stable main" | sudo tee -a /etc/apt/sources.list.d/grafana.list
sudo apt update
```

oss version

```bash
sudo apt install grafana
```

```bash
start grafana server
```

[https://grafana.com/docs/grafana/latest/setup-grafana/start-restart-grafana/](https://grafana.com/docs/grafana/latest/setup-grafana/start-restart-grafana/)

```bash
sudo systemctl daemon-reload
sudo systemctl start grafana-server
sudo systemctl status grafana-server
sudo systemctl status grafana-server
```

login

```
http://localhost:3000/
```

> default user/password: admin/admin

## first dashboards

[https://www.youtube.com/watch?v=YUabB_7H710](https://www.youtube.com/watch?v=YUabB_7H710)
