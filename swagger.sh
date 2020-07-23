CURR_DIR="$(pwd)"

# SwaggerUI folder containing files for serving the frontend for swagger
SWAGGERUI=${CURR_DIR}/dist

# Statik output folder will take the swagger.yml generated together with SwaggerUI folder and convert into single statik.go
STATIK=${CURR_DIR}/output

# Remove old swagger.yml
# rm -rf ${SWAGGERUI}/swagger.yml
# echo "swagger.yml deleted at ${SWAGGERUI}"

# Generate new swagger.yml
# swagger generate spec -o ${SWAGGERUI}/swagger.yml --scan-models
# echo "swagger.yml generated at ${SWAGGERUI}"

# Validate new swagger.yml against swagger specificiation 2.0
swagger validate ${SWAGGERUI}/swagger.yml

# Generate statik.go which is then serve by the app
statik -src=${SWAGGERUI}
