# Flip

An application to manage a Floating IP address on a Docker Swarm cluster.

The setup will consist of a Docker Swarm cluster with `flip` enabled on each managers. The `flip` daemon controls on which host the Floating IP will be configured. The Floating IP is only assigned to one host at a time (providing active / passive failover). In the event that the service fails on the active Docker host, Docker Swarm will relocate it on another member of the cluster and `flip` daemon will assign the Floating IP on that member.

## Get started

To enable `flip` for a given service, add the label `flip.enabled: "true"` to the service configuration.

There is 2 types of datasource to get the Floating IP from.

### File

Create a yaml file with the following content:

```
---
floating_ip:
  address: '192.168.12.226'
  netmask: '255.255.255.0'
  gateway: '192.168.12.1'
```

Then run the docker command:

```bash
docker run -d -v /var/run/docker.sock:/var/run/docker.sock -v /opt/flip/config.yml:/tmp/config.yml --net host --cap-add=NET_ADMIN --restart=always --name flip kassisol/flip:x.x.x -d file -o /tmp/config.yml -D
```

### Metadata
#### Kassisol

The metadata server can be downloaded from https://github.com/kassisol/metadata. Once running and provisioned with data, run the docker command to start flip.

```
docker run -d -v /var/run/docker.sock:/var/run/docker.sock --net host --cap-add=NET_ADMIN --restart=always --name flip kassisol/flip:x.x.x -d kassisol -o "url=http://metadata.example.com:8080;itype=public;index=0" -D
```

## User Feedback

### Issues

If you have any problems with or questions about this application, please contact us through a [GitHub](https://github.com/kassisol/flip/issues) issue.
