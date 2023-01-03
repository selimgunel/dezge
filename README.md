
ROADMAP:

To run a local couchdb: run
docker run -p 5984:5984  -v /home/nevroz/dezge/data:/opt/couchdb/data 
/home/nevroz/dezge/etc:/opt/couchdb/etc/local.d
-e COUCHDB_USER=admin -e COUCHDB_PASSWORD='SET A PASSWORD' -d  narslan/dezge-couchdb 

docker build -t narslan/dezge-couchdb .

