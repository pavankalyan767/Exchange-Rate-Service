<img width="997" height="797" alt="image" src="https://github.com/user-attachments/assets/f3614c43-751d-4242-bccc-308986c0ffd5" />


<img width="1549" height="948" alt="image" src="https://github.com/user-attachments/assets/1719963c-e0a4-4838-b99c-6c964df926b4" />

Sure! Here's your **Performance Test Report** converted into well-formatted **Markdown**, with a clean structure and table formatting:

---

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

The service demonstrated **exceptional performance under heavy load**. The majority of requests were served very quickly, and response time was consistently low, even with 5,000 concurrent users.

This data confirms the service's ability to maintain **high performance and reliability**, even under extreme conditions.

---

Let me know if you'd like this exported to a `.md` file or if you want a graph/visual version too.

