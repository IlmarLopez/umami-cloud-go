#!/bin/bash

# Redirect output to log file for debugging
exec > >(tee /var/log/user-data.log|logger -t user-data -s 2>/dev/console) 2>&1

echo "Starting Umami installation..."

# 1. Install Docker
yum update -y
yum install -y docker git jq
systemctl start docker
systemctl enable docker
usermod -aG docker ec2-user

# 2. Install Docker Compose v2
mkdir -p /usr/local/lib/docker/cli-plugins/
curl -SL https://github.com/docker/compose/releases/latest/download/docker-compose-linux-x86_64 -o /usr/local/lib/docker/cli-plugins/docker-compose
chmod +x /usr/local/lib/docker/cli-plugins/docker-compose
ln -s /usr/local/lib/docker/cli-plugins/docker-compose /usr/bin/docker-compose

# 3. Clone the Umami repository
cd /home/ec2-user
git clone https://github.com/umami-software/umami.git app-umami
cd app-umami

# 4. FETCH FROM SECRETS MANAGER (Zero Hardcoding!)
SECRET_JSON=$(aws secretsmanager get-secret-value --region us-east-2 --secret-id umami-db-credentials --query SecretString --output text)

# 5. Extract specific credentials using jq
DB_URL=$(echo $SECRET_JSON | jq -r '.DATABASE_URL')
DB_USER=$(echo $SECRET_JSON | jq -r '.POSTGRES_USER')
DB_PASS=$(echo $SECRET_JSON | jq -r '.POSTGRES_PASSWORD')
DB_NAME=$(echo $SECRET_JSON | jq -r '.POSTGRES_DB')
APP_SEC=$(echo $SECRET_JSON | jq -r '.APP_SECRET')

# 6. Dynamically generate the .env file for Docker
cat <<EOF > .env
DATABASE_URL=$DB_URL
APP_SECRET=$APP_SEC
POSTGRES_USER=$DB_USER
POSTGRES_PASSWORD=$DB_PASS
POSTGRES_DB=$DB_NAME
EOF

# 7. Start services
docker-compose up -d

echo "Installation completed successfully."