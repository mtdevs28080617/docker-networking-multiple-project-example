# docker-networking-multiple-project-example
Networking between multiple Docker projects


1. Run `docker network create app_bridge`
2. Run `make boot`
3. `docker exec -it app1_app bash -c "curl http://app2_app"`
4. `docker exec -it app2_app bash -c "curl http://app1_app"`