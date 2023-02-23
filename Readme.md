## Test Instructions

## Prerequisites:

Docker - You can install docker here:  https://www.docker.com/products/docker-desktop/


### Steps:

1. Build the image 
```
    docker build -t covid-app:latest .
```

2. Run the app using docker-compose
```
    docker-compose up -d
```


### App

You can now access the REST endpoint with this URL `http://localhost:3001/api/v1/covidapp`

### Endpoints

GET http://localhost:3001/api/v1/covidapp/top/confirmed - returns the top countries with confirmed covid 19 cases


### Postman Collection

Added the postman collection that can be used to perform the request under `/postman_collection`


### How I did my testing
- You can watch the recordings here https://drive.google.com/file/d/1yVFTEEFQrssrbBDt4GdfkYImSqB9zc7R/view?usp=sharing
