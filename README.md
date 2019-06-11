# demogo

Initial repository for Go based project

features:

1. Folder structure based on https://github.com/golang-standards/project-layout :heavy_check_mark:
2. Docker
    1. Optimized Dockerfile with multistage build :heavy_check_mark:
    2. Docker development mode with auto restart :heavy_check_mark:
3. Dependency management with Go modules :heavy_check_mark:
4. HTTP handler with [Gin](https://github.com/gin-gonic/gin)
    1. Built in middleware logger :heavy_check_mark:
    2. [CORS](https://github.com/gin-contrib/cors) security :heavy_multiplication_x:
    3. Customizable middleware function :heavy_multiplication_x:
5. Config management with [viper](https://github.com/spf13/viper) and toml format :heavy_check_mark:
6. JWT generator & verifier :heavy_multiplication_x:
