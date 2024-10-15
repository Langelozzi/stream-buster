# Code Standards

## Variable naming
- Variables should be named in camelCase
- Constants should be named in UPPERCASE
- If the variable is receiving a value from a return function it should contain the return type in the variable declaration

## Function naming
- Functions should be named in camelCase if they are private to the file
- Functions should be named in PascalCase if they are public to the rest of the module

## File naming
- Files should be all lowercase and snake_case if they contain multiple words
- Files in the `controllers` should be named `*_controller.go`
- Files in the `daos` should be named `*_dao.go`
- Files in the `daos/interfaces` should be named `*_dao_interface.go`
- Files in the `middlewares` should be named `*_middleware.go`
- Files in the `models` should be named `*.go`
- Files in the `routes/api/*` should be named `*_routes.go`
- Files in the `services` should be named `*_service.go`
- Files in the `services/interfaces` should be named `*_service_interface.go`
- Files in the `utils/dependency_injection` should be named `*_di.go`

## Creating structs (classes)
- Structs should be named in PascalCase
- They should be in their own file, and the file should have the following structure:
    - package declaration
    - import statements
    - struct declaration
        - fields should be named in camelCase if they are private to the struct
        - fields should be named in PascalCase if they are public to the rest of the module (this is the case for models)
    - "constructor" function with naming convention `New*` that returns the struct instance pointer
    - methods for the struct

## CRUD Methods
- Insert the name of the Table where the `*` is

### Controllers
- GetAll*Handler
- Get*Handler
- Create*Handler
- Create*BulkHandler
- Update*Handler
- Delete*Handler
- Delete*BulkHandler
### Services
- GetAll*
- Get*
- Create*
- Create*Bulk
- Update*
- Delete*
- Delete*Bulk
### Daos
- GetAll*Dao
- Get*Dao
- Create*Dao
- Create*BulkDao
- Update*Dao
- Delete*Dao
- Delete*BulkDao

## Dependency Injection
- Dependency injection should be used to inject services and daos into the controllers
- Whenever a service or dao is referenced, it should be passed as a parameter or stored in a variable of its interface type
    - NOTE: The interface type cannot be a pointer, but the underlying struct that is assigned to that variable can be a pointer

## Comments
- We want to have readable code that is self-explanatory, so comments should be used sparingly
    - This means writing good variable names, function names, struct names and organizing the code in a way that is easy to understand
- Comments should be used to explain why something is done in a certain way, or to explain a complex piece of code
- The API endpoints will be documented using Swagger annotations, so no need to write comments for the API endpoints themselves
