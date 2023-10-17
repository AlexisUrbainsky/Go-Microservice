# **Config**

<br/>

# **Dependencies**

<br/>

# **HTTP Framework**

    go get github.com/gofiber/fiber/v2

<br/>

# **ORM & Driver**

    go get -u gorm.io/gorm

    go get -u gorm.io/driver/postgres

<br/>

# **Read Env Variables**

    go get github.com/joho/godotenv

<br/>

# **LOGS**

    go get golang.org/x/exp/slog

<br/>

# **CompileDaemon**

    go get github.com/githubnemo/CompileDaemon

    CompileDaemon -command="./cmd.exe"

<br/>

# **Docker**

Build Image

    go build -o go-ms cmd/main.go

Execute

    docker build . -t go-ms

Config Ip to connect PostgreSQL with the container

    host all all 0.0.0.0/0	scram-sha-256

**Importante** if our database is local, we set the DB_HOST variable with your local ip, from ipconfig, in this case 192.168.1.85

    docker run --network="bridge" --hostname=e5f790b336d6 --mac-address=02:42:ac:11:00:02 --env=PORT=3000 --env=DB_HOST=192.168.1.85 --env=DB_NAME=gopostgres --env=DB_USER=postgres --env=DB_PASSWORD=alex --env=DB_PORT=5432 --env=LOGFILE=peopleservice.log --env=LOGPATH=./logs/ --env=PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin --env=GOLANG_VERSION=1.20.5 --env=GOPATH=/go --env=ENV=DEV --workdir=/app -p 3000:3000 --runtime=runc -d go-ms:latest

<br/>

# **Kubernetes**

Execute the commands

    kubectl apply -f .k8s\deployment.yaml

    kubectl apply -f .k8s\service.yaml

<br/>

# **Test**

<br/>
