# Indexer

## Introduction
This program is designed to search and process emails from a specified directory using concurrent workers. It utilizes Go's concurrency features to efficiently handle large datasets.

## Usage
The program accepts several command-line arguments to customize its behavior:

- `-dir`: Specifies the directory to search for email data. Defaults to `enron_data`.
- `-workers`: Number of concurrent workers to process emails. Defaults to 5.
- `-batch-size`: Number of emails to process per batch. Defaults to 10000.
- `-profile`: Optional flag to enable CPU and memory profiling.

To run the program, execute the following command:

```bash
./indexer [-dir directory] [-workers num_workers] [-batch-size batch_size] [-profile]
```