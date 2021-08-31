=============================
PSQL CLI
=============================
Container with psql CLI
To accept environment variables for DB URL, userid, password
In Kubernetes these environments should be supplied as secrets - and provide to pod as environment variables

=============================
REFERENCES
=============================
https://gist.github.com/petrkohut/2ed5dc83ca4afd508e9b
https://hub.docker.com/r/jbergknoff/postgresql-client
- $ docker run -it --rm jbergknoff/postgresql-client postgresql://user:pass@host:5432/db
- https://github.com/jbergknoff/Dockerfile/blob/master/postgresql-client/Dockerfile

=============================
STATUS
=============================

NOT YET WORKING. 
With entrypoint of psql it gives error:
psql: could not connect to server: No such file or directory
	Is the server running locally and accepting
	connections on Unix domain socket "/tmp/.s.PGSQL.5432"?

NEED TO ADD ENVIRONMENT VARIABLES


=============================
PSQL COMMAND FORMAT
=============================
psql -U username -h postgres_host_ip -p port
psql -U postgres -h 192.168.99.103 -p 31654
