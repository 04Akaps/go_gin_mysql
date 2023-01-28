# API

go + gin + mysql 


# module

1. github.com/spf13/viper
2. github.com/gin-gonic/gin

3. github.com/swaggo/swag/cmd/swag
4. go install github.com/swaggo/swag/cmd/swag@latest
5. github.com/swaggo/files
6. github.com/swaggo/gin-swagger

# migrate
migrate create -ext sql -dir ./db -seq create_table 


# 실행법

1. envSample.env파일을 app.env로 수정 후 시작

# Swagger

http://localhost:8080/swagger/index.html