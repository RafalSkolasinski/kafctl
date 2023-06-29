STRIMZI_NAMESPACE ?= kafka
CLUSTER_NAMESPACE ?= ${STRIMZI_NAMESPACE}


install-strimzi:
	kubectl create ns ${STRIMZI_NAMESPACE} || echo "Namespace kafka already exists"
	helm repo add strimzi https://strimzi.io/charts/
	helm install kafka-operator strimzi/strimzi-kafka-operator -n ${STRIMZI_NAMESPACE}

create-cluster:
	kubectl create ns ${CLUSTER_NAMESPACE} || echo "Namespace kafka already exists"
	kubectl apply -f kafka.yaml -n ${CLUSTER_NAMESPACE}


get-broker-ca:
	kubectl get secret -n ${CLUSTER_NAMESPACE} default-cluster-ca-cert -o json | jq -r '.data."ca.crt"' | base64 -d > configs/broker.ca.crt

get-client-ca:
	kubectl get secret -n ${CLUSTER_NAMESPACE} mtls -o json | jq -r '.data."ca.crt"' | base64 -d > configs/client.ca.crt
	kubectl get secret -n ${CLUSTER_NAMESPACE} mtls -o json | jq -r '.data."user.crt"' | base64 -d > configs/client.user.crt
	kubectl get secret -n ${CLUSTER_NAMESPACE} mtls -o json | jq -r '.data."user.key"' | base64 -d > configs/client.user.key

.ONESHELL: get-sasl-password
get-sasl-password:
	PASSWORD=$$(kubectl get secret -n ${CLUSTER_NAMESPACE} sasl -o json | jq .data.password -r | base64 -d)
	sed "s/sasl.password=.*/sasl.password=$${PASSWORD}/" configs/sasl.properties.template > configs/sasl.properties


deploy:
	kubectl apply -f kctl.yaml

remove:
	kubectl delete --grace-period=1 -f kctl.yaml || echo "already removed"
