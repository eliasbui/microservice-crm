# Microservices CRM with AI Chat Integration

A comprehensive enterprise-grade CRM system built with microservices architecture, featuring AI-powered chat assistance, real-time analytics, and multi-language service implementation.

## 🏗️ Architecture Overview

### Services
- **Auth Service** (C#/.NET Core) - JWT/OAuth 2.0 authentication and authorization
- **CRM Core Service** (Java/Spring Boot) - Customer and opportunity management
- **Chat & AI Service** (Golang) - Real-time chat with AI integration (Claude/OpenAI)
- **Automation Service** (Rust) - High-performance workflow automation
- **Notification Service** (BUN) - Email, SMS, and push notifications
- **Analytics Service** (Java) - Data aggregation and reporting
- **Integration Service** (Golang) - Third-party system integrations

### Frontend
- **Next.js + Vite** with TypeScript and Material UI
- Real-time dashboard and chat interface
- Mobile-responsive design with dark mode

### Infrastructure
- **API Gateway**: Nginx with load balancing and WAF
- **Databases**: PostgreSQL (master-slave), Redis Sentinel
- **Messaging**: RabbitMQ (message broker), Kafka (event streaming)
- **Monitoring**: Prometheus, Grafana, ELK Stack
- **Orchestration**: Kubernetes with Helm charts
- **Security**: OAuth 2.0, JWT, HashiCorp Vault

## 📁 Project Structure

```
microservice-crm/
├── services/
│   ├── auth-service/          # C#/.NET Core - Authentication
│   ├── crm-core-service/      # Java/Spring Boot - CRM Core
│   ├── chat-ai-service/       # Golang - Chat & AI
│   ├── automation-service/    # Rust - Workflow Automation
│   ├── notification-service/  # BUN - Notifications
│   ├── analytics-service/     # Java - Analytics & Reporting
│   └── integration-service/   # Golang - Third-party Integrations
├── frontend/
│   └── crm-web-app/          # Next.js + Vite + TypeScript
├── infrastructure/
│   ├── kubernetes/            # K8s manifests
│   ├── helm/                  # Helm charts
│   ├── docker-compose/        # Local development
│   ├── gateway/               # Nginx configuration
│   ├── monitoring/            # Prometheus, Grafana configs
│   └── logging/               # ELK Stack configs
├── database/
│   ├── migrations/            # Database migration scripts
│   └── schemas/               # Database schemas
├── scripts/
│   ├── setup/                 # Setup scripts
│   └── deployment/            # Deployment scripts
├── .github/
│   └── workflows/             # CI/CD pipelines
└── docs/                      # Documentation

```

## 🚀 Quick Start

### Prerequisites

- Docker & Docker Compose
- Kubernetes (minikube/kind for local)
- Helm 3.x
- Node.js 20+
- .NET 8 SDK
- Java 21 JDK
- Go 1.22+
- Rust 1.75+
- Bun 1.0+

### Local Development Setup

1. **Clone the repository**
```bash
git clone <repository-url>
cd microservice-crm
```

2. **Set up environment variables**
```bash
cp .env.example .env
# Edit .env with your configuration
```

3. **Start infrastructure services**
```bash
cd infrastructure/docker-compose
docker-compose up -d postgres redis rabbitmq kafka
```

4. **Start backend services**
```bash
docker-compose up -d
```

5. **Start frontend**
```bash
cd frontend/crm-web-app
npm install
npm run dev
```

6. **Access the application**
- Frontend: http://localhost:3000
- API Gateway: http://localhost:8080
- Grafana: http://localhost:3001
- Kibana: http://localhost:5601

## 🐳 Docker Compose

For local development with all services:

```bash
cd infrastructure/docker-compose
docker-compose up -d
```

This starts:
- All microservices
- PostgreSQL (master-slave)
- Redis Sentinel cluster
- RabbitMQ
- Kafka
- Prometheus & Grafana
- ELK Stack
- API Gateway

## ☸️ Kubernetes Deployment

### Using kubectl

```bash
# Create namespace
kubectl create namespace crm-system

# Apply configurations
kubectl apply -f infrastructure/kubernetes/namespace/
kubectl apply -f infrastructure/kubernetes/configmaps/
kubectl apply -f infrastructure/kubernetes/secrets/
kubectl apply -f infrastructure/kubernetes/deployments/
kubectl apply -f infrastructure/kubernetes/services/
kubectl apply -f infrastructure/kubernetes/ingress/
```

### Using Helm

```bash
# Install all services
helm install crm-system infrastructure/helm/crm-system \
  --namespace crm-system \
  --create-namespace \
  --values infrastructure/helm/crm-system/values.yaml

# Upgrade
helm upgrade crm-system infrastructure/helm/crm-system \
  --namespace crm-system
```

## 🔐 Security Setup

1. **Generate JWT Secret**
```bash
openssl rand -base64 32
```

2. **Set up Vault**
```bash
./scripts/setup/vault-setup.sh
```

3. **Configure OAuth 2.0**
- Edit `services/auth-service/appsettings.json`
- Set up client credentials

## 📊 Monitoring & Logging

### Prometheus Metrics
- Access: http://localhost:9090
- Scrapes metrics from all services

### Grafana Dashboards
- Access: http://localhost:3001
- Default credentials: admin/admin
- Pre-configured dashboards for each service

### Kibana Logs
- Access: http://localhost:5601
- All service logs aggregated via Filebeat/Logstash

## 🧪 Testing

```bash
# Run all tests
./scripts/test/run-all-tests.sh

# Individual services
cd services/auth-service && dotnet test
cd services/crm-core-service && ./gradlew test
cd services/chat-ai-service && go test ./...
cd services/automation-service && cargo test
cd services/notification-service && bun test
```

## 📡 API Documentation

- Swagger UI: http://localhost:8080/swagger
- OpenAPI specs available in each service's `/docs` directory

## 🔄 CI/CD Pipeline

GitHub Actions workflows:
- `.github/workflows/build.yml` - Build all services
- `.github/workflows/test.yml` - Run tests
- `.github/workflows/deploy-dev.yml` - Deploy to dev
- `.github/workflows/deploy-prod.yml` - Deploy to production

## 🗄️ Database

### Migrations

```bash
# Auth Service
cd services/auth-service
dotnet ef database update

# CRM Core Service
cd services/crm-core-service
./gradlew flywayMigrate

# Run all migrations
./scripts/deployment/migrate-databases.sh
```

### Backup & Restore

```bash
# Backup
./scripts/deployment/backup-database.sh

# Restore
./scripts/deployment/restore-database.sh backup-file.sql
```

## 🤖 AI Integration

### Configure AI Provider

Edit `.env`:
```bash
AI_PROVIDER=claude  # or openai
CLAUDE_API_KEY=your-api-key
OPENAI_API_KEY=your-api-key
```

### Supported Features
- Natural language customer queries
- Intelligent email drafting
- Sentiment analysis
- Recommendation engine
- Sales forecasting

## 🌐 Service Endpoints

| Service | Port | Health Check |
|---------|------|--------------|
| Auth Service | 5001 | /health |
| CRM Core Service | 8081 | /actuator/health |
| Chat & AI Service | 8082 | /health |
| Automation Service | 8083 | /health |
| Notification Service | 8084 | /health |
| Analytics Service | 8085 | /actuator/health |
| Integration Service | 8086 | /health |
| API Gateway | 8080 | /health |

## 📈 Scaling

### Horizontal Pod Autoscaling

```bash
kubectl autoscale deployment crm-core-service \
  --cpu-percent=70 \
  --min=2 \
  --max=10 \
  -n crm-system
```

### Manual Scaling

```bash
kubectl scale deployment crm-core-service \
  --replicas=5 \
  -n crm-system
```

## 🔧 Configuration

### Environment Variables

See `.env.example` for all available configuration options.

### Kubernetes ConfigMaps

Edit `infrastructure/kubernetes/configmaps/` for service configurations.

## 🛠️ Development

### Adding a New Service

1. Create service directory in `services/`
2. Add Dockerfile
3. Add to docker-compose.yml
4. Create Kubernetes manifests
5. Update Helm charts
6. Add to CI/CD pipeline

### Code Style

- C#: Follow Microsoft C# conventions
- Java: Google Java Style Guide
- Go: Effective Go guidelines
- Rust: Rust API guidelines
- TypeScript: Airbnb style guide

## 📝 License

MIT License - See LICENSE file for details

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📧 Support

For issues and questions:
- GitHub Issues: [Create an issue]
- Documentation: `docs/`
- Email: support@example.com

## 🗺️ Roadmap

- [ ] Mobile app (React Native)
- [ ] Advanced AI features (GPT-4 integration)
- [ ] Multi-tenancy support
- [ ] Advanced analytics with ML
- [ ] WhatsApp integration
- [ ] Video call support

---

**Built with ❤️ using modern microservices architecture**
