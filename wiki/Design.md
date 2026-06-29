## Architecture

My main goal is to keep the codebase easy to navigate, test, and scale. 
To do so responsibilities are intentionally separated :
 - **Router** is responsible for endpoint registration and configuration.
 - **Handlers** are responsible for HTTP concerns such as parsing requests and formatting responses.
 - **FizzBuzz logic** is responsible for generating the FizzBuzz output.
 - **Shared packages** contain reusable components such as configuration, helpers function and error responses.

### Project Structure

```
cmd/
└── api
    ├── config
    ├── docs
    ├── resource
    │   ├── common
    │   └── fizzbuzz
    └── router
```
`config`

Contains the application configuration.

Responsibilities:
 - Loading environment variables.
 - Providing default values.
 - Keeping all runtime configuration centralized.

`router`

Construct the HTTP server by initializing the shared dependencies once 
at startup before injecting them into handlers.

`ressource/fizzbuzz`

The actual REST API Implementations
It handles :
 - Parsing & Validate request parameters
 - Executing the FizzBuzz algorithm
 - Formating and returning JSON responses
 - Recording request statistics

`ressource/common`

Contains reusable components shared across resources.
Including :
 - Error response helpers
 - Test helpers
 - Response writing helpers

## Design Decision

Some of my thought to use a specific package or app

### Go
 - Go position at leboncoin 
 - Great standard library
 - `net/http` has most functionality for a production-ready API
 - Fast compilation + small artifacts

### Chi
 - Lightweight
 - `net/http` interfaces
 - Does the job right

I wanted to stay close to the Go standard library with a lightweight framework that is fully compatible with `net/http`.
Meaning that you can use chi easily without deep framework knowledge.

### go-playground/validator
 - Validators that supports struct level validation
 - Customizable error handling
 - Extensive validation rules

### OpenApi
 - Self-documenting endpoints.
 - Great improvement to maintainability
 - Interactive documentation through Swagger UI.

### In-Memory Statistics
The statistics endpoint stores request counts in memory.
This implementation was selected because:
 - The exercise does not require persistence.
 - It keeps the application simple.
 - It avoids introducing unnecessary infrastructure. 
Having an sql database just to prevent one map to be lost after a restart seemed overkill.

I made sure to allow concurrent requests to update map statistics safely by using mutex to give access one thread at a time.

On a more serious application deployed on multiple instances, 
it would likely be stored within a datastore like redis or postgreSQL to avoid losing the value to restart 
and allow multiple node to gather data for the stats.

### Configuration using environment variables
 - Environment-driven → application and configuration is separated
 - Easy to use within different environment, local-dev, CI, prod

### Rate limiting
I set some arbitrary values to protect it against extreme abuse. 
I didn't try to maximum the throughput since it's not a real end user application,
simply providing a fair access meanwhile safe guarding against obvious malicious intent.





