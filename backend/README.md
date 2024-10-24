# Docswap-backend

## Table of contents
- [Description](#description)
  - [Built with](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Usage](#usage)
- [Environment Variables](#environment-variables)
  - [DB_CONNECTION_STRING](#db_connection_string)
  - [CLIENT_URL](#client_url)
  - [BLOB_CONTAINER](#blob_container)
  - [AZURE_JWKS_URL](#azure_jwks_url)
  - [AZURE_B2C_URL](#azure_b2c_url)
  - [AZURE_CLIENT_ID](#azure_client_id)
  - [AZURE_ACCOUNT_NAME](#azure_account_name)
  - [AZURE_ACCOUNT_KEY](#azure_account_key)
- [Project Structure](#project-structure)
- [Additional Documentation](#additional-documentation)

## Description
The Docwap API is a RESTful API that enables users to upload, share, and purchase real estate documents using a token-based system, facilitating information exchange primarily for realtors.

### Built with
* [![Go][GolangIcon]](https://go.dev/)
* [![Gin][GinIcon]](https://gin-gonic.com/docs/introduction/)
* [![GORM][GormIcon]](https://gorm.io/)
* [![Swagger][SwaggerIcon]](https://swagger.io/)
* [![Azure][AzureIcon]](https://azure.microsoft.com/en-us/)
* [![MS SQL Server][SqlIcon]](https://www.microsoft.com/en-us/sql-server/sql-server-2022)

## Getting Started
The following instructions will guide you through setting up the project on your local machine.

### Prerequisites
1. Go(lang) [version 1.16 or higher, preferably v1.9.*]
   - Download and install Go from [here](https://go.dev/doc/install)
   - Verify the installation by running the following command in a terminal:
    ```shell
    go version
    ```
2. A SQL Server
   - Any SQL server will work, but we recommend using a dockerized version of [MS SQL Server](https://hub.docker.com/_/microsoft-mssql-server/)
   - A database must be created in the server and accessible via a database connection string

3. Install Make command [optional but recommended]
   - If you are using Windows, you can download it from [here](https://gnuwin32.sourceforge.net/packages/make.htm)
   - If you are using MacOS or Linux, you can download it using your terminal package manager (NOTE: Many
Linux distros have the `make` command preinstalled)
   - Verify that Make is installed by running the following command in a terminal:
    ```shell
    make -v
    ```

### Installation
1. Clone the repository by running the following command in a terminal:
   ```shell
   git clone https://github.com/DOC-SWAP/Docswap-backend.git
   ```
   
2. Cd into the project directory
   ```shell
   cd Docswap-backend/
   ```

3. Set up environment variables
   - All variables outlined in the [Environment Variables](#environment-variables) section require valid values
   - After setup, you should have a `.env` file in the root of the project with the following variables:

    ```
    DB_CONNECTION_STRING=
    CLIENT_URL=
    BLOB_CONTAINER=
    AZURE_JWKS_URL=
    AZURE_CLIENT_ID=
    AZURE_B2C_URL=
    AZURE_ACCOUNT_NAME=
    AZURE_ACCOUNT_KEY=
    ```

### Usage
1. When running the system for the first time, you must navigate to the `main.go` file in the root of the project
and uncomment the `database.InitializeDb()` function. This will ensure that the database tables are properly created
in your database before the application starts up. The main function should look like so:
    ```go
    func main() {
        // Uncomment the following line to run the db initialization for updates 
        database.InitializeDb()

        ...
    }
    ```

2. If you have the `make` command installed in your system, then you can run the following command in the root of the project
to install necessary dependencies, perform initialization and start the server:
   ```shell
   make run
   ```
   
   Alternatively, you can run all the commands individually in your terminal:

   ```shell
   go install github.com/swaggo/swag/cmd/swag@latest // Install swag command
   ```

   ```shell
   go get -u github.com/swaggo/swag/cmd/swag         // Import swag command
   ```
   
   ```shell
   swag init                                         // Initialize swagger documentation
   ```
   
   ```shell
   go mod tidy                                       // Install project dependencies
   ```
   
   ```shell
   go run main.go                                    // Start the application
   ```
   
3. The API should be running on `localhost:8080`
4. The API endpoint documentation is available to view at:
    ```
    localhost:8080/api/v1/swagger/index.html
    ```


## Environment variables
### `DB_CONNECTION_STRING`
Defines the parameters necessary for establishing a connection to the database. This will depend on the type of database
and where your database is located.
The `connection_string` can often be found via a DBMS or cloud portal, but this differs depending on the database itself.

```
DB_CONNECTION_STRING="{connection_string}"
```
<br>

### `CLIENT_URL`
The URL of the client application that is using this API. This variable is used for CORS purposed to secure the connection
to the API. If you want to allow all CORS connections, set the variable to `CLIENT_URL="*"`.

```
CLIENT_URL="{client_url}"
```
<br>

### `BLOB_CONTAINER`
The name of the container in your Azure Storage resource that will be used to store the uploaded documents. 
The `container_name` can be found in your Azure portal and is configurable through Azure.

```
BLOB_CONTAINER="{container_name}"
```
<br>

### `AZURE_JWKS_URL`
Indicates the url provided by Azure for fetching the JSON Web Keys, used to validate
the token signature. For more information, see documentation [here](https://learn.microsoft.com/en-us/entra/identity-platform/access-tokens#validate-tokens).
The `tenant` may also be called "Directory name" in the Azure portal, and the `user_flow_name`
is available under "User flows" in the B2C resource portal.

```
AZURE_JWKS_URL="https://{tenant}.b2clogin.com/{tenant}.onmicrosoft.com/{user_flow_name}/discovery/v2.0/keys"
```
<br>

### `AZURE_B2C_URL`
Contains the Azure B2C tenant url, used for validating the issuer of the token.
The `tenant` may also be called "Directory name" in the Azure portal, and the `tenant_directory_id` can be found
under "Directory + subscriptions" which is accessible when switching directories in the portal.

```
AZURE_B2C_URL="https://{tenant}.b2clogin.com/{tenant_directory_id}/v2.0/"
```

<br>

### `AZURE_CLIENT_ID`
Contains the value of the client ID for the front end application. It is used to
validate that the token was issued for the proper application.
The `docswap_front_end_azure_client_id` can be found under the Azure AD B2C portal -> Manage -> App registrations.

```
AZURE_CLIENT_ID={docswap_front_end_azure_client_id}
```

<br>

### `AZURE_ACCOUNT_NAME`
The name of the Azure Storage resource in the Azure portal.

```
AZURE_ACCOUNT_NAME={storage_resource_name}
```

<br>

### `AZURE_ACCOUNT_KEY`
The Azure Storage resource key that enables SAS (Shared Access Signature) keys to be generated, enabling users
to access the Azure Storage files. The `storage_key` can be found in the Azure portal storage resource under
`Security + Networking > Access keys`

```
AZURE_ACCOUNT_KEY={storage_key}
```


## Project Structure
```
├── .github/
    └── workflows/
├── controllers/
├── daos/
    └── interfaces/
├── middlewares/
├── models/
    ├── auth/
    └── search/
├── routes/
    └── api/
        └── v1/
├── services/
    └── interfaces/
├── technical_documentation/
├── tests/
    └── postman_tests/
├── utils/
    ├── auth/
    ├── database/
    └── dependency_injection/
├── .env
├── .gitignore
├── go.mod
├── main.go
├── Makefile
└── README.md
```

- `.github`: Files used for CI/CD with GitHub
  - `workflows`: Yml files that define GitHub action workflows
- `controllers`: Controllers that handle the requests and responses
- `daos`: Data access objects that interact with the database and other external services
  - `interfaces`: Interfaces that are used to define the methods that the daos should implement
- `middlewares`: Middlewares that are used to intercept requests and responses
- `models`: Models that are used to represent the data in the database
  - `auth`: Structures used for the authentication handling
  - `search`: Models used in the search functionality
- `routes`: Routes that are used to define the endpoints of the application
  - `api`: API routes
    - `v1`: Version 1 of the api routes
- `services`: Services that perform business logic and facilitate communication between controllers and daos
  - `interfaces`: Interfaces that are used to define the methods that the services should implement
- `technical_documentation`: Markdown files with detailed documentation on the project and API features
- `tests`: Testing files
  - `postman_tests`: JSON files that define postman tests for API endpoints. These can be imported into postman and executed
- `utils`: Utility functions that are used throughout the application
  - `auth`: Structures and functions that enable the authentication middleware
  - `database`: Functions that initialize and establish connection to the database
  - `dependency_injection`: Dependency injection files that are used to inject the services and daos into the controllers


## Additional Documentation
All additional technical documentation for this project is stored in the `technical_documenation/` directory.


<!-- Markdown links and images -->
[GolangIcon]: https://img.shields.io/badge/Golang-00ADD8?style=for-the-badge&logo=go&logoColor=white
[GinIcon]: https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white
[GormIcon]: https://img.shields.io/badge/GORM-00ADD8?style=for-the-badge&logo=go&logoColor=white
[AzureIcon]: https://img.shields.io/badge/Microsoft_Azure-0089D6?style=for-the-badge&logo=microsoft-azure&logoColor=white
[SqlIcon]: https://img.shields.io/badge/MS_SQL_Server-CC2927?style=for-the-badge&logo=microsoft-sql-server&logoColor=white
[SwaggerIcon]: https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=swagger&logoColor=black
