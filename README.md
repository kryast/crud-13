CRUD ke-13

POST
curl -X POST http://localhost:8080/category \
-H "Content-Type: application/json" \
-d '{"name": "Electronics"}'

GET
curl -X GET http://localhost:8080/category
curl -X GET http://localhost:8080/category/1

PUT
curl -X PUT http://localhost:8080/category/1 \
-H "Content-Type: application/json" \
-d '{"name": "Updated Electronics"}'