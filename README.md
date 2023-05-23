# Guitar API

## Table of Contents
- [Description](#description) :page_with_curl:
- [API Tools](#api-tools) :hammer_and_wrench:
- [Getting Started](#getting-started)
   - [Installation](#installation)
- [Recommended Use](#recommended-use)
- [Still to Come](#still-to-come)


## Description 
The Guitar API is a RESTful API that provides information about various guitar makes and modes, along with information about famous mucisians who have used those models at some point in their careers. This API allows developers to access ruitar-related information to build applications, websites, or any other projects that require guitar-related data.

## API Tools & Use
- Go Programming Language Standard Library
- PostgreSQL
   - For storing and retrieving API information
- `lib/pq`[^1] Postgres driver for `database/sql`
   - For interfacing with the postgres database
- Postman
   - Accessing Endpoints


## Getting Started
In it's current version, the API requires the above tools to be installed locally on one's machine. Continue below for installation docs to isntall each of the tools on your operating system.

### Installation
#### - [Go](https://go.dev/doc/install)
#### - [PostgreSQL](https://www.postgresql.org/download/)
   - [Getting Started Tutorial](https://www.postgresql.org/docs/current/tutorial.html)





## Still To Come
 - [ ] Web Based Authentication
 
[^1]: See [lib/pq docs](https://pkg.go.dev/github.com/lib/pq) for more information
