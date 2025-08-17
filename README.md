# Exchange Rate Service

A high-performance currency exchange rate service built with Go and Go-Kit framework, featuring real-time exchange rates, currency conversion, and comprehensive monitoring.

## 🚀 Quick Start

### Prerequisites
- Docker and Docker Compose
- Go 1.21+ (for local development)

### Running the Project

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd exchange-rate-service
   ```

2. **Setup environment configuration**
   ```bash
   cp .env.example .env
   ```
## Pleaes use the API keys , mentioned below , for convenience , since this is a private repo , i've made them available in the readme 
## or you can get new api keys at https://api.exchangerate.host/ , http://api.coinlayer.com/
3. **Enter your API keys in `.env` file**
   ```bash
   # Edit .env file and add your API keys
   
  FIAT_API_KEY=032b7d9d9fb6d0029737824b9fc3b43d
  FIAT_API_URL=https://api.exchangerate.host/
  CRYPTO_API_URL=http://api.coinlayer.com/
  CRYPTO_API_KEY=fed5a907404a4c87d6f3608dd891e589
   


4. **Start the entire stack with Docker Compose**
   ```bash
   docker compose up --build -d
   ```
   This command will:
   - Pull Grafana and Prometheus images
   - Build the exchange rate service
   - Start all services in detached mode

5. **Verify services are running**
   - **Exchange Rate Service**: http://localhost:8080
   - **Prometheus**: http://localhost:9000
   - **Grafana**: http://localhost:3000 (admin/admin)
   - data source for grafana http://prom-server:9090

## 📡 API Testing

### Test the Service Endpoints

#### Exchange Rate Fetching
```bash
# Fiat to Fiat conversion
curl "http://localhost:8080/fetch?base_currency=USD&target_currency=INR"

# Cryptocurrency conversion  
curl "http://localhost:8080/fetch?base_currency=BTC&target_currency=ETH"

# Mixed currency conversion
curl "http://localhost:8080/fetch?base_currency=USD&target_currency=BTC"

# Historical rate lookup
curl "http://localhost:8080/fetch?base_currency=USD&target_currency=INR&date=2025-08-01"
```

#### Amount Conversion
```bash
# Convert 100 USD to INR
curl "http://localhost:8080/convert?base_currency=USD&target_currency=INR&amount=100"

# Convert cryptocurrency amounts
curl "http://localhost:8080/convert?base_currency=BTC&target_currency=ETH&amount=0.5"

# Mixed currency conversion
curl "http://localhost:8080/convert?base_currency=ETH&target_currency=USD&amount=1.25"
```

#### Historical Data Retrieval
```bash
# Get historical rates for a date range
curl "http://localhost:8080/history?base_currency=USD&target_currency=INR&from=2025-07-14&to=2025-08-14"

# Monthly historical data
curl "http://localhost:8080/history?base_currency=EUR&target_currency=GBP&from=2025-07-01&to=2025-07-31"
```

## 📊 Monitoring & Observability

- **Prometheus Metrics**: http://localhost:9090/query
- **Grafana Dashboards**: http://localhost:3000 (Login: admin/admin)

### Key Metrics Available
- Request count and latency
- Cache hit rates
- External API response times
- System resource utilization

## 🧪 Testing

The project includes comprehensive testing covering:

### Cache Testing
- In-memory cache performance and reliability
- TTL management and automatic cleanup
- Thread-safe operations under concurrent load

### Conversion Testing
- Currency conversion accuracy
- Cross-currency calculations (Fiat-to-Fiat, Crypto-to-Crypto, Mixed)
- Historical rate conversions

### Running Tests
```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Run specific test suite
go test ./service/... -v
```

---

# Exchange Rate Service
## Final Project Submission

**Technologies:** Go • Go-Kit • Docker • Prometheus • Grafana • Production Ready

---

## Table of Contents

1. [Executive Summary](#1-executive-summary)
2. [Requirements Achievement](#2-requirements-achievement)
3. [System Architecture & Design](#3-system-architecture--design)
4. [Project Structure](#project-structure)
5. [Key Features Implementation](#4-key-features-implementation)
6. [Performance Analysis](#5-performance-analysis)
7. [API Documentation](#6-api-documentation)
8. [Testing & Quality Assurance](#7-testing--quality-assurance)
9. [Deployment & Operations](#8-deployment--operations)
10. [Monitoring & Observability](#9-monitoring--observability)
11. [Technical Decisions](#10-technical-decisions)

---

## 1. Executive Summary

Successfully developed and delivered a high-performance Exchange Rate Service that exceeds all project requirements. The solution is built using Go and Go-Kit framework, implementing clean architecture principles with comprehensive monitoring, testing, and deployment capabilities.

### Project Highlights

- **Performance Excellence:** Handles high concurrent load with excellent response times
- **Zero Error Rate:** 100% success rate during extensive load testing
- **Sub-Second Response:** Average response time of 0.68 seconds
- **Full Feature Coverage:** All requirements plus bonus features implemented
- **Production Ready:** Complete with monitoring, Docker deployment, and comprehensive testing

### Key Metrics

| Metric | Value | Description |
|--------|--------|-------------|
| Concurrent Users | 5K | Maximum concurrent users supported |
| Average Response Time | 0.68s | Sub-second response performance |
| Error Rate | 0% | Perfect reliability under load |
| Requests per Second | 7,249 | High throughput capacity |

---

## 2. Requirements Achievement

### Core Requirements Compliance

| Requirement | Status | Implementation Approach |
|------------|--------|------------------------|
| Ingest from Public APIs | ✅ Complete | Integrated exchangerate.host for fiat currencies and coinlayer.com for cryptocurrencies |
| RESTful API Endpoints | ✅ Complete | Three endpoints: /fetch, /convert, /history with comprehensive parameter validation |
| Hourly Rate Updates | ✅ Complete | Background goroutines with ticker-based scheduling for automated updates |
| Currency Support | ✅ Complete | All 5 required fiat currencies (USD, INR, EUR, JPY, GBP) plus 3 cryptocurrencies as specified in requirements |
| 90-Day Historical Limit | ✅ Complete | Date validation with automatic enforcement of historical data constraints |
| Clean Architecture | ✅ Complete | Go-Kit framework with proper separation of transport, service, and data layers |
| Docker Deployment | ✅ Complete | Multi-stage Dockerfile with docker-compose orchestration |

### Bonus Features Achievement

| Bonus Feature | Status | Technical Implementation |
|--------------|--------|-------------------------|
| In-Memory Caching | ✅ Implemented | Custom cache with TTL management, automatic cleanup, and thread-safe operations |
| Concurrency Handling | ✅ Implemented | RWMutex for cache safety, goroutines for background tasks, channels for coordination |
| Error Handling | ✅ Implemented | Graceful degradation, API failure recovery, comprehensive error logging |
| Unit Testing | ✅ Implemented | Comprehensive test suite covering core business logic and edge cases |
| Prometheus + Grafana | ✅ Implemented | Full observability stack with custom metrics and visual dashboards |
| Go-Kit Framework | ✅ Implemented | Complete implementation with middleware chain and clean endpoints |
| Cryptocurrency Support | ✅ Implemented | BTC, ETH, USDT support as per requirements with cross-conversion capabilities between fiat and crypto |

---

## 3. System Architecture & Design

### Design Philosophy

The system follows clean architecture principles with clear separation of concerns, making it maintainable, testable, and scalable. Built using Go-Kit framework to ensure production-grade reliability and observability.

### Key Architectural Decisions

- **Layered Architecture:** Transport → Service → Cache/Client layers with clear boundaries
- **Dependency Injection:** All components are injected, enabling easy testing and swapping
- **Interface-Based Design:** Service contracts defined through interfaces for flexibility
- **Middleware Chain:** Logging, metrics, and validation handled through middleware
- **Concurrent Design:** Thread-safe operations with proper synchronization primitives

### Component Overview

| Component | Responsibility | Key Features |
|-----------|---------------|--------------|
| HTTP Transport Layer | Request handling and routing | Go-Kit HTTP handlers, JSON encoding/decoding, parameter validation |
| Middleware Stack | Cross-cutting concerns | Request logging, Prometheus metrics, error handling, timing |
| Service Layer | Business logic execution | Rate calculations, currency conversions, unified rate fetching logic |
| Cache System | In-memory data storage | TTL management, automatic cleanup, thread-safe operations |
| API Client | External service integration | HTTP client with context, timeout handling, error recovery |
| Background Services | Automated data updates | Hourly rate fetching, cache maintenance, goroutine management |

---

## Project Structure

```
pavan:exchange-rate-service/ (main*)
.
├── bin
│   └── rate-exchange-service
├── cache
│   └── cache.go
├── client
│   └── client.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   └── constants.go
├── main.go
├── Makefile
├── prometheus.yml
├── service
│   ├── convert.go
│   ├── exchange_service_test.go
│   ├── fetch.go
│   ├── history.go
│   ├── middleware.go
│   ├── rate_fetcher.go
│   └── service.go
├── tmp
│   ├── build-errors.log
│   └── main
├── transport
│   ├── convert.go
│   ├── fetch.go
│   ├── history.go
│   ├── middleware.go
│   └── transport.go
└── types
    └── types.go

9 directories, 26 files
```

### Directory Structure Overview

| Directory | Purpose | Key Files |
|-----------|---------|-----------|
| `/bin` | Compiled binary output | rate-exchange-service executable |
| `/cache` | Caching implementation | cache.go - In-memory cache with TTL management |
| `/client` | External API client | client.go - HTTP client for external APIs |
| `/internal` | Internal configuration | constants.go - Currency definitions and configuration |
| `/service` | Business logic layer | Core service implementations and tests |
| `/transport` | HTTP transport layer | Go-Kit HTTP handlers and middleware |
| `/types` | Type definitions | Shared data structures and interfaces |

### Currency Configuration

The system currently supports only the currencies specified in the requirements document:

**Supported Fiat Currencies:**
- USD (United States Dollar)
- INR (Indian Rupee)  
- EUR (Euro)
- JPY (Japanese Yen)
- GBP (British Pound Sterling)

**Supported Cryptocurrencies:**
- BTC (Bitcoin)
- ETH (Ethereum)
- USDT (Tether)

**Adding New Currencies:**
To add support for additional currencies, simply update the currency constants in `internal/constants.go`. The system is designed to be easily extensible - new currencies can be added by including their currency codes in the appropriate constant maps without requiring changes to the core business logic.

---

## 4. Key Features Implementation

### Universal Currency Conversion
Supports all conversion types: Fiat-to-Fiat, Crypto-to-Crypto, and Mixed conversions using USD as base currency for optimal API efficiency.

### High-Performance Caching
Custom in-memory cache with separate stores for fiat and crypto rates, featuring TTL management and automatic cleanup routines.

### Real-Time Data Updates
Automated hourly updates using background goroutines with proper error handling and logging for data freshness.

### Concurrency Safety
Thread-safe operations using RWMutex, enabling thousands of concurrent requests without data races or corruption.

### Historical Data Management
90-day historical rate storage with date validation and range queries, optimized for read-heavy workloads.

### Robust Error Handling
Comprehensive error management with graceful degradation, API failure recovery, and detailed error logging.

---

## 5. Performance Analysis

### Load Testing Configuration

Conducted extensive performance testing using the 'hey' command-line tool to simulate real-world high-traffic scenarios with excellent results demonstrating the system's capability to handle production-level loads.

### Test Results Summary

| Metric | Value | Performance Impact |
|--------|-------|-------------------|
| Test Duration | 137.93 seconds | Sustained performance over extended periods |
| Requests per Second | 7,249.88 | High throughput suitable for production environments |
| Average Response Time | 0.6797 seconds | Sub-second response times for excellent user experience |
| Success Rate | 100% (0 errors) | Perfect reliability under load conditions |

### Response Time Distribution

| Percentile | Response Time | Analysis |
|------------|--------------|----------|
| 50th Percentile | 0.6579 seconds | Median response time shows consistent performance |
| 90th Percentile | 1.0058 seconds | 90% of requests completed within 1 second |
| 99th Percentile | 1.3962 seconds | Even outlier requests remain under 1.4 seconds |
| Fastest Request | 0.0004 seconds | Cache hits provide near-instantaneous responses |
| Slowest Request | 3.3196 seconds | Worst-case scenario still within acceptable bounds |

### Load Testing Results

![Load Testing Terminal Output](https://github.com/user-attachments/assets/f3614c43-751d-4242-bccc-308986c0ffd5)

The load testing demonstrates the service's ability to handle high concurrent traffic with consistent performance metrics.

### Performance Monitoring Dashboard

![Grafana Performance Dashboard](https://github.com/user-attachments/assets/1719963c-e0a4-4838-b99c-6c964df926b4)

The Grafana dashboard shows comprehensive performance metrics including request rates, response times, and system health indicators.

---

## 6. API Documentation

### Endpoint Overview

| Endpoint | Purpose | Currency Support | Key Features |
|----------|---------|------------------|--------------|
| /fetch | Get exchange rates between currencies | All combinations | Real-time rates, historical dates, cross-currency calculations |
| /convert | Convert amounts between currencies | All combinations | Amount conversion, date-specific rates, precision handling |
| /history | Historical rates for date ranges | Fiat currencies only | 90-day lookback, date validation, range queries |

### Response Examples

#### /fetch Endpoint Response
```json
{
  "base_currency": "USD",
  "target_currency": "INR",
  "rate": 83.25,
  "date": "2025-08-17"
}
```

#### /convert Endpoint Response
```json
{
  "base_currency": "USD",
  "target_currency": "INR",
  "amount": 100,
  "converted_amount": 8325.00,
  "rate": 83.25,
  "date": "2025-08-17"
}
```

#### /history Endpoint Response
```json
{
  "base_currency": "USD",
  "target_currency": "INR",
  "rates": [
    {
      "date": "2025-08-01",
      "rate": 83.15
    },
    {
      "date": "2025-08-02",
      "rate": 83.22
    }
  ]
}
```

---

## 7. Testing & Quality Assurance

### Testing Strategy

| Test Type | Coverage | Purpose |
|-----------|----------|---------|
| Unit Tests | Core business logic | Validate conversion algorithms, cache operations, error handling |
| Cache Testing | In-memory cache functionality | TTL management, thread safety, performance under load |
| Conversion Testing | Currency conversion accuracy | Cross-currency calculations, precision handling |
| Edge Case Testing | Error conditions | Invalid currencies, date ranges, network failures |

### Implemented Test Coverage

#### Cache Testing
- **TTL Management**: Verifies automatic expiration of cached data
- **Thread Safety**: Concurrent access testing with multiple goroutines
- **Cache Hit/Miss**: Performance validation for cache efficiency
- **Cleanup Operations**: Automatic removal of expired entries

#### Conversion Testing
- **Accuracy Validation**: Mathematical precision in currency conversions
- **Cross-Currency Logic**: USD-based conversion chain validation
- **Historical Conversions**: Date-specific rate applications
- **Error Scenarios**: Invalid currency pairs and amounts

### Quality Assurance Measures

- **Code Coverage:** Comprehensive test suite covering critical business logic paths
- **Error Simulation:** Tests for API failures, invalid inputs, and edge conditions
- **Concurrency Testing:** Thread safety verification under high load
- **Data Validation:** Input sanitization and format validation
- **Performance Monitoring:** Response time tracking and throughput analysis

---

## 8. Deployment & Operations

### Containerization Strategy

| Component | Technology | Configuration |
|-----------|------------|---------------|
| Application Container | Multi-stage Dockerfile | Go builder stage + Alpine runtime for minimal footprint |
| Service Orchestration | Docker Compose | Application + Prometheus + Grafana in unified network |
| Environment Management | .env configuration | API keys and URLs externalized for different environments |
| Development Tools | Air hot reload + Makefile | Streamlined development workflow with automatic rebuilds |

### Deployment Features

- **One-Command Deployment:** `docker compose up --build -d` starts entire stack
- **Environment Configuration:** Externalized API keys and endpoints via `.env` file
- **Service Discovery:** Internal Docker networking for service communication
- **Volume Management:** Persistent storage for Grafana dashboards and Prometheus data
- **Development Workflow:** Hot reloading and easy debugging setup

### Make Commands

```bash
# Build the application
make build

# Run tests
make test

# Run with hot reload
make dev

# Clean build artifacts
make clean

# Run load testing
make load-test
```

---

## 9. Monitoring & Observability

### Metrics Collection

| Metric Category | Specific Metrics | Business Value |
|-----------------|------------------|----------------|
| Request Metrics | Request count, latency distribution, error rates | API performance monitoring and SLA tracking |
| System Metrics | Goroutine count, memory usage, CPU utilization | Resource utilization and scaling decisions |
| Business Metrics | Currency conversion volumes, cache hit rates | Feature usage analytics and optimization opportunities |
| External API Metrics | API call frequency, response times, failure rates | Third-party dependency monitoring and cost optimization |

### Observability Stack

- **Prometheus Integration:** Custom metrics instrumentation with Go-Kit middleware
- **Grafana Dashboards:** Visual monitoring with alerts and threshold configuration
- **Structured Logging:** Comprehensive logging with contextual information
- **Health Checks:** Endpoint monitoring and dependency verification
- **Performance Tracking:** Real-time latency and throughput monitoring

### Available Dashboards

1. **Service Overview**: Request rates, response times, error rates
2. **System Metrics**: CPU, memory, goroutine usage
3. **Cache Performance**: Hit rates, TTL effectiveness
4. **External API Health**: Third-party service monitoring

---

## 10. Technical Decisions

### Key Architecture Choices

| Decision | Rationale | Benefits |
|----------|-----------|----------|
| USD as Base Currency | Minimize API calls for cross-currency calculations | Reduced external API dependency, improved performance |
| Separate Fiat/Crypto Caches | Different TTL requirements and data patterns | Optimized cache strategies, better data organization |
| In-Memory Caching | Fast access times and reduced API calls | Sub-second responses, cost optimization |
| Go-Kit Framework | Production-grade middleware and instrumentation | Built-in observability, clean architecture |
| 90-Day Historical Limit | Balance data freshness with storage efficiency | Relevant data retention, performance optimization |

### Technology Stack Justification

- **Go Language:** High concurrency support, excellent performance, strong standard library
- **Go-Kit Framework:** Microservices toolkit with built-in observability and middleware
- **Custom Caching:** Tailored TTL management and cleanup for specific use case
- **Docker Deployment:** Consistent environments and easy horizontal scaling
- **Prometheus/Grafana:** Industry-standard monitoring stack with rich ecosystem

---

## 📝 Development Setup

### Local Development

```bash
# Install dependencies
go mod download

# Run locally with hot reload
make dev

# Build binary
make build

# run with hot reloading
make serve

# Run tests
go test ./service
```

### Environment Variables

```bash
# .env file example
EXCHANGE_RATE_API_KEY=your_api_key_here
COINLAYER_API_KEY=your_coinlayer_api_key
PORT=8080
LOG_LEVEL=info
CACHE_TTL=3600

```

---



