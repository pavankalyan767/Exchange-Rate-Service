<img width="997" height="797" alt="image" src="https://github.com/user-attachments/assets/f3614c43-751d-4242-bccc-308986c0ffd5" />


<img width="1549" height="948" alt="image" src="https://github.com/user-attachments/assets/1719963c-e0a4-4838-b99c-6c964df926b4" />


# 📈 Performance Test Report: Exchange Rate Service

This report summarizes the results of a performance test on the exchange rate service's `/history` endpoint. The test was conducted using the `hey` command-line tool, simulating a high-traffic scenario to evaluate the service's stability and responsiveness.

---

## 🛠️ Test Configuration

| Parameter            | Value                                                                                               |
| -------------------- | --------------------------------------------------------------------------------------------------- |
| **Endpoint**         | `http://localhost:8080/history?base_currency=USD&target_currency=INR&from=2025-07-14&to=2025-08-14` |
| **Total Requests**   | 1,000,000 (`-n 1000000`)                                                                            |
| **Concurrent Users** | 5,000 (`-c 5000`)                                                                                   |

---

## 📊 Summary of Results

| Metric                     | Value                   |
| -------------------------- | ----------------------- |
| **Total Test Time**        | 137.93 seconds          |
| **Requests per Second**    | 7,249.88                |
| **Average Response Time**  | 0.6797 seconds          |
| **Total Data Transferred** | 1.09 GB                 |
| **Successful Requests**    | 1,000,000 (HTTP 200 OK) |

All 1,000,000 requests were successfully completed with a **HTTP 200 OK** status code, indicating that the service handled the entire load **without any failures or errors**.

---

## ⏱️ Latency Distribution

| Percentile | Response Time (seconds) |
| ---------- | ----------------------- |
| 50%        | 0.6579                  |
| 90%        | 1.0058                  |
| 99%        | 1.3962                  |
| Fastest    | 0.0004                  |
| Slowest    | 3.3196                  |




Got it! I'll add the **convert** endpoint with curl examples and integrate it smoothly into the existing markdown.

Here’s the full README markdown with the **convert** endpoint included:


# Exchange Rate Service API

This document provides examples of how to use the Exchange Rate Service API via curl commands. The service supports fetching current exchange rates, converting amounts, and fetching historical data (for fiat currencies only).

---

## API Endpoints

- **Fetch current exchange rate:** `/fetch`
- **Convert amount between currencies:** `/convert`
- **Fetch historical exchange rates:** `/history` (fiat currencies only)

---

## Fetch Exchange Rate Examples

### 1. Fetch Exchange Rate Between Fiat Currencies (USD to INR)

Fetch the current exchange rate from USD to INR.

```bash
curl "http://localhost:8080/fetch?base_currency=USD&target_currency=INR"
````

---

### 2. Fetch Exchange Rate Between Fiat Currencies (EUR to GBP)

```bash
curl "http://localhost:8080/fetch?base_currency=EUR&target_currency=GBP"
```

---

### 3. Fetch Exchange Rate Between Cryptocurrencies (BTC to ETH)

Fetch the exchange rate from Bitcoin (BTC) to Ethereum (ETH).

```bash
curl "http://localhost:8080/fetch?base_currency=BTC&target_currency=ETH"
```

---

### 4. Fetch Exchange Rate From Fiat to Crypto (USD to BTC)

```bash
curl "http://localhost:8080/fetch?base_currency=USD&target_currency=BTC"
```

---

### 5. Fetch Exchange Rate From Crypto to Fiat (ETH to USD)

```bash
curl "http://localhost:8080/fetch?base_currency=ETH&target_currency=USD"
```

---

### 6. Fetch Exchange Rate for a Specific Date (USD to INR on 2025-08-01)

```bash
curl "http://localhost:8080/fetch?base_currency=USD&target_currency=INR&date=2025-08-01"
```

---

## Convert Amount Examples

### 7. Convert 100 USD to INR

```bash
curl "http://localhost:8080/convert?base_currency=USD&target_currency=INR&amount=100"
```

---

### 8. Convert 50 EUR to GBP

```bash
curl "http://localhost:8080/convert?base_currency=EUR&target_currency=GBP&amount=50"
```

---

### 9. Convert 0.5 BTC to ETH

```bash
curl "http://localhost:8080/convert?base_currency=BTC&target_currency=ETH&amount=0.5"
```

---

### 10. Convert 100 USD to BTC

```bash
curl "http://localhost:8080/convert?base_currency=USD&target_currency=BTC&amount=100"
```

---

### 11. Convert 1.25 ETH to USD

```bash
curl "http://localhost:8080/convert?base_currency=ETH&target_currency=USD&amount=1.25"
```

---

## Fetch Historical Exchange Rate Examples (Fiat Only)

### 12. Historical Rates from USD to INR (2025-07-14 to 2025-08-14)

```bash
curl "http://localhost:8080/history?base_currency=USD&target_currency=INR&from=2025-07-14&to=2025-08-14"
```

---

### 13. Historical Rates from EUR to GBP (July 2025)

```bash
curl "http://localhost:8080/history?base_currency=EUR&target_currency=GBP&from=2025-07-01&to=2025-07-31"
```

---

### 14. Crypto History Not Supported

History endpoint does **not** support cryptocurrency pairs. Example request:

```bash
curl "http://localhost:8080/history?base_currency=BTC&target_currency=ETH&from=2025-07-01&to=2025-07-31"
```

This request will likely return an error or no data.

---

## Notes

* Base currency and target currency codes must be valid ISO currency codes for fiat currencies or supported crypto symbols.
* Amount must be a positive number.
* Date format: `YYYY-MM-DD`.
* Historical data is only available for fiat currency pairs.

---

Feel free to modify the commands to suit your testing or integration needs.

```







