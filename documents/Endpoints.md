# List of all the API endpoints and there respective requirements



## Endpoint `/api/me`
This endpoint will return a json response which contains information about the current owner of the cookie they maintain.

It returns all the details of the cookie owners account like username, password (hashed), salt, account level & there email

### Example response

Success schema:
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Generated schema",
  "type": "object",
  "properties": {
    "id": {
      "type": "number"
    },
    "username": {
      "type": "string"
    },
    "password": {
      "type": "string"
    },
    "email": {
      "type": "string"
    },
    "salt": {
      "type": "string"
    },
    "account_level": {
      "type": "number"
    },
    "parent": {
      "type": "number"
    }
  },
  "required": [
    "id",
    "username",
    "password",
    "email",
    "salt",
    "account_level",
    "parent"
  ]
}
```

Error schema:
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Generated schema",
  "type": "object",
  "properties": {
    "error": {
      "type": "string"
    }
  },
  "required": [
    "error"
  ]
}
```

## Endpoint `/api/signout`
This endpoint is a sample endpoint which is used to destory brower cookies & there server sided component **token**

It only returns one json object when it's done successfully. otherwise it returns two json objects

### Example response

Success Schema:
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Generated schema",
  "type": "object",
  "properties": {
    "status": {
      "type": "boolean"
    }
  },
  "required": [
    "status"
  ]
}
```

Error schema:
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Generated schema",
  "type": "object",
  "properties": {
    "error": {
      "type": "string"
    }
  },
  "required": [
    "error"
  ]
}
```

## Endpoint `/api/users/getusers`
This endpoint will use the cookie provided to find the owners details, once we have the owners details we will then try find all of there children inside the database.

This requires no additional parameters and only involes 2 database queries to perform this request

### Example responses
**There is only an error response due to it being dynamicly populated**

Error schema:
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Generated schema",
  "type": "object",
  "properties": {
    "error": {
      "type": "string"
    }
  },
  "required": [
    "error"
  ]
}
```

### Endpoint `/api/users/createuser`
this endpoint will allow the creation of users via an API endpoint, allows us to create users via the api.

## Example responses
The response for this card is dynamiclly produced meaning it changes, we only have a static error response

Error schema:
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Generated schema for Root",
  "type": "object",
  "properties": {
    "error": {
      "type": "string"
    },
    "status": {
      "type": "boolean"
    }
  },
  "required": [
    "error",
    "status"
  ]
}
```

## Endpoint `/api/pages/dashboard`
This endpoint returns information which we intend for usage on the dashboard for filling widgets & graphs in

## Example responses

Error schema: 
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Generated schema for Root",
  "type": "object",
  "properties": {
    "error": {
      "type": "string"
    },
    "status": {
      "type": "boolean"
    }
  },
  "required": [
    "error",
    "status"
  ]
}
```