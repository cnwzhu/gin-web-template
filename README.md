# gin-web-template
gin sqlx logrus log split project template

### build

- linux
  ```shell
  env GOOS=linux GOARCH=amd64 go build ./cmd/app
   ```
- windows
    ```shell
   env GOOS=windows GOARCH=amd64 go build ./cmd/app
   ```

### use

- linux
  ```shell
  app -help
   ```
- windows
    ```shell
  app.exe -help
   ```

### swagger文档安装与生成

- 安装
  ```shell
  go get -u github.com/swaggo/swag/cmd/swag
  ```

- 生成
  ```shell
  swag init -g ./cmd/app/main.go -d ./
  ```