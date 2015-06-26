services
--------

Package `services` contains the defintions for all the Elos API services. Services are another `aeolus` term, but are basically structures and logic that are not directly related to what the API logic and structures do, but provide information needed for the elos API to complete a request. For example, a database service, provides an access point to the Elos database. The database could be directly on the API struct, but that's wrong, because it isn't directly a member of the API.

For those of you who are microservice crazy, these interfaces in services could be intermediaries to RPC for other hosts with their own service logic.

