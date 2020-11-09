# api-repository-golang-v1

For run the application
------------------------

go run main.go

Run with docker
----------------
for this example I am mount a specific windows directory to test my go app.
docker run -p 8080:8080 -v C:\imagenes:/imagenes lacaja/api-repo

for run the app on test or prod
docker run -p 8080:8080  lacaja/api-repo

Swagger UI
-----------
http://localhost:8080/swagger/index.html#/

 
