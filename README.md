# Authentication Service

This is a basic authentication service built using Go, allowing users to sign up, sign in, and refresh tokens. It uses JWT (JSON Web Tokens) for secure token generation and validation.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Setup Instructions](#setup-instructions)
3. [Run the Application](#run-the-application)
4. [Test the API Requests](#test-the-api-requests)
5. [API Endpoints](#api-endpoints)

## Prerequisites

Before you begin, make sure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.18 or later)
- [MongoDB](https://www.mongodb.com/try/download/community) or a running MongoDB instance
- `curl` for testing the API requests

## Setup Instructions

### 1. Clone the repository
```bash
git clone <repository-url>
cd <repository-folder>

.env 
MONGO_URI=mongodb://localhost:27017  # Replace with your MongoDB URI
JWT_SECRET=your_secret_key          # Replace with your secret key for JWT

