# How to run

```bash
podman build . -t opencv-go-podman
podman run --device /dev/video0 opencv-go-podman:latest
```


## Go into the container

```bash
podman ps
podman exec -it <container id or name> /bin/bash
```