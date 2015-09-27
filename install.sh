sudo apt-get update
curl -sSL https://get.docker.com/ | sh
sudo usermod -aG docker ubuntu
newgrp docker
docker pull sequenceiq/ambari:1.7.0
curl -Lo .amb j.mp/docker-ambari-170 && . .amb
amb-deploy-cluster 3
