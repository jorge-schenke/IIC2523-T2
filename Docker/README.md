# Parte 1
## Constrrucción de imagen en docker
Luego de instalar docker, lo primero que hicimos fue clonar el repositorio público que nos entrega la documentación en el tutorial:
![image](1/1.png)
Una vez clonado el repositorio, cremamos el Dockerfile con las especificaciones del build, que eran las siguientes:<br>
```
FROM node:12-alpine
RUN apk add --no-cache python2 g++ make
WORKDIR /app
COPY . .
RUN yarn install --production
CMD ["node", "src/index.js"]
EXPOSE 3000
```
Una vez configurado, corrimos docker build para crear la imagen:
![image](1/2.png)
Con nuesta imagen ya creada, fuimos capaces de correr la app:
![image](1/3.png)
![image](1/4.png)

## Compartiendo la app:
Una vez funcionando la app, comenzamos el proceso para poder compartirla:<br>
Primero, nos creamos un usuario de Docker Hub.<br>
Después, creamos un nuevo repositorio público.<br>
Una vez creado, subimos nuestra aplicación al repo, para lo cual tuvimos que crear un tag para la imagen:
![image](1/5.png)
![image](1/6.png)
Finalmente, pudimos pushear la imagen al repo y correrla en una máquina remota:
![image](1/7.png)
![image](1/8.png)

## Múltiples contenedores + BDD
En primer lugar, creamos la network que contendría el container con la app y la BDD:
![image](1/9.png)
Luego, iniciamos una BDD MySQL, y la agregamos al network:
![image](1/10.png)
Y verificamos que la BDD se haya creado correctamente:
![image](1/11-DBList.png)
Para correr la app en el mismo container, primero debemos encontrarlo, para lo cual utilizamos un container externo:
![image](1/12-netshoot.png)
Y buscamos la IP del container que contiene la BDD:
![image](1/13-mysqlip.png)
![image](1/14-mysqlconnected.png)
![image](1/17-bothrunning.png)

## Docker compose
Primero, creamos nuestro docker-compose.yml y le agregamos el siguiente contenido:
```
version: "3.7"

services:
  app:
    image: node:12-alpine
    command: sh -c "yarn install && yarn run dev"
    ports:
      - 3000:3000
    working_dir: /app
    volumes:
      - ./:/app
    environment:
      MYSQL_HOST: mysql
      MYSQL_USER: root
      MYSQL_PASSWORD: secret
      MYSQL_DB: todos

  mysql:
    image: mysql:5.7
    volumes:
      - todo-mysql-data:/var/lib/mysql
    environment:
      MYSQL_ROOT_PASSWORD: secret
      MYSQL_DATABASE: todos

volumes:
  todo-mysql-data:
```
Una vez configurado, corremos con compose:
![image](1/18-composerun.png)
Podemos checkear que tanto la app como la BDD estén corriendo:
![image](1/19-containerrunning.png)
Y cerrar el compose:
![image](1/20-composedown.png)