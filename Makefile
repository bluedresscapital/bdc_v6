run_backend:
	go run backend/server.go

run_frontend:
	cd frontend && yarn start

db_portforward:
	kubectl --context kind-kind port-forward service/cockroachdb-public 26257

# NOTE(ma): Tbh idk yet what the exact difference cockroach's sql is vs postgresql
crdb_sql:
	kubectl exec -it cockroachdb-client-secure \
		-- ./cockroach sql \
		--certs-dir=/cockroach/cockroach-certs \
		--host=cockroachdb-public

psql:
	psql -U bdc bdc -p 26257 -h 127.0.0.1
