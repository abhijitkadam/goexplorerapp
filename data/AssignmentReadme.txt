Aggregator API server Assignment:
Section A contains prerequisite and setup to start on assignment

A] Data services setup:
Steps to setup data services to be consumed by the Aggregator API server that you will build.
1. install nodejs

2. install json-server to create fake services
npm install -g json-server

3. Run service silver which serves silver data from file silverdata.json.json on port 8085
    json-server --watch silverdata.json --port 8085

    You can access the data from browser: http://localhost:8085/inventory
    or from curl: curl http://localhost:8085/inventory

4. Run service gold which serves gold data from file golddata.json on port 8086. Run on different terminal window so that you have 2 services running.
    json-server --watch golddata.json --port 8086

    You can access the data from browser: http://localhost:8086/inventory
    or from curl: curl http://localhost:8086/inventory


B] Aggregator API server details:

Build an REST API server in Java or any other tech stack that you are comfertable with.

The API server should have an HTTP GET endpoint "/hotels":
    http://<server>:<port>/hotels

    - It should give list of all the hotels to the user who is querying above endpoint.
    - The list should come from silverdata & gold data. The Aggregator API service should 
    consume (call) data services mentioned in section-A above to get the combined data.
    - You need to use http client to fetch the data from http://localhost:8085/inventory 
    & http://localhost:8086/inventory
    - The server name and port of the data services server can be read as configuration information. 
        Good not to hardcode with localhost.
    - No need to do any caching in the server.
    - Use validation for wrong inputs and do error handling.
    - Unit testing is good to have not mandatory. 

    queries to support:
    1) http://<server>:<port>/hotels   
        => above should give all hotels combined
    
    2) http://<server>:<port>/hotels?orderby=price 
        => above should give all hotels in ascending order of price

    3) http://<server>:<port>/hotels?source=silver
        => above should give all hotels from silver source. If source=gold then only gold source data. This
            is to filter room by data source.

    4) http://<server>:<port>/hotels?roomtype=LARGE
        => above should give all hotels with room type LARGE. Also can give other options like 
            roomtype=X-LARGE or roomtype=SMALL. This is to filter rooms by type.

    5) Queries and filters should work in combination like:
        http://<server>:<port>/hotels?roomtype=LARGE&roomtype=X-LARGE&orderby=price