
ROADMAP:
A user api
A PGN viewer,
To run a local couchdb: run
docker run -p 5984:5984  -v /home/nevroz/elschach/data:/opt/couchdb/data 
-e COUCHDB_USER=admin -e COUCHDB_PASSWORD='' -d -d couchdb 