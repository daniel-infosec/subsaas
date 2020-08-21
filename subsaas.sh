#/bin/bash
image=$(docker inspect --type=image subsaas:latest)
image=`echo "${image}" | head -1`
if [ $image = "[]" ]; then
    docker build . -t subsaas:latest
fi

docker run subsaas:latest subsaas $1 $2
