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
docker service scale myapp_listerner-service=3
```
