# Overview
This is an implementation of the swagger PetStore API spec for Jumbo Interactive.

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

# TODO
* how to test individual services
* how to run individual services
* how to run everything together   

checkpoint: 7h 10m spent so far