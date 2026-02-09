## setting up traefik

- create certs for local https support

```
mkdir -p ./dynamic/certs
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout certs/local.key -out certs/local.crt \
  -subj "/CN=*.docker.localhost"
```

- generate password for traefik dashboard

  ```
  htpasswd -nb <admin_name> "<your_passowrd>" | sed -e 's/\$/\$\$/g'

  eg.
  htpasswd -nb roshan "rustorgo" | sed -e 's/\$/\$\$/g'
  ```

  - copy the output of `htpasswd` and past it in [compose_file](./dynamic/compose.yaml) line 61
    ```
    -"traefik.http.middlewares.dashboard-auth.basicauth.users=<your_generated_hash>"
    ```

**NOTE : use this credential when entering the trawfik dashbaord**

---

## running the server
- install dependency `make install`
- run all the service `make service-up`
- run the godploy server `make start`

- urls :

```
server = http://localhost:8080/
traefik_dashboard = https://dashboard.docker.localhost/dashboard/
```
