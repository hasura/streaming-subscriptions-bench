# streaming-subscriptions-bench

This is a very simple way to validate the runtime characterstics of streaming subscriptions on the database.

## Benchmarking methodology

1. There are 2 parts to benchmarking: 1) a streaming subscriber, 2) concurrent mutations
2. We start a streaming subscription and start concurrent mutations (configurable concurrency, default: 1000)
3. We check pg_stat_statements to check the max_time of execution of the SQL query.We also monitor the benchmarking app and the database for load (cpu/mem)

## Todo

1. We should test with concurrent subscriptions as well.