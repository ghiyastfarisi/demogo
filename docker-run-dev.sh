docker rm -f demogo-dev
docker run -dit --restart unless-stopped --name demogo-dev -v $(pwd):/app -p 8080:8080 demogo:dev