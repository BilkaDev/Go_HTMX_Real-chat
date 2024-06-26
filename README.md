
# GO_HTMX Real-Time application

This Real-Time Chat Application is a web-based messaging platform that enables users to communicate instantly via a streamlined interface. Built using HTMX, this application provides a reactive experience while leveraging server-side technologies for rendering updates without full page refreshes.


## Features
- User Authentication: Secure login, signup, and logout functionalities.
- Real-Time Messaging: Instantly send and receive messages without needing to reload the page.

## Tech Stack

- **HTMX**: Handles dynamic HTML content updates.
- **Go (Golang)**: Back-end server handling HTTP requests.
- **Docker**: Used to containerize the application, ensuring consistent environments and deployment ease.

## Installation

Install my-project with Makefile

```bash
  go env -w CGO_ENABLED=1
  make run
```

Or

```bash
  docker build -t real-time-chat-app .
  docker run -p 3000:3000 real-time-chat-app
```
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`STATE`=dev
`PORT`=3000
`JWT_SECRET`=secret_key


## Endpoint

### Register a new user
```bash
curl --location --request POST 'http://localhost:3000/api/v1/auth/signup' \
--form 'fullName="Andrzej"' \
--form 'userName="Hiema"' \
--form 'email="example@example.com"' \
--form 'password="123456"' \
--form 'gender="male"' \
--form 'repassword="123456"'
```

### Authenticate a User
```bash
curl --location --request POST 'http://localhost:3000/api/v1/auth/login' \
--header 'Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyTmFtZSI6IkhpZW1hIiwiVXNlcklkIjoyMjMsImV4cCI6MTcxNDI5NDYyOX0._jKR3yfoyw1t-yi1Dng9i_Z5HjoaaAv8uPxlTad3lxg' \
--form 'fullName="Andrzej"' \
--form 'userName="Hiema"' \
--form 'email="test@wp.pl"' \
--form 'password="123456"' \
--form 'gender="male"'
```

###  Logout user
```bash
curl --location --request POST 'http://localhost:3000/api/v1/auth/logout' \
--form 'fullName="Andrzej"' \
--form 'userName="Hiema"' \
--form 'email="em"' \
--form 'password="123456"' \
--form 'gender="male"'
```

### Get users for left side bar
```bash
curl --location --request GET 'http://localhost:3000/api/v1/user'
```

### Show a message
```bash
curl --location --request GET 'http://localhost:3000/api/v1/message/{:id}' \
--form 'fullName="Andrzej"' \
--form 'userName="Hiema"' \
--form 'email="em"' \
--form 'password="123456"' \
--form 'gender="male"'
```

### Send a message
```bash
curl --location --request POST 'http://localhost:3000/api/v1/message/send/{:id}' \
--form 'message="Hi"'
```
### Conenct with web socket
```bash
'http://localhost:3000/api/v1/ws/message'
```
