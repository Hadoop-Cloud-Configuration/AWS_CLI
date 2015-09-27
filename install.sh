sudo apt-get update
curl -sSL https://get.docker.com/ | sh
sudo usermod -aG docker ubuntu
newgrp docker
docker pull sequenceiq/ambari:1.7.0
# curl -Lo .amb j.mp/docker-ambari-170 && . .amb
docker run -d --dns 127.0.0.1 -p 8080:8080 --entrypoint /usr/local/serf/bin/start-serf-agent.sh -e KEYCHAIN= --name amb0 -h amb0.mycorp.kom sequenceiq/ambari:1.7.0 --tag ambari-server=true 
