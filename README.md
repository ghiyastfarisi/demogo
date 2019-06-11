# demogo

Initial repository for Go based project

features:

1. Folder structure based on https://github.com/golang-standards/project-layout [done]
2. Docker
  - Optimized Dockerfile with multistage build [done]
  - Docker development mode with auto restart [done]
3. Dependency management with Go modules [done]
4. HTTP handler with [Gin](https://github.com/gin-gonic/gin)
  - Built in middleware logger [done]
  - [CORS](https://github.com/gin-contrib/cors) security
  - Customizable middleware function
5. Config management with [viper](https://github.com/spf13/viper) and toml format [done]
6. JWT generator & verifier
