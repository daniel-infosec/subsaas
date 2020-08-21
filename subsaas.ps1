$image = docker inspect --type=image subsaas:latest
if ($image -eq "[]"){
    docker build . -t subsaas:latest
}

$param1 = $args[0]
$param2 = $args[1]

docker run subsaas:latest subsaas $param1 $param2