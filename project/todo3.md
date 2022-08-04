# Todo 3: Deploy with Docker Swarm

> Goal: Manage multiple nodes with Docker Swarm

Initialize a Swarm
```console
docker swarm init
```

```console
docker stack deploy -c swarm.yml myapp
```

To scale up a service
```console
docker service scale myapp_logger-service=3
```
To update a service
```console
docker service update --image misterpuffin/logger-service:1.0.1 myapp_logget-service
```

## Task: To update logger-service image from v1.0.0 to v1.0.1 with little downtime


