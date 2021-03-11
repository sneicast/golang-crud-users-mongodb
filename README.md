# golang-crud-users-mongodb

En este ejemplo encontraras el CRUD de usuarios, utilizando como lenguaje golang y base de datos mongoDb

## Estructura del proyecto

```sh
golang-crud-users-mongodb/
├── src                            # Contiene el codigo de la aplicacion
      ├── database                 # Contiene la configuracion de la base de datos
      ├── models                   # Contiene los Modelos 
      ├── repository               # Contiene los metodos encargados de comunicarse con la db
      ├── service                  # Contiene la logica del negocio y se comunica con los archivos del repository                  
```



## Configurar conexiona la base de datos MongoDB
En el archivo con el nombre de mongodb.conn.go se encuentra las variables para configurar la conexion a la base de datos mongoDb

```go
var (
	usr      = "usuarioDB"
	pwd      = "ContraseñaDB"
	host     = "localhost:27017"
	database = "nombreDB"
)
```

## iniciar servidor 

Para iniciar nuestro servidor de golang se debe ejecutar el siguiente comando

```bash
go run main.go
```

## url expuestas

Para consumir nuestras apis se debe realizar por medio de las siguientes urls

* GET http://localhost:3000/users
* GET http://localhost:3000/users/:id
* POST http://localhost:3000/users
```Json
{
    "name": "Nombre",
    "email": "correo@micorreo.com"
}
```

* PUT http://localhost:3000/users/:id
```Json
{
    "name": "Nombre nuevo",
    "email": "nuevoCorreo@micorreo.com"
}
```
* DELETE http://localhost:3000/users/:id



