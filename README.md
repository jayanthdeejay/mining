# DeBruijn Sequence Generator for Cryptocurrency Addresses

This GoLang project aims to generate DeBruijn Sequences of length 256 over the alphabet 0 and 1. The primary objective is to generate a DeBruijn Sequence that covers all possible 256-bit binary strings. For each generated string, the application computes Bitcoin and Ethereum addresses from the corresponding 256-bit private key. It then checks these generated addresses against a database to determine if they already exist. If an address exists in the database, the private key is stored for future reference.

## Algorithm Used

The project utilizes the Granddaddy algorithm to find the next bit in the DeBruijn Sequence. The source of inspiration for the Granddaddy algorithm comes from the work of Joe Sawada et. al. You can contact Joe Sawada at [email](jsawada@uoguelph.ca) or [site](http://debruijnsequence.org/db/home) for further information.

## Sequence Generation and Key Processing

The DeBruijn Sequence is generated incrementally. After a new bit is generated, the first bit is dropped, and the new bit is added to the end of the sequence. This updated key is then pushed to either a RabbitMQ channel or a Redis list, depending on your configuration.

The application employs a producer-consumer framework. The producer is responsible for generating 256-bit strings and pushing them to a queue in either RabbitMQ or Redis. Distributed workers (consumers) then consume these strings and process them into cryptocurrency addresses. The supported address formats include P2PKH, P2SH, P2WPKH for Bitcoin and Ethereum addresses.

## Rate Limiting

To ensure manageable processing load, the producer is designed to generate a maximum of 100,000 keys per minute. This rate limit allows the consumer to catch up with the processing and ensures system stability.

## Project Purpose

The author of this project is perfectly aware of the fact that generating a DeBruijn Sequence of 256-bit length over the alphabet 0 and 1 will take at least a million years, and that this is a CPU-bound algorithm. The primary motivation behind working on this project was to learn GoLang and gain experience with RabbitMQ. This project serves as a learning exercise and may not be practical for real-world use cases due to the extreme computational requirements.

## Future Work

Future work on this project will involve parallelizing the key processing by leveraging GPU cores. This optimization will enhance the efficiency of generating and processing cryptocurrency addresses, potentially increasing the throughput and scalability of the system.

Feel free to contribute to this project or use it as a reference for your own cryptocurrency address generation needs.

**Note**: Make sure to check the project's documentation and configuration files for specific setup instructions and usage guidelines.
