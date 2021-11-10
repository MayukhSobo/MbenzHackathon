docker build -t mbenz-routing -f Dockerfile .
docker run --rm --name mbenz-route -p 9999:9999 --network mbenz_default mbenz-routing
