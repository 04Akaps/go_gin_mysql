# API

go + gin + mysql 

로그 파일 생성 및 swagger를 구성하며, makefile을 통한 docker에서 mysql실행이 담겨 있습니다.

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