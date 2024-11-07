# Monify

## Overview
**Monify** is a backend application built with Go and Beego, designed to manage customer data, loans, team structures, and roles within an organization. It features **JWT-based authentication**  for secure access, supports concurrent processing with **Goroutines**, and includes **Docker** setup for easy database management with **PostgreSQL**. One of Monify's key highlights is its dynamic hierarchy management system, allowing organizations to scale and adjust their role structures seamlessly.

Monify leverages **Beego ORM** for model definitions but uses **direct database connections** for querying. This approach enhances flexibility and performance in database operations while maintaining strong security measures, including protection against SQL injection.

## Features
- **JWT Authentication**: Secure API endpoints with JSON Web Tokens, ensuring authorized access only.
- **Dynamic Hierarchy Management**: Flexible role-based structure that allows inserting or deleting levels without disrupting the organizational flow.
- **Goroutines**: Enhanced performance with concurrent processing, optimizing data retrieval and updates.
- **Dockerized PostgreSQL**: Easy setup with Docker to run PostgreSQL in an isolated environment.
- **Customer Management**: CRUD operations for customer profiles, including contact and financial details.
- **Loan Tracking**: Manage loan data, overdue amounts, and installment schedules.
- **Team and Role Management**: Set up teams, define roles, and organize users within a hierarchical structure.
- **API Documentation**: Auto-generated Swagger documentation for easy API exploration and testing.
- **Direct Database Querying**: Combines Beego ORM for model setup with direct DB queries, ensuring optimized performance and secure querying practices.

## Technologies Used
- **Go**: Backend programming language.
- **Beego**: Web framework for building Go applications.
- **JWT**: For secure user authentication.
- **PostgreSQL**: Database management system.
- **Docker**: For containerizing PostgreSQL for easy setup.
- **Redis**: (Optional) For session management and caching.
- **MinIO**: (Optional) Object storage, compatible with Amazon S3.

## Installation

### Prerequisites
- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [Docker](https://docs.docker.com/get-docker/) (for PostgreSQL setup)
- [PostgreSQL](https://www.postgresql.org/download/) (optional, if not using Docker)
- [MinIO](https://min.io/download) (optional, for file storage)
- [Redis](https://redis.io/download) (optional, for session management)

### Setup

1. **Clone the Repository**
   ```bash
   git clone https://github.com/your-username/monify.git
   cd monify
