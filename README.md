# Go-Swagger

Example project demonstrate go-swagger and swagger ui

### How do I get set up? ###
 * Download swagger-ui/dist from https://github.com/swagger-api/swagger-ui and place into your project
 * Get Dependencies (if you are not using go.mod)
    ``` 
   go get -u -v github.com/go-swagger/go-swagger/cmd/swagger
   go get github.com/rakyll/statik  
   ```
 * Replace url inside index.html of swagger-ui/dist with "./swagger.yml"
 * Run script to generate swagger.yml and statik.go
    ```
    . ./swagger.sh
    ```
 * Run project using 
    ```
    make run
    ```