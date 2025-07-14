# BudSafe

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Go Version](https://img.shields.io/badge/Go-1.24.3-blue.svg)
![Next.js](https://img.shields.io/badge/Next.js-15.3.2-black.svg)
![GraphQL](https://img.shields.io/badge/GraphQL-API-e10098)
![GCP](https://img.shields.io/badge/Google_Cloud-Deployed-4285F4)

A regulatory compliance and license tracking dashboard designed to help cannabis businesses stay compliant with evolving state laws. BudSafe provides a centralized platform for managing business licenses, monitoring renewal deadlines, and accessing state-specific regulatory updates.

## Table of Contents

- [BudSafe](#budsafe)
  - [Table of Contents](#table-of-contents)
  - [About The Project](#about-the-project)
  - [Live Demo](#live-demo)
  - [Built With](#built-with)
  - [Project Architecture](#project-architecture)
  - [API Reference](#api-reference)
  - [Deployment](#deployment)
  - [License](#license)

## About The Project

BudSafe is a full-stack web application built to tackle the complex challenge of regulatory compliance in the cannabis industry. As state laws evolve, businesses face a significant burden in tracking license renewals, compliance requirements, and regulatory changes.

This dashboard provides a clear, centralized, and easy-to-use interface to:

- Track all business and operational licenses in one place.
- Receive notifications for upcoming renewal deadlines and compliance checks.
- Access a knowledge base of state-specific regulations.
- Securely manage documents related to compliance and licensing.

The application is designed with a modern, scalable architecture to ensure reliability and allow for easy expansion as more states and regulatory frameworks are added.

## Live Demo

_[A link to the live, deployed version of the project will be available here once created.]_

## Built With

This project showcases a modern, decoupled technology stack:

**Backend:**

- [Go](https://golang.org/)
- [GraphQL](https://graphql.org/) with [gqlgen](https://gqlgen.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [sqlx](https://github.com/jmoiron/sqlx) for database interaction

**Frontend:**

- [Next.js](https://nextjs.org/) (React)
- [TypeScript](https://www.typescriptlang.org/)
- [Tailwind CSS](https://tailwindcss.com/)

**Infrastructure & Auth:**

- [Google Cloud Platform (GCP)](https://cloud.google.com/)
  - **Cloud Run:** for scalable, serverless deployment.
  - **Cloud SQL:** for managed PostgreSQL hosting.
- [Firebase Authentication](https://firebase.google.com/docs/auth) for secure user management.
- [Docker](https://www.docker.com/) for containerization.

## Project Architecture

This repository contains the source code for the entire BudSafe application, structured as a monorepo with two primary services:

1.  **`./frontend`**: A **Next.js** application that serves as the user-facing client. It handles all UI rendering and state management, communicating with the backend via GraphQL.

2.  **`./backend`**: A **Go** application that provides a robust **GraphQL API**. It is responsible for all business logic, database interactions, and serving as the single source of truth for all data.

**Authentication Flow:**
User authentication is handled by **Firebase Authentication**. The frontend client communicates directly with Firebase to sign up or log in a user. Upon success, Firebase issues a secure ID Token (JWT). This token is then included in all subsequent requests to the Go backend. The backend middleware verifies the token using the Firebase Admin SDK before processing any request, ensuring all API endpoints are secure.

## API Reference

The backend exposes a GraphQL API for all data operations. The schema, defined in the `backend/graph` directory, serves as the contract between the frontend and backend. It includes queries for fetching data and mutations for creating, updating, and deleting resources.

## Deployment

The application is designed for modern cloud infrastructure and is deployed as two separate containerized services on **GCP Cloud Run**:

1.  `budsafe-backend`: The Go GraphQL API.
2.  `budsafe-frontend`: The Next.js web application.

A `Dockerfile` is provided in both the `backend` and `frontend` directories to facilitate this process. This serverless architecture ensures scalability, reliability, and cost-efficiency.

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

