# demogo

Initial repository for Go based project

features:

1. Folder structure based on https://github.com/golang-standards/project-layout [x]
2. Docker
  1. Optimized Dockerfile with multistage build [x]
  2. Docker development mode with auto restart [x]
3. Dependency management with Go modules [x]
4. HTTP handler with [Gin](https://github.com/gin-gonic/gin)
  1. Built in middleware logger [ ]
  2. [CORS](https://github.com/gin-contrib/cors) security [ ]
  3. Customizable middleware function [ ]
5. Config management with [viper](https://github.com/spf13/viper) and toml format [x]
6. JWT generator & verifier [ ]
