sudo su
sed -i -e 's/#PermitRootLogin/PermitRootLogin/g' /etc/ssh/sshd_config
cat ~/.ssh/authorized_keys >> /root/.ssh/authorized_keys
