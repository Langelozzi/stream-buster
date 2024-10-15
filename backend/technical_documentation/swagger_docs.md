# Swagger Documentation

## Overview

Swagger is a set of tools for designing, building, documenting, and consuming RESTful web services. It offers a user-friendly framework to describe all elements of a REST API, ensuring clear communication and a streamlined workflow between frontend and backend teams.

## Adding Swagger Annotations to Routes

When creating a new route, ensure to add Swagger annotations to each controller method that defines the route. This helps in documenting the functionality, parameters, and the expected responses for each API endpoint.

For detailed guidance on adding these annotations, refer to the documentation [here](https://github.com/swaggo/swag?tab=readme-ov-file#general-api-info).

## Adding Annotations to Models

Similarly, when creating a new model, you should also add Swagger annotations. This includes details about the model properties, types, and any constraints or relevant descriptions that help understand the model structure.

I have found that the best way to do this is just to ask chatGpt to generate the annotations.

### Controller Method Annotation Template for Swaggo

Below is a template for annotating controller methods in Swaggo. Replace `<MethodDescription>`, `<Parameters>`, `<Responses>`, etc., with appropriate descriptions and details relevant to your specific API method.

```go
// <MethodDescription>
// @Summary Summary of the endpoint
// @Description Detailed description of the endpoint
// @Tags <Tag>
// @Accept  <InputFormats> // e.g., json
// @Produce <OutputFormats> // e.g., json
// @Param   name   path    <Type>  <IsRequired> "Description of parameter"
// @Success 200 {object} <SuccessResponseType> "Description of successful response"
// @Failure <ErrorCode> {object} <ErrorResponseType> "Description of error response"
// @Router /route/to/endpoint [method]
```

Replace `<Tag>`, `<InputFormats>`, `<OutputFormats>`, `<Type>`, `<IsRequired>`, `<SuccessResponseType>`, `<ErrorResponseType>`, `<ErrorCode>`, and `[method]` with the relevant details for your API. This template should be modified according to the specific requirements of each endpoint.
## Example
[document_controller.go](https://github.com/DOC-SWAP/Docswap-backend/blob/develop/controllers/document_controller.go)
```go
// GetDocument retrieves a document by its ID
// @Summary Retrieve document
// @Description get document by ID
// @Tags documents
// @Accept  json
// @Produce  json
// @Param id path int true "Document ID"
// @Success 200 {object} models.Document "Successfully retrieved the document"
// @Failure 400 {object} map[string]interface{} "Error: Bad Request"
// @Router /documents/{id} [get]
func (contr *DocumentController) GetDocument(c *gin.Context) {}
```
## What this actually does
if you run the server and go to `localhost:8080/api/v1/index.html` you will see the swagger documentation for the API.
