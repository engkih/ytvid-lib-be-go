# YT Library API
This is API for my personal project website (YT Library) and made using Go. The reason I create this is to learn about Go, authentication, middleware, Gin, and implement what I have been learn into real case project.

## Database
This API using MySQL as database utulizing GORM. The MySql database connected throug url stored in `./database/database.go` specifily in : 
```
database, err := gorm.Open(mysql.Open("username:password@tcp(localhost:3306)/ytlibdb_go?charset=utf8mb4&parseTime=True&loc=Local"))
```
You can change the `username` and the `password` to match your MySql registered user.
The database named `ytlibdb_go` and contain two table :
- `user`
- `video`

## API Lists
### GET /api/user
Return all videos from the video table.
- URL Params: None
- Request Body: None
- Response Body:
  - Success:
    - Code: 200
    - Body:
        ```
        [
          {
            "VideoUrl": string,
            "VideoTitle": string
          }
        ]
        ```

### GET /api/video/:vidId
Return specific video detail based on `Id`
- URL Params: `:vidId`
- Request Body: None
- Response Body:
  - Succes:
    - Code: 200
    - Body:
        ```
        [
          {
            "VideoUrl": string,
            "VideoTitle": string
          }
        ]
        ```
