pavan:exchange-rate-service/ (main) $ hey -n 1000000 -c 5000 http://localhost:8080/convert\?base_currency\=USD\&target_currency\=INR\&amount\=100

Summary:
  Total:        32.0886 secs
  Slowest:      2.3439 secs
  Fastest:      0.0002 secs
  Average:      0.1546 secs
  Requests/sec: 31163.7555

  Total data:   39000000 bytes
  Size/request: 39 bytes

Response time histogram:
  0.000 [1]     |
  0.235 [777211]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.469 [177677]        |■■■■■■■■■
  0.703 [36106] |■■
  0.938 [7415]  |
  1.172 [1241]  |
  1.406 [293]   |
  1.641 [41]    |
  1.875 [10]    |
  2.110 [3]     |
  2.344 [2]     |


Latency distribution:
  10% in 0.0026 secs
  25% in 0.0460 secs
  50% in 0.1154 secs
  75% in 0.2207 secs
  90% in 0.3524 secs
  95% in 0.4531 secs
  99% in 0.6884 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0007 secs, 0.0002 secs, 2.3439 secs
  DNS-lookup:   0.0008 secs, 0.0000 secs, 0.3918 secs
  req write:    0.0002 secs, 0.0000 secs, 0.2414 secs
  resp wait:    0.1520 secs, 0.0001 secs, 2.3438 secs
  resp read:    0.0008 secs, 0.0000 secs, 0.1399 secs

Status code distribution:
  [200] 1000000 responses
