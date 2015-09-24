sudo apt-get update
curl -sSL https://get.docker.com/ | sh
sudo docker pull sequenceiq/ambari:1.7.0
sudo usermod -aG docker ubuntu
newgrp docker
curl -Lo .amb j.mp/docker-ambari-170 && . .amb
amb-deploy-cluster 3
