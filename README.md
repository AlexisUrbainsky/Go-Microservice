# **Configuracion del Proyecto**

### A continuacion veremos todo lo necesario para inicializar nuestro proyecto y instalaremos todas las dependencias necesarias

<br/>

# **Iniciamos el proyecto**

    C:\User\Alex\workspace>mkdir people

    go mod init people

<br/>

# **Instalamos las dependencias del proyecto**

<br/>

# **Framework para recibir y realizar peticiones HTTP**

    go get github.com/gofiber/fiber/v2

<br/>

# **ORM y Driver de la base de datos que estemos usando**

    go get -u gorm.io/gorm

    go get -u gorm.io/driver/postgres

<br/>

# **Lector de variables de entorno**

    go get github.com/joho/godotenv

<br/>

# **LOGS**

    go get golang.org/x/exp/slog

<br/>

# **CompileDaemon**

Cuando hagamos un cambio en nuestro codigo automaticamente vuelva a deployar la aplicacion evitando parar y levantar el proyecto todo el tiempo

    go get github.com/githubnemo/CompileDaemon

Una vez instalado para levantar el proyecto ejecutar el comando desde la carpeta cmd

    CompileDaemon -command="./cmd.exe"

Al ejecutar este comando se hace el build de la app compila el .exe del proyecto ejemplo gin.exe

<br/>

# **Docker**

Ejecutar los siguientes comandos para correr docker

    go build -o go-ms cmd/main.go

Comprobamos que todo funciona asi que podemos ejecutar la creacion de la imagen

    docker build . -t go-ms

Antes de levantar el contenedor de Docker deberemos configurar nuestra base de datos (PostgreSQL) para que reciba peticiones desde la ip de nuestro docker por lo que agregamos la siguiente linea al archivo pg_hba.conf

    host all all 0.0.0.0/0	scram-sha-256

**Importante** si tenemos nuestra base de datos local deberemos usar en DB_HOST la ip (ipv4) que obtenemos al hacer ipconfig en el cmd en este caso 192.168.1.85

    docker run --network="bridge" --hostname=e5f790b336d6 --mac-address=02:42:ac:11:00:02 --env=PORT=3000 --env=DB_HOST=192.168.1.85 --env=DB_NAME=gopostgres --env=DB_USER=postgres --env=DB_PASSWORD=alex --env=DB_PORT=5432 --env=LOGFILE=peopleservice.log --env=LOGPATH=./logs/ --env=PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin --env=GOLANG_VERSION=1.20.5 --env=GOPATH=/go --env=ENV=DEV --workdir=/app -p 3000:3000 --runtime=runc -d go-ms:latest

<br/>

# **Kubernetes**

Deberemos crear en la carpeta k8s los siguientes archivos

    deployment.yaml

    service.yaml

Luego ejecutaremos los siguientes comandos

    kubectl apply -f .k8s\deployment.yaml

    kubectl apply -f .k8s\service.yaml

<br/>

# **Test**

<br/>
