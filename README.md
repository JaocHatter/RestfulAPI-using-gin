# RestfulAPI-using-gin
solo pràctica...
Querys locales en la terminal:
  *VER TODOS LOS ITEMS:
  $ curl http://localhost:8080/countries
  *VER UN ITEM EN ESPECÌFICO:
  $ curl http://localhost:8080/countries/<id>
  *AGREGAR UN ITEM:
    example:
  $ curl -X POST -H "Content-Type: application/json" -d '{"ID": "6", "name": "Mexico", "continent": "North America", "minsalary": 1200.00, "language": "Spanish"}' http://localhost:8080/countries
  *ELIMINAR UN ITEM:
  $ curl -X DELETE http://localhost:8080/countries/3
