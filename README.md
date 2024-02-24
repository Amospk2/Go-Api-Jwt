# User Management API

This User Management API provides a RESTful service for managing users, including operations such as creating, retrieving, updating, and deleting user records. It also features a secure authentication mechanism for user login and access control.

## Features

- Create new users
- Retrieve all users
- Retrieve a single user by ID
- Update user details
- Delete users
- User authentication and token generation

## Prerequisites

Before you begin, ensure you have met the following requirements:
- Go 1.15+ installed
- PostgreSQL 12.0+ installed
- An environment for running Go applications

## Installation 

1. Clone the repository to your local machine:

```bash
git clone https://github.com/your-username/Go-Api-Jwt.git
cd Go-Api-Jwt
```

2. Install the project dependencies:

```bash
go mod tidy
```

## Running the Application

To start the server, run the following command from the root of the project:

```bash
go run main.go
```

The API will be available at [http://localhost:9999](http://localhost:9999).

## Usage

### Authenticating a User

1. To authenticate, send a POST request to `/login` with a JSON body containing the user's email and password:

```bash
curl -X POST http://localhost:9999/login \
-H "Content-Type: application/json" \
-d '{"email":"user@example.com", "password":"password123"}'
```

2. On successful authentication, you'll receive a JWT token in the response. Use this token to authenticate further requests to protected endpoints.

### Working with Users

- **Create a User**

```bash
curl -X POST http://localhost:9999/users \
-H "Authorization: Bearer YOUR_JWT_TOKEN" \
-H "Content-Type: application/json" \
-d '{"name":"John Doe", "email":"john@example.com", "age":"30"}'
```

- **Get All Users**

```bash
curl http://localhost:9999/users \
-H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.


### Notes:
- Replace `https://github.com/your-username/user-management-api.git` with your actual repository URL.
- Customize the installation and usage instructions based on the actual setup and capabilities of your API.
- If your project has additional dependencies or requires more complex setup steps, be sure to include those details.
- Adjust the example `curl` commands according to your API's endpoints and expected request formats.