
TRAEFIK_CONTAINER_NAME="godploy-traefik"


isTraefikRunning() {
    if [ -n "$(docker ps -q -f name=^/$TRAEFIK_CONTAINER_NAME$)" ]; then
        echo "$TRAEFIK_CONTAINER_NAME is running"
        return 0
    else
        echo "$TRAEFIK_CONTAINER_NAME is not running"
        exit 0
    fi
}

setupGodploy(){
    isTraefikRunning
    echo "continue downloading .."
}

setupGodploy
