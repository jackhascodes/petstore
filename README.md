# Overview
This is an implementation of the swagger PetStore API spec for Jumbo Interactive.
The final service which has been delivered is the Pet service `./pet/`. Other services have been stubbed as individual 
applications within this repo to demonstrate the split as explained in the "Design" section below.

# Time Tracking and Thought Process
I have made regular commits and made notes in the commit messages on when I started and stopped.
As of this writing I have spent ~8h 10m. The approach I took to coding this can be followed through the commit log.
It should be noted that in an actual development situation, I would squash these commits and write up a more appropriate 
commit message. Please forgive the messy nature of some of the commit messages herein.

# Running
In the pet service, there is a `docker` directory containing docker-compose files for testing and running the Pet service.

`./pet/docker/docker-compose-test.yml` can be used to test the internal pet package. It includes setup for a dummy database
with migrated data.

```
cd ./pet/docker
docker-compose -f docker-compose-test.yml build
docker-compose -f docker-compose-test.yml up
```
__NOTE:__ There *may* be some cases where the database fails to boot fully before the tests begin. There are ways around
this, but please see the "what I would improve" section below.

To run the actual Pet service:
```
cd ./pet/docker
docker-compose -f docker-compose.yml build
docker-compose -f docker-compose.yml up
```
This docker compose file makes the database and application available via exposed ports.
Please refer to the `swagger.yaml` for API details.
A simple test can be done using cURL, though it is recommended you use something like Postman to validate the API.
Note also that Authentication has been stubbed and hard-coded in the docker-compose file as an env var.
Any requests made wihtout a header `x-api-key: testkey` will respond `401 Unauthorized`.
```
$ ~/go/src/github.com/jackhascodes/petstore $ curl -v http://localhost:8081/pet/1 -H "x-api-key: testkey"
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8081 (#0)
> GET /pet/1 HTTP/1.1
> Host: localhost:8081
> User-Agent: curl/7.47.0
> Accept: */*
> x-api-key: testkey
> 
< HTTP/1.1 200 OK
< Date: Thu, 17 Oct 2019 04:02:43 GMT
< Content-Length: 239
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host localhost left intact
{"id":1,"category":{"id":1,"name":"dogs"},"name":"Nacho","photoUrls":["https://s3.public/img1-1.png","https://s3.public/img1-2.png","https://s3.public/img1-3.png"],"tags":[{"id":1,"name":"friendly"},{"id":3,"name":"easy-care"}],"status":1}
```


## Design
The Swagger PetStore API lends itself to straightforward service boundaries.
* __Pet service__: this will take care of the CRUD operations via direct http requests and event subscriptions. It will 
be backed by its own database.
* __Store service__: this will act as an orchestrator via direct http requests and an event service to handle orders and
publish order events. Essentially a gateway service, it will not be backed by a database.
* __Order service__: this will track orders and have mock-functionality to simulate orders being processed. It will 
communicate order status to the Pet service via the event service and receive order requests from the Store service. It 
will be backed by its own database.
* __User service__: this will be stubbed for time allowances.
* __Event service__: this will be an out-of-the-box message-queue service which will communicate events from the Store 
and Order services. 

## Considerations
Consideration was given to making a separate inventory service, but this was discarded as a solution due to the limited 
nature of use cases outlined in the API spec. Given that it would essentially be a simple aggregator of pet statuses 
it was decided to fold that functionality into the Pet service.

The use of a message-queue to handle events, and in fact to use events at all, is intended to allow for long-lived 
order simulation and also to demonstrate the extensibility of this design. For example, integration with other services
(e.g. SalesForce, etc) can be done with minimal (if any) interference to existing services. NATs was chosen because of 
the simple integration with Go and a very easy setup for demonstration purposes.

## Final Implementation
With 8 hours, I managed to build out the Pet service. I did not incorporate the inventory endpoint, though the service 
level functionality is there. The other services (Order, Store, User) are stubbed in this repository to show how I would
have tackled them given the time. I also did not implement a message-queue handler in the Pet service, and so did not 
include NATs in the docker composition. 

## What Could Be Improved
So much.

The database dependency would be better served with a bash script (or similar) which is able to verify that the database
is accepting connections. This would enable integration tests to begin only when the database is ready.
  
The code is not as DRY as I would like. There are many places where it would be better to go back and refactor duplicate
functionality.

