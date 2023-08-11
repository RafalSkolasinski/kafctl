# Basic Examples Using Strimzi Kafka

## Requirements

K8s cluster with LoadBalancer.
I assume Kind + MetalLB.
If you use something else you may need to adjust IPs in the example configs.

One can use this Ansible [playbook](https://github.com/SeldonIO/seldon-core/blob/v2/ansible/playbooks/kind-cluster.yaml) for quick setup.


## Install Kafka

```bash
make install-strimzi create-cluster
```

## Get required certificates and password

```bash
make get-broker-ca get-client-ca get-sasl-password
```


## Trying different configuration

See example configurations under `kafctl/examples/configs` folder.
Choose one you want to try using `-c` flag, e.g.

```bash
$ kafctl get messages -c configs/plaintext.properties my-topic
```
