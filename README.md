# curso_go_api_twiter

#paso 1
Tener instalado go

#crear la carpeta "Go" y dentro crear estas 3 carpetas
src/
bin/
kpg/
y dentro de src/ se crea otra carpeta con el nombre github.com y dentro crear otra carpeta con el nombre de usuario de github ejemplo:
src/github.com/luisesteban91/

#paso 2 
Dentro de la carpeta del usuario github crear el proyecto ejemplo:
src/github.com/luisesteban91/CURSO_GO_API_TWITER

#paso 4
configurar el main.go

#paso 5
posicionarte a nivel raiz de tu terminal y ejecutar el comando: "run go main.go" ejemplo:
go run main.go 
go: downloading github.com/gorilla/mux v1.8.0
go: downloading github.com/rs/cors v1.8.0
go: downloading github.com/joho/godotenv v1.3.0
go: downloading go.mongodb.org/mongo-driver v1.7.0
go: downloading golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
go: downloading github.com/dgrijalva/jwt-go v3.2.0+incompatible
go: downloading github.com/pkg/errors v0.9.1

#paso 8
Crear archivo .env con la variable
GO_PASSWORD_BD:123456

#paso 7
compilar el proyecto con el comando: "go build main.go"
