## prerequisites

- `docker` v29.x
- `docker-compose`
- `openssl`

## setting up local env

- create certs for local https support

```
mkdir -p ./dynamic/certs
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout ./dynamic/certs/local.key -out ./dynamic/certs/local.crt \
  -subj "/CN=*.godploy.localhost"
```

- add `127.0.0.1  *.godploy.localhost` in new line to your `/etc/hosts` file
- run `make install` to install the dependencies
- run `make img-build` to build the docker images

## running the services

- run `make setup` to setup traefik
- run `make dev` to run dev containers
- you can access services at
  - Traefik dashboard : `https://traefik.godploy.localhost:8443/` (to access the dashboard username : `godploy`, password : `godploy`)
  - Godploy web : `https://web.godploy.localhost`
  - Godploy server : `https://server.godploy.localhost`

## watch the services

- run `make web-logs` to watch the web service logs
- run `make server-logs` to watch the server service logs
- run `make traefik-logs` to watch the traefik service logs

## stopping the development environment

- run `make services-rm` to stop and remove the services
