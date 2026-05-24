# rest

## actions

Source: [http://www.vinaysahni.com/best-practices-for-a-pragmatic-restful-api](http://www.vinaysahni.com/best-practices-for-a-pragmatic-restful-api)

Verbs: GET, POST, PUT, PATCH, DELETE

Verb | Action | Description 
:------------- |:-------------| :-----
GET | /tickets | retrieves a list of tickets
GET | /tickets/12 | retrieves a specific ticket
POST | /tickets | creates a new ticket
PUT | /tickets/12 | updates ticket 12
PATCH | /tickets/12 | partially updates ticket 12
DELETE | /tickets/12 | delete ticket 12

Relations

Verb | Action | Description 
:------------- |:-------------| :-----
GET | /tickets/12/messages | retrieves a list of messages for ticket #12
GET | /tickets/12/messages/5 | retrIeves message #5 for ticket #12
POST | /tickets/12/messages | creates a new message in ticket #12
PUT | /tickets/12/messages/5 | updates message #5 for ticket #12
PATCH | /tickets/12/messages/5 | partially update message #5 for ticket #12
DELETE | /tickets/12/messages/5 | deletes message #5 for ticket #12

Actions that doesnt fit on crud

1. Reestructure the action to appear like a resource. For example an activate action could be mapped to a boolean activated and updated via a PATCH
2. Treat it like a sub-resource. PUT /tickets/{id}/activated
3. Multi-resource search doesnt apply to any resource endpoint. Create a /search endpoint.

## ssl everywhere

## have good documentation

An API is only as good as its documentation.

## versioning

https://api.myapp.com/v1/  
https://api.myapp.com/v2/

## result filtering, sorting, searching

filtering: 

`GET /tickets?state=open  `

sorting: 

`GET /tickets?sort=priority`  
`GET /tickets?sort=-priority,created_at`

searching: 

`GET /tickets?q=searchString&state=open&sort=-priority,created_at`

## alias 

Make experience more pleasant for consumer

Recently closed tickets: 

`GET /tickets/recently_closed`

## limit fields

`GET /tickets?fields=id,subject,custome_name,&state=open&sort=-updated_at`

## updates & creation should return representation

PUT, POST, PATCH 
Prevent another hit to the API. The API should return the updated or created representation of the resource as part of the response.

## try to avoid envelope

```json
{
	"data" : {
		"id" : 1,
		"name" : "jim"
	}
}
```

## authentication

RESTFUL should be stateless. Should not depend on cookies or sessions. 

## caching

ETag, Last-Modified

## error handling

Useful error message:

```json
{
	"code" : 1234,
	"message" : "something bad happened",
	"description" : "more details about the error"
}
```

## pagination

Example

`https://api.github.com/user/repos?page=3&per_page=100`

An API that requires sending a count can use a custom HTTP header like X-Total-Count.

## embed

`GET /tickets?embed=customer&fields=id,customer.id,customer.name`

## http status codes

Status Code | Description 
:------------- | :-----
200 | OK - Response to a successful GET, PUT, PATCH or DELETE. Can also be used for a POST that doesn't result in a creation.
201 | Created - Response to a POST that results in a creation. Should be combined with a Location header pointing to the location of the new resource
202 | Accepted - The request has been accepted for processing, but the processing has not been completed
204 | No Content - Response to a successful request that won't be returning a body (like a DELETE request)
301 | Moved Permanently - This and all future requests should be directed to the given URI
304 | Not Modified - Used when HTTP caching headers are in play
308 | Permanent Redirect - The request and all future requests should be repeated using another URI
400 | Bad Request - The request is malformed, such as if the body does not parse
401 | Unauthorized - When no or invalid authentication details are provided. Also useful to trigger an auth popup if the API is used from a browser
403 | Forbidden - When authentication succeeded but authenticated user doesn't have access to the resource
404 | Not Found - When a non-existent resource is requested
405 | Method Not Allowed - When an HTTP method is being requested that isn't allowed for the authenticated user
410 | Gone - Indicates that the resource at this end point is no longer available. Useful as a blanket response for old API versions
415 | Unsupported Media Type - If incorrect content type was provided as part of the request
422 | Unprocessable Entity - Used for validation errors
429 | Too Many Requests - When a request is rejected due to rate limiting
500 | Internal Server Error - A generic error message, given when an unexpected condition was encountered and no more specific message is suitable
501 | Not Implemented - The server either does not recognize the request method, or it lacks the ability to fulfil the request
503 | Service Unavailable - The server cannot handle the request (because it is overloaded or down for maintenance)


| Code | Name                         | When to Use (REST APIs) |
|------|------------------------------|--------------------------|
| 200  | OK                           | Successful GET, PUT, PATCH (with response body) |
| 201  | Created                      | Resource successfully created (POST) |
| 204  | No Content                   | Successful request with no response body (DELETE, some PUT/PATCH) |
| 400  | Bad Request                  | Invalid request format, malformed JSON, validation error (generic) |
| 401  | Unauthorized                 | Authentication required or invalid token |
| 403  | Forbidden                    | Authenticated but not allowed to access resource |
| 404  | Not Found                    | Resource does not exist |
| 409  | Conflict                     | Business rule conflict (e.g., duplicate email) |
| 422  | Unprocessable Entity         | Valid JSON but semantic/validation errors |
| 500  | Internal Server Error        | Unexpected server failure |
| 502  | Bad Gateway                  | Upstream service failed (microservices) |
| 503  | Service Unavailable          | Service overloaded or under maintenance |
