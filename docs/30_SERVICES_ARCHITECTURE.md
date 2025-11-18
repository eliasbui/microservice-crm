# 30 Microservices Architecture

## Complete Service Directory Structure

```
microservice-crm/
├── services/
│   ├── 01-user-auth-service/          (C#/.NET)    - Port 5001
│   ├── 02-customer-service/           (Java)       - Port 5002
│   ├── 03-contact-service/            (Java)       - Port 5003
│   ├── 04-account-service/            (Go)         - Port 5004
│   ├── 05-opportunity-service/        (Go)         - Port 5005
│   ├── 06-deal-service/               (Rust)       - Port 5006
│   ├── 07-activity-service/           (Go)         - Port 5007
│   ├── 08-interaction-service/        (Java)       - Port 5008
│   ├── 09-chat-service/               (Go)         - Port 5009
│   ├── 10-ai-service/                 (Rust)       - Port 5010
│   ├── 11-email-service/              (BUN)        - Port 5011
│   ├── 12-sms-service/                (Go)         - Port 5012
│   ├── 13-notification-service/       (Java)       - Port 5013
│   ├── 14-sales-pipeline-service/     (Java)       - Port 5014
│   ├── 15-proposal-service/           (Go)         - Port 5015
│   ├── 16-quote-service/              (C#/.NET)    - Port 5016
│   ├── 17-campaign-service/           (Java)       - Port 5017
│   ├── 18-lead-service/               (Go)         - Port 5018
│   ├── 19-analytics-service/          (Java)       - Port 5019
│   ├── 20-reporting-service/          (Go)         - Port 5020
│   ├── 21-dashboard-service/          (BUN)        - Port 5021
│   ├── 22-integration-service/        (Go)         - Port 5022
│   ├── 23-configuration-service/      (Go)         - Port 5023
│   ├── 24-audit-service/              (Java)       - Port 5024
│   ├── 25-admin-service/              (C#/.NET)    - Port 5025
│   ├── 26-api-gateway/                (Nginx)      - Port 8080
│   ├── 27-service-discovery/          (Consul)     - Port 8500
│   ├── 28-config-server/              (Spring)     - Port 8888
│   ├── 29-message-bus/                (RabbitMQ)   - Port 5672
│   └── 30-cache-service/              (Redis)      - Port 6379
```

## Service-Database Mapping

| Service | Primary DB | Secondary DB | Purpose |
|---------|-----------|--------------|---------|
| 01-user-auth-service | PostgreSQL | Redis | Identity, JWT, sessions |
| 02-customer-service | PostgreSQL | MongoDB | Customer master data + profiles |
| 03-contact-service | PostgreSQL | - | Contact information |
| 04-account-service | PostgreSQL | - | Account hierarchies |
| 05-opportunity-service | PostgreSQL | - | Sales opportunities |
| 06-deal-service | PostgreSQL | - | Deal management |
| 07-activity-service | Cassandra | - | Time-series activities |
| 08-interaction-service | PostgreSQL | - | Communication history |
| 09-chat-service | MongoDB | Redis | Real-time messaging |
| 10-ai-service | - | Redis | AI/ML processing |
| 11-email-service | MySQL | - | Email templates & sending |
| 12-sms-service | MySQL | - | SMS gateway |
| 13-notification-service | MySQL | MongoDB | Multi-channel notifications |
| 14-sales-pipeline-service | PostgreSQL | - | Pipeline management |
| 15-proposal-service | MongoDB | - | Proposal documents |
| 16-quote-service | PostgreSQL | - | Quote generation |
| 17-campaign-service | PostgreSQL | - | Marketing campaigns |
| 18-lead-service | MongoDB | - | Lead tracking & scoring |
| 19-analytics-service | Cassandra | PostgreSQL | Time-series + aggregations |
| 20-reporting-service | PostgreSQL | - | Report generation |
| 21-dashboard-service | MongoDB | Redis | Dashboard configurations |
| 22-integration-service | PostgreSQL | CouchDB+MariaDB | Third-party integrations |
| 23-configuration-service | CouchDB | - | Feature flags |
| 24-audit-service | PostgreSQL | - | Audit trails |
| 25-admin-service | PostgreSQL | - | System administration |

## Standard Directory Structure (Per Service)

### C# Services (01, 16, 25)
```
{service-name}/
├── Controllers/
│   └── {Entity}Controller.cs
├── Models/
│   ├── Entities/
│   └── DTOs/
├── Services/
│   ├── I{Service}.cs
│   └── {Service}.cs
├── Data/
│   └── ApplicationDbContext.cs
├── {ServiceName}.csproj
├── Program.cs
├── appsettings.json
├── Dockerfile
├── .dockerignore
└── .env.template
```

### Java Services (02, 03, 08, 13, 14, 17, 19, 24)
```
{service-name}/
├── src/
│   ├── main/
│   │   ├── java/com/crm/{service}/
│   │   │   ├── {Service}Application.java
│   │   │   ├── controller/
│   │   │   ├── service/
│   │   │   ├── repository/
│   │   │   ├── entity/
│   │   │   ├── dto/
│   │   │   └── config/
│   │   └── resources/
│   │       ├── application.yml
│   │       └── db/migration/
│   └── test/
├── build.gradle
├── settings.gradle
├── Dockerfile
└── .env.template
```

### Go Services (04, 05, 07, 09, 12, 15, 18, 20, 22, 23)
```
{service-name}/
├── cmd/
│   └── main.go
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
│   └── config/
├── pkg/
│   └── utils/
├── go.mod
├── go.sum
├── Dockerfile
└── .env.template
```

### Rust Services (06, 10)
```
{service-name}/
├── src/
│   ├── main.rs
│   ├── handlers/
│   ├── services/
│   ├── models/
│   └── config.rs
├── Cargo.toml
├── Cargo.lock
├── Dockerfile
└── .env.template
```

### BUN Services (11, 21)
```
{service-name}/
├── src/
│   ├── index.ts
│   ├── controllers/
│   ├── services/
│   ├── models/
│   └── config/
├── package.json
├── tsconfig.json
├── Dockerfile
└── .env.template
```

## Kubernetes Structure (Per Service)

```
infrastructure/kubernetes/services/
└── {service-name}/
    ├── deployment.yaml
    ├── service.yaml
    ├── configmap.yaml
    ├── hpa.yaml           (Horizontal Pod Autoscaler)
    └── pdb.yaml           (Pod Disruption Budget)
```

## Service Communication Patterns

### Synchronous (HTTP/REST)
- API Gateway → All Services
- Service-to-Service direct calls (limited)

### Asynchronous (Events)
- RabbitMQ: Command/Task queues
- Kafka: Event streaming
- Redis Pub/Sub: Real-time notifications

### Service Discovery
- Consul for service registration
- DNS-based discovery in Kubernetes

## API Gateway Routes

```
/api/v1/auth/*          → user-auth-service
/api/v1/customers/*     → customer-service
/api/v1/contacts/*      → contact-service
/api/v1/accounts/*      → account-service
/api/v1/opportunities/* → opportunity-service
/api/v1/deals/*         → deal-service
/api/v1/activities/*    → activity-service
/api/v1/interactions/*  → interaction-service
/api/v1/chat/*          → chat-service
/api/v1/ai/*            → ai-service
/api/v1/email/*         → email-service
/api/v1/sms/*           → sms-service
/api/v1/notifications/* → notification-service
/api/v1/pipeline/*      → sales-pipeline-service
/api/v1/proposals/*     → proposal-service
/api/v1/quotes/*        → quote-service
/api/v1/campaigns/*     → campaign-service
/api/v1/leads/*         → lead-service
/api/v1/analytics/*     → analytics-service
/api/v1/reports/*       → reporting-service
/api/v1/dashboards/*    → dashboard-service
/api/v1/integrations/*  → integration-service
/api/v1/config/*        → configuration-service
/api/v1/audit/*         → audit-service
/api/v1/admin/*         → admin-service
```

## Technology Stack Summary

### Languages Distribution
- **C# (.NET 8)**: 3 services (12%)
- **Java (Spring Boot 3)**: 8 services (32%)
- **Go (1.22)**: 10 services (40%)
- **Rust (1.75)**: 2 services (8%)
- **BUN (1.0)**: 2 services (8%)

### Databases Distribution
- **PostgreSQL**: 15 services (primary)
- **MongoDB**: 6 services
- **Cassandra**: 2 services
- **MySQL**: 4 services
- **CouchDB**: 2 services
- **MariaDB**: 1 service
- **Redis**: All services (caching)

## Inter-Service Dependencies

### Core Dependencies
```
user-auth-service → (all services depend on this for authentication)
customer-service → contact-service, account-service
opportunity-service → customer-service, account-service
deal-service → opportunity-service
```

### Communication Dependencies
```
notification-service → email-service, sms-service
chat-service → ai-service
```

### Analytics Dependencies
```
analytics-service → activity-service, interaction-service
reporting-service → analytics-service
dashboard-service → analytics-service, reporting-service
```

## Deployment Order

### Phase 1: Infrastructure
1. message-bus (RabbitMQ)
2. cache-service (Redis)
3. service-discovery (Consul)
4. config-server (Spring Cloud)

### Phase 2: Core Services
5. user-auth-service
6. customer-service
7. contact-service
8. account-service

### Phase 3: Domain Services
9. opportunity-service
10. deal-service
11. activity-service
12. interaction-service

### Phase 4: Communication
13. chat-service
14. ai-service
15. email-service
16. sms-service
17. notification-service

### Phase 5: Sales & Marketing
18. sales-pipeline-service
19. proposal-service
20. quote-service
21. campaign-service
22. lead-service

### Phase 6: Analytics
23. analytics-service
24. reporting-service
25. dashboard-service

### Phase 7: Integration & Admin
26. integration-service
27. configuration-service
28. audit-service
29. admin-service

### Phase 8: Gateway
30. api-gateway

## Resource Requirements (Kubernetes)

### Small Services (Go, BUN)
```yaml
requests:
  memory: "256Mi"
  cpu: "250m"
limits:
  memory: "512Mi"
  cpu: "500m"
replicas: 2
```

### Medium Services (Java, C#)
```yaml
requests:
  memory: "512Mi"
  cpu: "500m"
limits:
  memory: "1Gi"
  cpu: "1000m"
replicas: 2-3
```

### Large Services (Rust, Heavy Processing)
```yaml
requests:
  memory: "1Gi"
  cpu: "1000m"
limits:
  memory: "2Gi"
  cpu: "2000m"
replicas: 3-5
```

## Monitoring & Observability

### Metrics (Prometheus)
- Each service exposes `/metrics` endpoint
- Custom business metrics per service
- JVM metrics for Java services
- Runtime metrics for Go/Rust

### Logging (ELK)
- Structured JSON logging
- Correlation IDs for tracing
- Log levels: ERROR, WARN, INFO, DEBUG

### Tracing (Jaeger/Zipkin)
- Distributed tracing across services
- Request flow visualization
- Performance bottleneck identification

### Health Checks
- `/health` - Overall health
- `/ready` - Readiness probe
- `/live` - Liveness probe

## Security Considerations

### Authentication
- JWT tokens from user-auth-service
- Token validation in API Gateway
- Service-to-service mTLS (optional)

### Authorization
- Role-based access control (RBAC)
- Resource-level permissions
- API scopes

### Network Security
- Service mesh (Istio/Linkerd) for mTLS
- Network policies in Kubernetes
- Secrets management (Vault)

## Scaling Strategy

### Horizontal Scaling
- HPA based on CPU/Memory
- Custom metrics (request rate, queue depth)
- Min/Max replica configuration

### Vertical Scaling
- Resource limit adjustments
- JVM heap tuning for Java services

### Database Scaling
- Read replicas for PostgreSQL
- Sharding for MongoDB
- Cluster expansion for Cassandra

## Disaster Recovery

### Backup Strategy
- Database backups (daily)
- Configuration backups
- Docker image registry backups

### High Availability
- Multi-zone deployment
- Database replication
- Stateless service design

### Failover
- Automatic pod rescheduling
- Database automatic failover
- Load balancer health checks
