# Transactions History Backend
## Endpoints available
- GET `/users`: Returns all the users whit their names and total balance
- GET `/transactions/user/{id}`: Returns the transactions of the user with the given id
- GET `/summary/user/{id}`: Returns the summary of all transactions of the user with the given id including the transactions per month
- GET `/summary/email/{id}/to/{email}`: Sends the summary of all transactions of the user with the given id to the given email

## Description
This is a project to show the transactions history of a user. It is a simple project that uses the following technologies:
- Go 1.19
- Gin Gonic
- Gorilla Mux
- Gomail
- Docker
- Docker Compose
- AWS S3
- AWS Lambda
- AWS Simple Email Service

## Architecture
The project uses a simple architecture with the following layers:
- **Controller**: This layer is in charge of receiving the requests and sending the responses. It is the entry point of the application.
- **Service/Handlers**: This layer is in charge of the business logic. It is the core of the application.
- **Repository**: This layer is in charge of the database operations. It is the data access layer of the application (in this case is a CSV file hosted in a AWS S3 Bucket).

## How to run in local environment
1. Clone the repository [transaction-history-backend](https://github.com/dobleub/transaction-history-backend)
2. Inside `transaction-history-backend/deployments/local-docker` folder, add a `.env` file with the following content:
```bash
AWS_LAMBDA_RUNTIME_API=
AWS_ACCESSKEYID=
AWS_SECRETACCESSKEY=
AWS_DEFAULTREGION=
AWS_BUCKET=
EMAIL_SMTP_HOST=
EMAIL_SMTP_PORT=
EMAIL_SMTP_SECURE=
EMAIL_SMTP_USERNAME=
EMAIL_SMTP_PASSWORD=
EMAIL_IAM_USERNAME=
```
3. Once `.env` is completed, run the following command in the same folder as before in order to to start the server:
```bash
docker-compose up -d --build
```
- Container name es `transactions-history-backend` and exposes the following ports:
	- 0.0.0.0:3003->3030/tcp
4. The server will be running in your host on port 3003. You can test it with the following command:
```bash
curl http://localhost:3003/users
```
5. To stop the server, run the following command:
```bash
docker stop transactions-history-backend
```
6. To remove the container, run the following command:
```bash
docker rm transactions-history-backend
```

## How to deploy in production environment
1. Follow previous steps to clone the repository and build the docker image
2. Make some changes
3. Push the changes to the repository at `develop` branch
4. Create a pull request to `master` branch
5. Once the pull request is approved, the changes will be deployed automatically to production environment

## How to run tests
1. Follow previous steps to clone the repository and build the docker image
2. Run the following command in the same folder as before in order to to start the server:
```bash
docker exec -it transactions-history-backend go test ./...
```
