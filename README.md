# Go - JWT Authentication w/ Cookies
This project provides a simple JWT (JSON Web Token) authentication system using Golang (gin/gorm) and PostgreSQL for the database. It includes functionalities for user signup, login, and JWT validation. It securely stores user credentials and uses JWT for maintaining user sessions. The JWT is stored in a cookie on the client side for persistent sessions. A great starting point for anyone looking to implement authentication in their Go web projects.

## Table of Contents

- [Go - JWT Authentication w/ Cookies](#go---jwt-authentication-w-cookies)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Project Architecture](#project-architecture)
  - [Contributing](#contributing)

## Getting Started

This section will guide you through setting up the project on your local machine for development and testing purposes.

## Prerequisites

You need to have Golang and PostgreSQL installed on your machine. Additionally, you should have the gin and gorm packages for Golang.

## Installation

To get a local copy up and running, follow these steps:

1. Clone the repository:
`bash git clone https://github.com/GingFreecss2/Go-JWT-auth-cookies`

2. Navigate into the cloned repository:
`bash cd your-repository-name`

3. Install the dependencies:
`bash go mod download`

4. Start the server:
`bash go run main.go`


## Usage

The project includes three main functionalities:

- `Signup`: This allows new users to create an account.
- `Login`: This allows users to log in to their account. Upon successful login, a JWT is returned to the user in a cookie.
- `Validate`: This validates the JWT in incoming requests.

## Project Architecture

![Project Architecture](/img/go_jwt_auth_cookies.png)

This image represents the overall architecture of the project. It shows the flow of data between the various components of the project including the client, controllers, JWT package, User model, and PostgreSQL database.

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.
