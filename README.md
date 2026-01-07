# Authentication and Authorization with Go

This project is a backend API developed in Go.
Its main goal is to implement a basic authentication and authorization system using JWT.

The project is designed as an independent service, focused on practicing good backend practices such as layer separation, error handling, basic security, and environment variable management.

## Project Goal

- Implement user registration and login

- Protect routes using JWT

- Handle basic roles (user / admin)

- Simulate a reusable authentication service for other APIs

- This project is part of my backend portfolio.

## Technologies Used

- Language: Go (Golang)
- Web Framework: Gin Gonic
- Database: MariaDB / MySQL
- Persistence: database/sql
- Environment Variables: godotenv
- Security: JWT + bcrypt

## Planned Features

- User registration
- Login with JWT generation
- Authentication middleware
- Protected private routes
- Role-based access control
- Logout (token invalidation)

## Planned Endpoints

```
http
POST   /auth/register
POST   /auth/login
POST   /auth/refresh
POST   /auth/logout
GET    /me

```

## Environment Variables

The project uses a .env file to manage sensitive information such as:
Database credentials
JWT secret keys
A .env.example file is provided as a reference.

## Architecture
The project follows a simplified hexagonal architecture (Hexagonal Lite)

Business logic is separated from infrastructure concerns to keep the codebase clean and maintainable