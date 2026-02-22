# Docker Swarm & Service SDK Assignments (Homelab & VPS Perspective)

These assignments are designed to help you understand Docker Swarm as an end-user deploying to their own VPS or homelab. Instead of typing `docker service ...` in the terminal, your goal is to write Go scripts using the Docker SDK (`github.com/moby/moby/client`) that accomplish the exact same things. 

This builds the foundation: once you know how a developer uses Swarm natively via the SDK, wrapping it into your own product later becomes much easier.

---

## Level 1: Swarm Basics (Easy)
*Focus: Getting the cluster running and deploying simple stateless apps.*

### 1. Initialize the Swarm
**CLI Equivalent:** `docker swarm init`
*   **Task:** Write a script that checks the current Swarm status of the Docker engine. If it's inactive, initialize a new single-node Swarm manager.
*   **SDK Focus:** `client.Info`, `client.SwarmInit`, `swarm.InitRequest`
*   **Guiding Question:** *When you inspect the `Info` struct, what is the difference between `Swarm.LocalNodeState` being `active` vs `pending`?*

### 2. The "Hello World" Service
**CLI Equivalent:** `docker service create --name web -p 8080:80 nginx:alpine`
*   **Task:** Deploy a single-replica Nginx service. Map port 80 inside the container to port 8080 on your host's routing mesh.
*   **SDK Focus:** `client.ServiceCreate`, `swarm.ServiceSpec`, `swarm.EndpointSpec`, `swarm.PortConfig`
*   **Guiding Question:** *In `PortConfig`, there is a `PublishMode` field. What is the difference between `PublishModeIngress` and `PublishModeHost` in a multi-node cluster?*

### 3. Service Discovery & Inspection
**CLI Equivalent:** `docker service ls` and `docker service inspect web`
*   **Task:** Write a script that lists all services. Find the service named "web", retrieve its ID, and print out the image it's running and its published ports.
*   **SDK Focus:** `client.ServiceList`, `filters.NewArgs`, `client.ServiceInspectWithRaw`
*   **Guiding Question:** *The SDK often requires a Service ID, not a name, for operations. How do you efficiently use `filters` in `ServiceList` to resolve a name to an ID without pulling down the entire list of services?*

### 4. Teardown
**CLI Equivalent:** `docker service rm web`
*   **Task:** Write a function that takes a service name, finds its ID, and completely removes it.
*   **SDK Focus:** `client.ServiceRemove`

---

## Level 2: Day-2 Operations (Medium)
*Focus: State, configurations, updates, and networking.*

### 5. Scaling Up
**CLI Equivalent:** `docker service scale web=3`
*   **Task:** Write a script that takes the "web" service and scales it to 3 replicas.
*   **SDK Focus:** `client.ServiceUpdate`, `swarm.ServiceSpec.Mode.Replicated`
*   **Guiding Question:** *`ServiceUpdate` requires you to pass the current `Version` of the service. Why does the Docker API use this "optimistic locking" mechanism? What happens if someone else updates the service right before your script does?*

### 6. Rolling Updates & Rollbacks
**CLI Equivalent:** `docker service update --image nginx:latest --update-delay 10s web`
*   **Task:** Update the "web" service to use a new image tag. Configure the update so it happens one container at a time, pausing for 10 seconds between each.
*   **SDK Focus:** `swarm.UpdateConfig`, `client.ServiceUpdate`
*   **Guiding Question:** *If the new image is broken and crashes immediately, what does Swarm do by default? How would you configure `UpdateConfig` to automatically rollback if the update fails?*

### 7. Overlay Networks & Stateful Volumes
**CLI Equivalent:** `docker network create -d overlay mynet` & `docker service create --network mynet --mount type=volume,source=dbdata,target=/var/lib/postgresql/data postgres`
*   **Task:** 
    1. Create an overlay network.
    2. Deploy a Postgres service attached to this network. 
    3. Mount a named volume to ensure the database data persists if the container restarts. Do NOT publish the database port to the host.
*   **SDK Focus:** `client.NetworkCreate`, `client.ServiceCreate`, `swarm.NetworkAttachmentConfig`, `mount.Mount`
*   **Guiding Question:** *Since you didn't publish the Postgres port, how would another service (like a backend API) connect to it? What hostname does Docker's internal DNS assign to the database container?*

### 8. Managing Secrets and Configs
**CLI Equivalent:** `echo "my-db-pass" | docker secret create db_pass -` & `docker service create --secret db_pass ...`
*   **Task:** 
    1. Create a Docker Secret via the SDK containing a dummy database password.
    2. Create a Docker Config via the SDK containing a basic JSON config file.
    3. Deploy a service (e.g., Alpine) that mounts both the secret and the config. Write a script to verify the service started.
*   **SDK Focus:** `client.SecretCreate`, `client.ConfigCreate`, `swarm.SecretReference`, `swarm.ConfigReference`
*   **Guiding Question:** *Where are Secrets mounted inside the container filesystem by default? Can a container modify the secret file once it's mounted?*

---

## Level 3: Advanced Swarm Control (Hard)
*Focus: Scheduling constraints, troubleshooting, and edge proxying.*

### 9. Node Labels & Placement Constraints
**CLI Equivalent:** `docker node update --label-add disk=ssd <node-id>` & `docker service create --constraint 'node.labels.disk == ssd' ...`
*   **Task:**
    1. List the nodes in your Swarm and add a custom label `disk=ssd` to your current node.
    2. Deploy a Redis service, but add a placement constraint so it is *only* allowed to run on nodes with the label `disk=ssd`.
*   **SDK Focus:** `client.NodeList`, `client.NodeUpdate`, `swarm.Placement`
*   **Guiding Question:** *What happens to the service if you deploy it with a constraint, but no nodes in the cluster match that constraint? Does the service creation fail, or does something else happen?*

### 10. Task Monitoring & Troubleshooting (The Crash Loop)
**CLI Equivalent:** `docker service ps <service>` and `docker service logs <service>`
*   **Task:** 
    1. Deploy a service using an image that immediately crashes (e.g., `alpine` with command `["sh", "-c", "exit 1"]`).
    2. Write a script that lists the "Tasks" (individual container instances) for this service. 
    3. Detect that the tasks are failing, and fetch the logs of the failed tasks to print to the terminal.
*   **SDK Focus:** `client.TaskList`, `swarm.TaskState`, `client.ServiceLogs`
*   **Guiding Question:** *A Service is the desired state; a Task is the actual container trying to reach that state. If a container crashes, Swarm spins up a new Task. How do you query the API to find only the Tasks that are currently in a `FAILED` state?*

### 11. The Homelab Entrypoint (Traefik Integration)
**CLI Equivalent:** Deploying Traefik to the edge, and using labels on other services to route traffic.
*   **Task:** 
    1. Deploy Traefik as a service, publishing ports 80 and 8080. You'll need to mount the Docker socket `/var/run/docker.sock` so Traefik can read Swarm events.
    2. Deploy a `whoami` service. Instead of publishing ports on `whoami`, attach it to the same network as Traefik and apply Traefik routing labels directly to the `ServiceSpec.Labels`.
*   **SDK Focus:** `mount.Mount` (Bind mount), `swarm.ServiceSpec.Labels`
*   **Guiding Question:** *When integrating Traefik with Docker Swarm, Traefik looks at labels. Do you place the `traefik.http.routers...` labels on the Container spec, or the Service spec? Why does Swarm architecture mandate one over the other?*