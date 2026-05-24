# docker

## build

```bash
docker build -t hello:latest .
```

## run

```bash
docker run -it -p 5000:80 --name upictures-web upictures
```
> port 5000 host, 80 docker

```bash
# start MySQL
docker run -p 3306:3306 --name cookbook -e MYSQL_ROOT_PASSWORD=dev -d mysql:8.4.4
# create database
docker exec -it cookbook mysql -uroot -p -e "create database cookbook;"
```

```bash
docker run -it -p 5000:80 --name upictures-web -v /Users/mario/Desktop/upictures/pictures:/app/wwwroot/pictures upictures
```
> volume

```bash
docker exec -it [container-name] /bin/bash
```
> interactive /bin/bash  

```bash
docker run -p 80:8080 --name caca -d --rm
```
> detached, delete after run

```bash
docker run --restart unless-stopped ...
```
> the container restart automatically, even after system shutdown

## ip address

```bash
docker inspect <container-name> | grep "IPAddress"
```

## example static web 

```bash
vim Dockerfile
```

```docker
FROM nginx:alpine
COPY . /usr/share/nginx/html/
  
docker build -t mywebsite .
docker run -d -p 8080:80 mywebsite
```

## example image based on existing container

create a new container based on an existing to change execution parameters

```bash
docker ps
docker commit [hash] aldebaran
docker run --name=aldebaran -dit -v ~/home/aldebaran:/home/ubuntu -p 5000:80 aldebaran
```

## develop with docker

```bash
docker run --rm -it -p 3000:3000 your-image
```
> then it is deleted when stopped (--rm)  
> for multiple versions, use tags for your images (e.g., your-image:v1, your-image:v2), and always use --rm when running them.

regularly remove stopped containers with:

```bash
docker container prune
```

regularly remove unused images with:

```bash
docker image prune
```

# Docker Multi-Environment Setup

This project supports multiple Docker environments for development and production.

## File Structure

```
├── docker-compose.dev.yml     # Development environment
├── docker-compose.prod.yml    # Production environment
├── .env.dev                   # Development environment variables
├── .env.prod                  # Production environment variables
├── .env.example               # Example environment file
├── dev.sh                     # Development helper script
├── prod.sh                    # Production helper script
└── DOCKER_ENVIRONMENTS.md     # This documentation
```

## Quick Start

### Development Environment

1. **Setup environment variables:**
   ```bash
   cp .env.example .env.dev
   # Edit .env.dev with your development settings
   ```

2. **Start development environment:**
   ```bash
   ./dev.sh up
   ```

3. **View logs:**
   ```bash
   ./dev.sh logs
   ```

4. **Stop development environment:**
   ```bash
   ./dev.sh down
   ```

### Production Environment

1. **Setup environment variables:**
   ```bash
   cp .env.example .env.prod
   # Edit .env.prod with your production settings
   ```

2. **Start production environment:**
   ```bash
   ./prod.sh up
   ```

3. **View logs:**
   ```bash
   ./prod.sh logs
   ```

4. **Stop production environment:**
   ```bash
   ./prod.sh down
   ```

## Helper Scripts

Both `dev.sh` and `prod.sh` support the following commands:

- `up` - Start the environment (default)
- `down` - Stop the environment
- `restart` - Restart the environment
- `logs` - Show logs from all services
- `build` - Build/rebuild images
- `shell` - Open shell in API container
- `clean` - Remove containers and volumes

## Manual Docker Compose Commands

If you prefer to use Docker Compose directly:

### Development
```bash
# Start
docker-compose -f docker-compose.dev.yml --env-file .env.dev up -d

# Stop
docker-compose -f docker-compose.dev.yml --env-file .env.dev down

# View logs
docker-compose -f docker-compose.dev.yml --env-file .env.dev logs -f
```

### Production
```bash
# Start
docker-compose -f docker-compose.prod.yml --env-file .env.prod up -d

# Stop
docker-compose -f docker-compose.prod.yml --env-file .env.prod down

# View logs
docker-compose -f docker-compose.prod.yml --env-file .env.prod logs -f
```

## Environment Variables

### Required Variables

- `DB_HOST` - Database host
- `DB_PORT` - Database port
- `DB_USER` - Database username
- `DB_PASSWORD` - Database password
- `DB_CONNECTION_STRING` - Full database connection string
- `MEDIA_PATH` - Path to media files

### Optional Variables

- `COMPOSE_PROJECT_NAME` - Docker Compose project name

## Key Differences Between Environments

### Development (`docker-compose.dev.yml`)
- Uses `ASPNETCORE_ENVIRONMENT=Development`
- Container names have `-dev` suffix
- Source code mounted for hot reload
- Logs directory mounted for easier debugging
- Uses local database connection

### Production (`docker-compose.prod.yml`)
- Uses `ASPNETCORE_ENVIRONMENT=Production`
- Container names have `-prod` suffix
- Optimized for production deployment
- Uses production database connection
- Production media path

## Best Practices

1. **Never commit `.env.dev` or `.env.prod`** - These files contain sensitive information
2. **Use `.env.example`** as a template for new environments
3. **Test both environments** before deploying to production
4. **Use different ports** if running both environments simultaneously
5. **Keep environment files secure** and limit access

## Troubleshooting

### Common Issues

1. **Environment file not found:**
   - Ensure `.env.dev` or `.env.prod` exists
   - Check file permissions

2. **Database connection issues:**
   - Verify database credentials in environment file
   - Ensure database server is running
   - Check network connectivity

3. **Port conflicts:**
   - Both environments use the same ports by default
   - Modify ports in compose files if running simultaneously

4. **Permission issues:**
   - Ensure helper scripts are executable: `chmod +x dev.sh prod.sh`
   - Check Docker daemon permissions

### Getting Help

- Check container logs: `./dev.sh logs` or `./prod.sh logs`
- Inspect containers: `docker ps -a`
- Check environment variables: `docker-compose config`

## db wait service 

```yaml
db-wait:
  container_name: nostalgia-db-wait-dev
  image: alpine:latest
  command: >
    sh -c "
      apk add --no-cache mysql-client &&
      until mysql -h ${DB_HOST} -P ${DB_PORT} -u ${DB_USER} -p${DB_PASSWORD} -e 'SELECT 1'; do
        echo 'Waiting for database...';
        sleep 2;
      done;
      echo 'Database is ready!'
    "
  networks:
    - nostalgia-network
```