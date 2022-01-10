# plan_adquisiciones_mid

Api intermediaria entre el cliente de plan de adquisiciones y las apis necesarios para la gestión de la información para estos mismos.
Api mid para el subsistema de plan de adquisiciones que hace parte del sistema kronos

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones

- [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
- [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno

```shell
# Ejemplo que se debe actualizar acorde al proyecto
PLAN_ADQUISICIONES_CRUD_URL = [descripción]
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y .env

### Ejecución del Proyecto

```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/plan_adquisiciones_mid

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/plan_adquisiciones_mid

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
PLAN_ADQUISICIONES_CRUD_HTTP_PORT=8080 PLAN_ADQUISICIONES_CRUD_PGURL=127.0.0.1 PLAN_ADQUISICIONES_CRUD_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile

```shell
# Implementado para despliegue del Sistema de integración continua CI.
```

### Ejecución docker-compose

```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/plan_adquisiciones_mid

#2. Moverse a la carpeta del repositorio
cd solicitudes_crud

#3. Crear un fichero con el nombre **custom.env**
touch .env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Apis Requeridas

1. [plan_adquisiciones_crud](https://github.com/udistrital/plan_adquisiciones_crud)

### Ejecución Pruebas

Pruebas unitarias

```shell
# Not Data
```

## Estado CI

| Develop | Release 1.2.0 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/plan_adquisiciones_mid/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/plan_adquisiciones_mid) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/plan_adquisiciones_mid/status.svg?ref=refs/heads/release/1.2.0)](https://hubci.portaloas.udistrital.edu.co/udistrital/plan_adquisiciones_mid) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/plan_adquisiciones_mid/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/plan_adquisiciones_mid) |

## Licencia

This file is part of plan_adquisiciones_mid

plan_adquisiciones_mid is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

plan_adquisiciones_mid is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with plan_adquisiciones_mid. If not, see https://www.gnu.org/licenses/.
