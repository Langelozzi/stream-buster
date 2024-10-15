# Search Endpoints

Search Endpoints are RESTful API endpoints that allows you to search for records in the database using where clauses. 
A search endpoint can be implemented for any entity in the database, if it is implemented in the backend.
The core logic for search is in the `utils/search_helpers.go` file, and is completely generic.

The endpoints are accessible via the following URLs:

```
POST /api/v1/<entity>/search
```

## Under the hood

This endpoint uses dynamic SQL building to create queries that are executed directly in the database. This enables extremely efficient searching
using WHERE clauses. When creating search objects for the request body, it can be helpful to imagine what the WHERE statement
would look like, and build the search object to reflect your desired result.

## Request Body

The request body must be a JSON object that depicts the search model in this project. The search model is as follows:

```json
{
  "Params": [
    {
      "Field": "string",
      "Operator": "string",
      "Value": "any",
      "AssociationForeignKey": "string?"
    }
  ],
  "LogicalOperator": "string"
}
```

### `Params (Required)`
An array of search parameters for the query.

#### `Field (Required)`
The field/column of the database table you are trying to search. This is case-sensitive and must match exactly
what the field name is in the database. If the field you want to search is on an associated table, then you must use dot
syntax (associated_table.field_name) and provide the `AssociationForeignKey` value (see below).

<br>

#### `Operator (Required)`
The operation to be performed in the WHERE clause. The available operators are:
- `=`: Equals. Accepts strings, numbers, booleans, and dates.
- `!=`: Not Equals. Accepts strings, numbers, booleans, and dates.
- `>`: Greater Than. Accepts numbers and dates.
- `>=`: Greater Than or Equal To. Accepts numbers and dates.
- `<`: Less Than. Accepts numbers and dates.
- `<=`: Less Than or Equal To. Accepts numbers and dates.
- `LIKE`: String searching using wildcards. Accepts strings that use wildcards. For more info on SQL wildcards: 
https://www.w3schools.com/sql/sql_wildcards.asp
- `NOT LIKE`: String searching using wildcards. Accepts strings that use wildcards.
- `IN`: Accepts an array of values. Accepts strings, numbers, booleans, and dates.
- `NOT IN`: Accepts an array of values. Accepts strings, numbers, booleans, and dates.
- `IS NULL`: Accepts no value. Used to check if a field is null.
- `IS NOT NULL`: Accepts no value. Used to check if a field is not null.

<br>

#### `Value (Optional)`
The value used for the right operand of the WHERE clause. This can be any valid json datatype, so long as it
is a valid datatype for the operator. For example, the `IN` operator accepts an array, but the `=` operator accepts single values.

When using the `IS NULL` and `IS NOT NULL` operators, the value can be omitted.

<br>

#### `AssociationForeignKey (Optional)`
Required if the `Field` references an associated table instead of a direct column.
The value of this should be the foreign key of the associated table that you are trying to search.
This is used to build the query with the correct join statement.

<br>

### `LogicalOperator (Optional)`
Indicates which logical operator to use in the query when multiple parameters are specified. 
Must have one of the following values: `AND` or `OR` (case-sensitive).

Required if there are multiple parameters in the `Params` array.


## Response
If the search is successful, the response will be a JSON array with the results of the search, and a status code of 200 will be returned. 
If no results are found, an empty array will be returned.

If the search fails, a status code of 400 will be returned, along with an error message in the response body, in the form of a JSON object.


## Searching on Associated Tables
If the field that you want to search on is on an associated table, some specific required additions must be present in the request body:
- You must specify the associated table and field/column using dot notation in the `Field` value of the search parameter.
The table name and field/column name must be exactly what is defined in the database (not the model).
  - The format should look like: `associated_table.field_name`
- The `AssociatedForeignKey` value must be provided in the search parameter. This value should be the exact string name of the
foreign key on the associated table in the database.


## Examples
### Searching for documents with a PDF file type
```
POST /api/v1/document/search
```

Request
```json
{
  "Params": [
    {
      "Field": "file_type",
      "Operator": "=",
      "Value": ".pdf"
    }
  ]
}
```

Response
```
[...]
```

<br>

### Searching for documents with specific ID's
```
POST /api/v1/document/search
```

Request
```json
{
  "Params": [
    {
      "Field": "id",
      "Operator": "IN",
      "Value": [10020, 13028, 10016, 28394]
    }
  ]
}
```

Response
```
[...]
```

<br>

### Searching for documents that have been approved (approved_at and approved_by are not null)
```
POST /api/v1/document/search
```

Request
```json
{
  "Params": [
    {
      "Field": "approved_at",
      "Operator": "IS NOT NULL"
    },
    {
      "Field": "approved_by",
      "Operator": "IS NOT NULL"
    }
  ],
  "LogicalOperator": "AND"
}
```

Response
```
[...]
```
<br>

### Searching for non-deleted documents owned by a specific user (user_id 2)
```
POST /api/v1/document/search
```

Request
```json
{
  "Params": [
    {
      "Field": "user_documents.user_id",
      "Operator": "=",
      "Value": 2,
      "AssociationForeignKey": "document_id"
    },
    {
      "Field": "user_documents.is_owner",
      "Operator": "=",
      "Value": true,
      "AssociationForeignKey": "document_id"
    },
    {
      "Field": "user_documents.deleted_at",
      "Operator": "IS NULL",
      "AssociationForeignKey": "document_id"
    }
  ],
  "LogicalOperator": "AND"
}
```

Response
```
[...]
```
<br>

### Searching for documents uploaded between a certain date range
```
POST /api/v1/document/search
```

Request
```json
{
  "Params": [
    {
      "Field": "uploaded_at",
      "Operator": ">=",
      "Value": "2024-01-01T00:00:00Z"
    },
    {
      "Field": "uploaded_at",
      "Operator": "<",
      "Value": "2024-02-01T00:00:00Z"
    }
  ],
  "LogicalOperator": "AND"
}

```

Response
```
[...]
```

<br>

### Searching for documents whose file names contain the word 'report'
```
POST /api/v1/document/search
```

Request
```json
{
  "Params": [
    {
      "Field": "file_name",
      "Operator": "LIKE",
      "Value": "%report%"
    }
  ]
}


```

Response
```
[...]
```
