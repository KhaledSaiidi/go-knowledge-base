# Go Serverless User API
Tiny Go Lambda that fronts DynamoDB for basic user CRUD, exposed through API Gateway.

## Fast Setup
- Table: `LambdaInGoUser`; PK `email` (String); attributes `email`, `firstName`, `lastName`.
- Build & ship:
  ```bash
  GOOS=linux GOARCH=amd64 go build -o main ./cmd
  zip function.zip main
  # deploy to Lambda runtime go1.x, handler=main
  ``` 
- Env: set `AWS_REGION` (or `AWS_DEFAULT_REGION`). IAM needs `dynamodb:GetItem|Scan|PutItem|DeleteItem` on the table.

## API
- `GET /?email={email}`: one user; omit `email` to list all.
- `POST /`: create `{ "email": "...", "firstName": "...", "lastName": "..." }`.
- `PUT /`: upsert same body shape.
- `DELETE /?email={email}`: delete by email.
Errors return HTTP 400 with `{"error": "..."};` success echoes the user (or list) except DELETE (empty body).

## Code Map
- `cmd/main.go`: Lambda entry; routes HTTP verbs to handlers; table name constant.
- `pkg/handlers`: API Gateway method dispatch + JSON responses.
- `pkg/user`: DynamoDB CRUD and email validation.
- `pkg/validators`: Email regex helper.
