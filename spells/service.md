# service

## configure as a service

```bash
sudo vim /etc/systemd/system/nmovies.service
```

add the following content:

```
[Unit]
Description=NMovies

[Service]
WorkingDirectory=/home/mario/src/nmovies
ExecStart=/home/mario/src/nmovies/main
Restart=always
RestartSec=10
User=mario
Environment=DISPLAY=:0

[Install]
WantedBy=multi-user.target
```

then:

```bash
sudo systemctl enable nmovies.service
sudo systemctl start nmovies.service
sudo systemctl status nmovies.service
```

view logs:

```bash
journalctl -u nmovies -b
```

## remove a service

```bash
sudo systemctl stop nmovies.service
sudo systemctl disable nmovies.service
sudo rm /etc/systemd/system/nmovies.service
sudo systemctl daemon-reload
```