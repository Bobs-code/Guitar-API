# Guitar API

## Description 
The Guitar API is a RESTful API that provides information about various guitar makes and modes, along with information about famous mucisians who have used those models at some point in their careers. This API allows developers to access ruitar-related information to build applications, websites, or any other projects that require guitar-related data.

## Requirements
- Go Programming Language
- PostgreSQL
   - For storing and retrieving API database information
- pgAdmin
- API Testing Tool
   - [Postman](https://www.postman.com/) or something similar

## Getting Started
In it's current version, the API requires the above tools to be installed locally on one's machine. Continue below for installation docs to install each of the tools on your operating system.

### Setup
- [Go](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
   - [Getting Started Tutorial](https://www.postgresql.org/docs/current/tutorial.html)
- [pgAdmin](https://www.pgadmin.org/download/)
- Download repo into local file directory 
 - `git clone https://github.com/Bobs-code/Guitar-API.git`

#### Import database file into Postgres
Using pgAdmin: 
  1. Navigate to and select "Databases" > "Create" > "Database..."
  2. Name database and select "Save"
  3. Right Click Database and select "Restore..."
  4. Navigate to the `.sql` file located in the db folder and select the file.
  5. Click "Restore"
  6. API Data should now be present in the database you have created. To view this data go to "Schemas" > "Tables". Right click any of the tables and then select "View/Edit Data" > "All rows". 

## Official Documentation
[Guitar API Documentation](https://app.swaggerhub.com/apis-docs/BOBGRIFF29/guitar-api/1.0.0)

## Feature Roadmap
 - [ ] Advanced Routing using Chi Router 
 - [ ] Middleware 
 - [ ] Authentication
