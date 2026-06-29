# Feature details

## Validation rules
The validation of each fizzbuzz query parameters with the following rules :
 - `first_multiple` and `second_multiple` must be `> 0`
 - `limit_integer` must be `> 0` and `<= 50000`
 - All query strings are required

## Error handling
Each error from the validator is translated into a JSON list of errors that is understable by the end user.

Example:
```json
{
  "errors": [
    "FirstMultiple must be greater than 0"
  ]
}
```

## Rate Limiting
Currently, 60 request per minute per IP is allowed. 
The choice was arbitrary simply to avoid someone extreme abuse case.
This value is going to be adjusted and improved if exposed to an actual user base. 
Giving it more depth, like a per minute amount but also a burst amount...

## Testing
We have unit tests for the fizzbuzzapi, and a simple benchmark
You can run them with the following commands
```bash
go test -v ./... # Run all tests
go test -bench="." ./cmd/api/resource/fizzbuzz/ -run=^$ # Run benchmark only
```

## CI/CD
We currently have 3 active workflow
 - code_quality : runs go linter and go units tests
 - swagger_check : check if any the swagger documentation needs to be updated
 - docker_build_and_push : on every main commit if the code_quality succeeds, it will build the docker image and publish
it to two tags `latest_dev` and `main_{full_sha}_{pipeline_id}`


