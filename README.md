# kctl

Simple CLI to test Kafka configurations.

## Installation

```bash
go install github.com/RafalSkolasinski/kctl@latest
```

or run `make install` command from root of the repository.

## Basic usage

See [examples](./examples/README.md) for different configuration examples.

### Creating Topics

```bash
$ kctl create topics -c configs/plaintext.properties my-new-topic
CONFIGURATION bootstrap.servers: 172.18.255.92:9092
CONFIGURATION security.protocol: PLAINTEXT
Created: my-new-topic (Topic 'my-new-topic' already exists.)
```

### Listing Topics

```bash
$ kctl get topics -c configs/plaintext.properties my-new-topic
CONFIGURATION bootstrap.servers: 172.18.255.92:9092
CONFIGURATION security.protocol: PLAINTEXT
Topic (partitions: 2): my-topic
Topic (partitions: 2): my-new-topic
Topic (partitions: 50): __consumer_offsets
```

### Listening for Messages

```bash
kctl get messages -c configs/plaintext.properties my-topic
CONFIGURATION bootstrap.servers: 172.18.255.92:9092
CONFIGURATION security.protocol: PLAINTEXT
Message on my-topic[1]@31: {"key":"2ed0c7aa","message":"msg-1"}
Message on my-topic[1]@32: {"key":"af5a97d8","message":"msg-2"}
Message on my-topic[1]@33: {"key":"36e452c1","message":"msg-3"}
Message on my-topic[1]@34: {"key":"a059502b","message":"msg-5"}
Message on my-topic[1]@35: {"key":"d563ddc0","message":"msg-6"}
Message on my-topic[1]@36: {"key":"45b20975","message":"msg-7"}
Message on my-topic[1]@37: {"key":"8ab8fdb6","message":"msg-9"}
Message on my-topic[0]@39: {"key":"8a14f7f4","message":"msg-0"}
Message on my-topic[0]@40: {"key":"8cca8c42","message":"msg-4"}
Message on my-topic[0]@41: {"key":"f89a2e88","message":"msg-8"}
```

### Sending for Messages

```bash
$ kctl create messages -c configs/plaintext.properties my-topic
CONFIGURATION bootstrap.servers: 172.18.255.92:9092
CONFIGURATION security.protocol: PLAINTEXT
Delivered message to my-topic[1]@15
Delivered message to my-topic[1]@16
Delivered message to my-topic[1]@17
Delivered message to my-topic[1]@18
Delivered message to my-topic[1]@19
Delivered message to my-topic[0]@25
Delivered message to my-topic[0]@26
Delivered message to my-topic[0]@27
Delivered message to my-topic[0]@28
Delivered message to my-topic[0]@29
```
