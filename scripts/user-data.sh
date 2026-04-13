#!/bin/bash
# Output log for debugging (if something fails, check /var/log/user-data.log)
exec > >(tee /var/log/user-data.log|logger -t user-data -s 2>/dev/console) 2>&1

echo "Starting Umami installation..."

# 1. Install Docker
yum update -y
yum install -y docker
systemctl start docker
systemctl enable docker
usermod -aG docker ec2-user

# 2. Install Docker Compose v2
mkdir -p /usr/local/lib/docker/cli-plugins/
curl -SL https://github.com/docker/compose/releases/latest/download/docker-compose-linux-x86_64 -o /usr/local/lib/docker/cli-plugins/docker-compose
chmod +x /usr/local/lib/docker/cli-plugins/docker-compose
ln -s /usr/local/lib/docker/cli-plugins/docker-compose /usr/bin/docker-compose

# 3. Configure App
mkdir -p /home/ec2-user/app-umami
cd /home/ec2-user/app-umami

# 4. Create .env (Generate a random secret)
cat <<EOF > .env
POSTGRES_USER=umami
POSTGRES_PASSWORD=umami
POSTGRES_DB=umami
APP_SECRET=$(openssl rand -base64 32)
DATABASE_URL=postgresql://umami:umami@db:5432/umami
EOF

# 5. Download compose file and start services
curl -L https://raw.githubusercontent.com/IlmarLopez/umami-cloud-go/dev/docker-compose.yml -o docker-compose.yml
docker-compose up -d

echo "Installation completed."