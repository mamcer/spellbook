# nginx

[https://www.nginx.com/blog/deploying-nginx-nginx-plus-docker/](https://www.nginx.com/blog/deploying-nginx-nginx-plus-docker/)

## basic nginx docker container

```bash
docker run --name mynginx1 -p 80:80 -d nginx
```

## sharing files from host

```bash
docker container, files: /usr/share/nginx/html, configuration:  /etc/nginx
host, files: /var/www, configuration: /var/nginx/conf

    docker run --name mynginx2 --mount type=bind source=/var/www,target=/usr/share/nginx/html,readonly --mount type=bind,source=/var/nginx/conf,target=/etc/nginx/conf,readonly -p 80:80 -d nginx
```

## see logs

```bash
docker logs container-name
```

## Create a docker static server

`docker-compose.yml`

```yaml
web:
  restart: always
  image: nginx
  ports:
    - "88:80"
  volumes:
    - /home/mario/Desktop/www:/usr/share/nginx/html
```

Then execute

```bash
docker-compose up
```

on localhost:88 you should be able to see any index.html (for example) you added in /home/mario/Desktop/www

## Multiple sites on the same IP address

For this to be done right different sites should be added to the DNS server with the same IP

In our case we will configure hosts file with two sites: caca.com and coco.com

```
/etc/hosts
```

```
192.168.100.100 caca.com
192.168.100.100 coco.com
```

in Nginx

```bash
sudo vim /etc/nginx/nginx.conf
```

```
server {
   listen 80;
   root /home/mario/Desktop/www3;
   index index.html index.htm;
   server_name _;
   location / {
       try_files $uri $uri/ =404;
   }
}
```
```
server {
listen 80;
   root /home/mario/Desktop/www;
index index.html index.htm;
   server_name caca.com;
   location / {
       try_files $uri $uri/ =404;
   }
}
```
```
server {
   listen 80;
   root /home/mario/Desktop/www2;
   index index.html index.htm;
   server_name coco.com;
   location / {
       try_files $uri $uri/ =404;
   }
}
```

files in www3 should be de default (192.168.100.100) www2 (coco.com) and www (caca.com)

Reload nginx

```bash
sudo nginx -s reload
```

Review any error

```bash
sudo tail -n 20 /var/log/nginx/error.log
```

You can test it with any browser or with curl

```bash
curl coco.com --resolve coco.com:80:192.168.100.100
```