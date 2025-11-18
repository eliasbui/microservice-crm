# Database Architecture Documentation

## Overview

The CRM microservices system uses a polyglot persistence architecture with 6 different database technologies, each optimized for specific use cases and data patterns.

## Database Technologies

### 1. PostgreSQL (Relational Database)
**Version:** 16
**Configuration:** Master-Slave Replication

**Use Cases:**
- **CRM Core Service**: Transactional data (Customers, Opportunities)
- **Analytics Service**: Aggregated reports and metrics
- **Integration Service**: Integration configurations and sync history

**Features:**
- ACID compliance for data integrity
- Advanced SQL capabilities
- Referential integrity with foreign keys
- Strong consistency
- Master-slave replication for read scaling

**Connection Details:**
- Master: `postgres-master:5432`
- Slave: `postgres-slave:5433`

---

### 2. MongoDB (Document Database)
**Version:** 7.0
**Configuration:** 3-Node Replica Set (rs0)

**Use Cases:**
- **CRM Core Service**: Customer profiles with flexible schema
- **Chat & AI Service**: Chat messages and sessions
- **Notification Service**: Notification history

**Features:**
- Flexible schema for evolving data models
- Rich query capabilities with aggregation pipeline
- Horizontal scaling with sharding
- Document validation
- Change streams for real-time data

**Connection String:**
```
mongodb://crm_user:password@mongodb-primary:27017,mongodb-secondary1:27017,mongodb-secondary2:27017/crm_db?replicaSet=rs0
```

**Collections:**
- `customer_profiles`: Rich customer data with nested documents
- `chat_messages`: Conversation history
- `chat_sessions`: Active and historical chat sessions
- `notifications`: Notification delivery tracking

---

### 3. Apache Cassandra (Wide-Column Store)
**Version:** 4.1
**Configuration:** 3-Node Cluster

**Use Cases:**
- **CRM Core Service**: Customer event logs (time-series)
- **Analytics Service**: Time-series metrics and KPIs

**Features:**
- Linear scalability
- High write throughput
- Tunable consistency
- Time-to-live (TTL) for automatic data expiration
- Optimized for time-series data

**Contact Points:**
- cassandra-1:9042
- cassandra-2:9043
- cassandra-3:9044

**Keyspace:** `crm_keyspace`

**Tables:**
- `customer_event_logs`: Click streams, page views, interactions
- `analytics_metrics`: Hourly/daily aggregated metrics
- `user_activity_summary`: User activity tracking
- `customer_engagement_scores`: Time-series engagement data

---

### 4. MySQL (Relational Database)
**Version:** 8.0

**Use Cases:**
- **Notification Service**: Email templates and campaigns

**Features:**
- ACID compliance
- JSON column type for semi-structured data
- Full-text search capabilities
- Stored procedures and triggers

**Connection:**
- Host: `mysql:3306`
- Database: `crm_notifications`

**Tables:**
- `notification_templates`: Reusable email/SMS/push templates
- `email_campaigns`: Campaign management
- `campaign_recipients`: Recipient tracking
- `notification_queue`: Async notification processing

---

### 5. MariaDB (Relational Database)
**Version:** 10.6

**Use Cases:**
- **Integration Service**: Integration configurations and sync logs

**Features:**
- MySQL-compatible
- Enhanced performance features
- Columnar storage engine (ColumnStore)
- Advanced replication

**Connection:**
- Host: `mariadb:3307`
- Database: `crm_integrations`

**Tables:**
- `integration_connections`: Third-party system credentials
- `sync_jobs`: ETL job tracking
- `sync_history`: Detailed sync logs
- `field_mappings`: Data transformation rules
- `webhook_endpoints`: Outbound webhooks

---

### 6. CouchDB (Document Database)
**Version:** 3.2

**Use Cases:**
- **Integration Service**: Flexible schema integrations, document storage

**Features:**
- Schema-free JSON documents
- HTTP/REST API
- Master-master replication
- Conflict resolution (MVCC)
- MapReduce views

**Connection:**
- URL: `http://couchdb:5984`

**Databases:**
- `integration_configs`: Integration configuration documents
- `flexible_entities`: Entities from various sources with varying schemas
- `document_store`: File metadata and attachments
- `audit_logs`: Comprehensive audit trail

---

## Service-to-Database Mapping

### CRM Core Service
| Database | Purpose |
|----------|---------|
| PostgreSQL | Core transactional data (customers, opportunities) |
| MongoDB | Customer profiles with rich metadata |
| Cassandra | Event logs and time-series data |

### Chat & AI Service
| Database | Purpose |
|----------|---------|
| MongoDB | Chat messages and sessions |
| Redis | Session state and caching |

### Analytics Service
| Database | Purpose |
|----------|---------|
| Cassandra | Time-series metrics |
| PostgreSQL | Aggregated reports |

### Notification Service
| Database | Purpose |
|----------|---------|
| MongoDB | Notification history |
| MySQL | Templates and campaigns |

### Integration Service
| Database | Purpose |
|----------|---------|
| PostgreSQL | Structured integration logs |
| MariaDB | Sync jobs and mappings |
| CouchDB | Flexible entity storage |

---

## Connection Pooling Configuration

### PostgreSQL (HikariCP)
```yaml
maximum-pool-size: 10
minimum-idle: 2
connection-timeout: 30000
idle-timeout: 600000
max-lifetime: 1800000
```

### MongoDB
```yaml
maxPoolSize: 50
minPoolSize: 10
maxIdleTimeMS: 60000
```

### Cassandra
```yaml
pool:
  local:
    size: 4
  remote:
    size: 2
```

### MySQL/MariaDB
```yaml
max-pool-size: 20
min-pool-size: 5
connection-timeout: 30000
```

---

## Backup Strategies

### PostgreSQL
- **Method**: WAL archiving + pg_basebackup
- **Frequency**: Daily full backup, continuous WAL archiving
- **Retention**: 30 days

### MongoDB
- **Method**: mongodump or Ops Manager
- **Frequency**: Daily snapshots
- **Retention**: 7 days for daily, 4 weeks for weekly

### Cassandra
- **Method**: Snapshots + incremental backups
- **Frequency**: Daily snapshots
- **Retention**: 14 days

### MySQL/MariaDB
- **Method**: mysqldump or binary log replication
- **Frequency**: Daily full backup
- **Retention**: 14 days

### CouchDB
- **Method**: Replication to backup server
- **Frequency**: Continuous replication
- **Retention**: Point-in-time recovery up to 30 days

---

## Performance Optimization

### Indexing Strategy
- **PostgreSQL**: B-tree indexes on foreign keys, partial indexes for filtered queries
- **MongoDB**: Compound indexes on frequently queried fields
- **Cassandra**: Partition key design for even data distribution
- **MySQL/MariaDB**: Covering indexes for common queries

### Caching Strategy
- **Redis**: Hot data caching with 5-minute TTL
- **Application-level**: Result caching for complex aggregations

### Partitioning
- **PostgreSQL**: Table partitioning by date for historical data
- **MongoDB**: Sharding by customer_id for horizontal scaling
- **Cassandra**: Time-based partitioning for event logs

---

## Monitoring

### Metrics to Monitor
1. **Connection Pool**: Active connections, wait time
2. **Query Performance**: Slow queries, execution time
3. **Replication Lag**: Master-slave delay (PostgreSQL, MongoDB)
4. **Disk Usage**: Storage utilization and growth rate
5. **Cache Hit Ratio**: Redis cache effectiveness

### Tools
- **Prometheus**: Metric collection from all databases
- **Grafana**: Visualization and alerting
- **Database-specific**:
  - PostgreSQL: pg_stat_statements
  - MongoDB: MongoDB Cloud Manager
  - Cassandra: nodetool, DataStax OpsCenter

---

## Migration Scripts Location

```
database/
├── cassandra/init/
│   └── 01_create_keyspace.cql
├── mongodb/init/
│   └── 01_init_replica_set.js
├── mysql/init/
│   └── 01_create_tables.sql
├── mariadb/init/
│   └── 01_create_tables.sql
├── couchdb/init/
│   └── 01_init_databases.sh
└── schemas/
    └── 01_create_tables.sql  (PostgreSQL)
```

---

## Best Practices

1. **Use the Right Database for the Right Job**
   - Transactional data → PostgreSQL
   - Flexible schemas → MongoDB/CouchDB
   - Time-series → Cassandra
   - Templates/Campaigns → MySQL/MariaDB

2. **Connection Management**
   - Always use connection pooling
   - Close connections properly
   - Monitor pool exhaustion

3. **Data Consistency**
   - Use transactions for multi-table operations (SQL databases)
   - Implement application-level consistency for cross-database operations
   - Handle eventual consistency in NoSQL databases

4. **Security**
   - Use separate credentials for each service
   - Encrypt connections (TLS/SSL)
   - Regular security patches
   - Principle of least privilege

5. **Disaster Recovery**
   - Regular backup testing
   - Document recovery procedures
   - Maintain backup retention policies
   - Test failover scenarios

---

## Quick Reference

### Database Ports
| Database | Port |
|----------|------|
| PostgreSQL Master | 5432 |
| PostgreSQL Slave | 5433 |
| MongoDB Primary | 27017 |
| MongoDB Secondary 1 | 27018 |
| MongoDB Secondary 2 | 27019 |
| Cassandra Node 1 | 9042 |
| Cassandra Node 2 | 9043 |
| Cassandra Node 3 | 9044 |
| MySQL | 3306 |
| MariaDB | 3307 |
| CouchDB | 5984 |

### Admin URLs
- **RabbitMQ**: http://localhost:15672
- **CouchDB**: http://localhost:5984/_utils
- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3001
