# google_mid
API MID de servicios de google.

### Variables de Entorno
```shell
GOOGLE_MID_HTTP_PORT: [Puerto de ejecución API]
GOOGLE_MID_GMAIL_CLIENT_ID: [Client Id de la cuenta que realizará el envío de correos]
GOOGLE_MID_GMAIL_CLIENT_SECRET: [Secreto de la cuenta que realizará el envío de correos]
GOOGLE_MID_GMAIL_ACCESS_TOKEN: [Access token de la cuenta que realizará el envío de correos]
GOOGLE_MID_GMAIL_REFRESH_TOKEN: [Refresh token de la cuenta que realizará el envío de correos]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/google_mid

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/google_mid

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
GOOGLE_MID_HTTP_PORT=8080 GOOGLE_MID_GMAIL_CLIENT_ID=127.0.0.1 GOOGLE_MID_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile
```shell
# Implementado para despliegue del Sistema de integración continua CI.
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/google_mid

#2. Moverse a la carpeta del repositorio
cd google_mid

#3. Crear un fichero con el nombre **custom.env**
touch .env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# Not Data
```

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/google_mid/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/google_mid) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/google_mid/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/google_mid) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/google_mid/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/google_mid) |



## Licencia

This file is part of google_mid.

google_mid is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

google_mid is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with google_mid. If not, see https://www.gnu.org/licenses/.
