#!/bin/bash
read -p 'MonogDB username >>> ' username
read -p 'MongoDB password >>> ' password

{
  echo "MONGO_INITDB_ROOT_USERNAME=$username"
  echo "MONGO_INITDB_ROOT_PASSWORD=$password"
  echo "MONGO_INITDB_DATABASE=pino"           # Might want to customise this?
  echo "ME_CONFIG_MONGODB_ADMINUSERNAME=$username"
  echo "ME_CONFIG_MONGODB_ADMINPASSWORD=$password"
  echo "DB_USER=$username"
  echo "DB_PASS=$password"
} >> ./.env
