operator: create_crd apply_operator

create_crd:
	kubectl apply -f crdb.cockroachlabs.com_crdbclusters.yaml

apply_operator:
	kubectl apply -f operator.yaml

crdb:
	kubectl apply -f example.yaml

portforward:
	kubectl --context kind-kind port-forward service/cockroachdb-public 26257

shell:
	kubectl exec -it cockroachdb-client-secure \
		-- ./cockroach sql \
		--certs-dir=/cockroach/cockroach-certs \
		--host=cockroachdb-public

psql:
	psql -U bdc bdc -p 26257 -h 127.0.0.1
