#!/bin/bash

# Initialize ENV Variables peers
# A list of all the ports
# of the woot peers
declare -a peers=()

# Name of the WOOT cluster network
# connecting the peers
network="woot_network"

# number of peers to be provisioned
# Default 3 peers are provisioned
peers_count=$1

# Err check number of peers
# If no peers count is given defaul to 3
if [[ $peers_count -eq "" ]]; then
    peers_count=3
fi

echo "Number of peers: $peers_count"

# Exit when there is less than 1 peer
if [[ $peers_count -le 1 ]]; then
    echo "Number of peers cannot be less than 1"
    exit 255
fi

# Exit when there are more than 1000 peers
if [[ $peers_count -ge 1000 ]]; then
    echo "Number of peers cannot be more than 1000"
    exit 255
fi

# Check if port is available and then
# append to peers starting from 8000
available_port=8000
provisioned_ports_count=0

echo "Cleaning previous stale peers"
docker ps -a | awk '$2 ~ /woot/ {print $1}' | xargs -I {} docker rm -f {}
docker network rm "$network"

echo "Reserving ports for peers"

for port in {8000..9000}; do
    if [[ provisioned_ports_count -eq peers_count ]]; then
        break
    fi

    netstat -an | grep $port
    if [[ $? -ne 0 ]]; then
        peers+=($port)
        ((provisioned_ports_count++))
    fi
done

if [[ provisioned_ports_count -ne peers_count ]]; then
    echo "Unable to reserve ports for peers"
    exit 255
fi

echo "Reserved ports:" ${peers[*]}
comma_separated_peers=$(
    IFS=,
    echo "${peers[*]}"
)

# Docker create peers from peer list
# and pass PORT = peers[[i]]
echo "Provisioning WOOT Docker Cluster"

echo "Building WOOT Docker Image"
docker build -t woot -f Dockerfile .

if [[ $? -ne 0 ]]; then
    echo "Unable To Build WOOT Docker Image"
    exit 255
fi

echo "Building WOOT Cluster Network"
docker network create "$network"


for peer_index in "${!peers[@]}"; do
    declare -a peer_id_list=()

    for ((id = 0; id < $peers_count; ++id)); do
        if [[ id -ne peer_index ]]; then
            peer_id_list+=(peer-$id)
        fi
    done
    
    comma_separated_peer_id_list=$(
        IFS=,
        echo "${peer_id_list[*]}"
    )
    
    docker run -p "${peers[$peer_index]}":8080 --net $network -e "PEERS="$comma_separated_peer_id_list"" -e "NETWORK="$network"" --name="peer-$peer_index" -d woot
done

# Docker list peers on success
echo "WOOT Cluster Nodes"
docker ps | grep 'woot'
docker network ls | grep "$network"
