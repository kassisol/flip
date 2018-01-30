FROM scratch

LABEL com.kassisol.image.maintainer="flip@kassisol.com"
LABEL com.kassisol.image.description="An application to manage a Floating IP address on a Docker Swarm cluster"

COPY build/flip /flip

ENTRYPOINT ["/flip"]
