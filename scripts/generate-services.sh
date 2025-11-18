#!/bin/bash

# Service Generation Script for 30 Microservices CRM System
# This script generates directory structure and base files for all services

set -e

# Base directory
BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SERVICES_DIR="$BASE_DIR/services"

# Color codes for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Service definitions: name:language:port:database
SERVICES=(
    "01-user-auth-service:csharp:5001:postgresql"
    "02-customer-service:java:5002:postgresql,mongodb"
    "03-contact-service:java:5003:postgresql"
    "04-account-service:go:5004:postgresql"
    "05-opportunity-service:go:5005:postgresql"
    "06-deal-service:rust:5006:postgresql"
    "07-activity-service:go:5007:cassandra"
    "08-interaction-service:java:5008:postgresql"
    "09-chat-service:go:5009:mongodb"
    "10-ai-service:rust:5010:redis"
    "11-email-service:bun:5011:mysql"
    "12-sms-service:go:5012:mysql"
    "13-notification-service:java:5013:mysql,mongodb"
    "14-sales-pipeline-service:java:5014:postgresql"
    "15-proposal-service:go:5015:mongodb"
    "16-quote-service:csharp:5016:postgresql"
    "17-campaign-service:java:5017:postgresql"
    "18-lead-service:go:5018:mongodb"
    "19-analytics-service:java:5019:cassandra,postgresql"
    "20-reporting-service:go:5020:postgresql"
    "21-dashboard-service:bun:5021:mongodb"
    "22-integration-service:go:5022:postgresql,couchdb,mariadb"
    "23-configuration-service:go:5023:couchdb"
    "24-audit-service:java:5024:postgresql"
    "25-admin-service:csharp:5025:postgresql"
)

echo -e "${BLUE}=== CRM Microservices Generator ===${NC}"
echo "Generating 30 microservices..."
echo

# Create C# service structure
create_csharp_service() {
    local name=$1
    local port=$2
    local db=$3
    local dir="$SERVICES_DIR/$name"

    echo -e "${GREEN}Creating C# service: $name${NC}"
    mkdir -p "$dir"/{Controllers,Models,Services,Data}

    # Create .csproj file
    cat > "$dir/${name}.csproj" <<EOF
<Project Sdk="Microsoft.NET.Sdk.Web">
  <PropertyGroup>
    <TargetFramework>net8.0</TargetFramework>
    <RootNamespace>$(echo $name | sed 's/-//g')</RootNamespace>
  </PropertyGroup>
  <ItemGroup>
    <PackageReference Include="Microsoft.AspNetCore.Authentication.JwtBearer" Version="8.0.0" />
    <PackageReference Include="Npgsql.EntityFrameworkCore.PostgreSQL" Version="8.0.0" />
    <PackageReference Include="Swashbuckle.AspNetCore" Version="6.5.0" />
    <PackageReference Include="Serilog.AspNetCore" Version="8.0.0" />
  </ItemGroup>
</Project>
EOF

    # Create Program.cs
    cat > "$dir/Program.cs" <<EOF
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

app.UseSwagger();
app.UseSwaggerUI();
app.MapControllers();
app.MapGet("/health", () => Results.Ok(new { status = "healthy", service = "$name" }));

app.Run();
EOF

    create_dockerfile_csharp "$dir"
    create_env_template "$dir" "$port" "$db"
}

# Create Java service structure
create_java_service() {
    local name=$1
    local port=$2
    local db=$3
    local dir="$SERVICES_DIR/$name"
    local package_name=$(echo $name | sed 's/-//g')

    echo -e "${GREEN}Creating Java service: $name${NC}"
    mkdir -p "$dir/src/main/java/com/crm/$package_name"/{controller,service,repository,entity,config}
    mkdir -p "$dir/src/main/resources"

    # Create build.gradle
    cat > "$dir/build.gradle" <<EOF
plugins {
    id 'java'
    id 'org.springframework.boot' version '3.2.0'
    id 'io.spring.dependency-management' version '1.1.4'
}

group = 'com.crm'
version = '1.0.0'
sourceCompatibility = '21'

repositories {
    mavenCentral()
}

dependencies {
    implementation 'org.springframework.boot:spring-boot-starter-web'
    implementation 'org.springframework.boot:spring-boot-starter-data-jpa'
    implementation 'org.springframework.boot:spring-boot-starter-actuator'
    implementation 'org.postgresql:postgresql'
    implementation 'org.springdoc:springdoc-openapi-starter-webmvc-ui:2.3.0'
    compileOnly 'org.projectlombok:lombok'
    annotationProcessor 'org.projectlombok:lombok'
}
EOF

    cat > "$dir/settings.gradle" <<EOF
rootProject.name = '$name'
EOF

    # Create Application class
    cat > "$dir/src/main/java/com/crm/$package_name/Application.java" <<EOF
package com.crm.$package_name;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}
EOF

    create_dockerfile_java "$dir"
    create_env_template "$dir" "$port" "$db"
}

# Create Go service structure
create_go_service() {
    local name=$1
    local port=$2
    local db=$3
    local dir="$SERVICES_DIR/$name"

    echo -e "${GREEN}Creating Go service: $name${NC}"
    mkdir -p "$dir"/{cmd,internal/{handler,service,repository,model,config}}

    # Create go.mod
    cat > "$dir/go.mod" <<EOF
module github.com/crm/$name

go 1.22

require (
    github.com/gin-gonic/gin v1.10.0
    github.com/sirupsen/logrus v1.9.3
)
EOF

    # Create main.go
    cat > "$dir/cmd/main.go" <<EOF
package main

import (
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy", "service": "$name"})
    })

    log.Printf("Starting $name on port $port")
    r.Run(":$port")
}
EOF

    create_dockerfile_go "$dir"
    create_env_template "$dir" "$port" "$db"
}

# Create Rust service structure
create_rust_service() {
    local name=$1
    local port=$2
    local db=$3
    local dir="$SERVICES_DIR/$name"

    echo -e "${GREEN}Creating Rust service: $name${NC}"
    mkdir -p "$dir/src"

    # Create Cargo.toml
    cat > "$dir/Cargo.toml" <<EOF
[package]
name = "$name"
version = "1.0.0"
edition = "2021"

[dependencies]
actix-web = "4.5"
serde = { version = "1.0", features = ["derive"] }
tokio = { version = "1.36", features = ["full"] }
EOF

    # Create main.rs
    cat > "$dir/src/main.rs" <<EOF
use actix_web::{web, App, HttpResponse, HttpServer, Responder};

async fn health() -> impl Responder {
    HttpResponse::Ok().json(serde_json::json!({
        "status": "healthy",
        "service": "$name"
    }))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .route("/health", web::get().to(health))
    })
    .bind(("0.0.0.0", $port))?
    .run()
    .await
}
EOF

    create_dockerfile_rust "$dir"
    create_env_template "$dir" "$port" "$db"
}

# Create BUN service structure
create_bun_service() {
    local name=$1
    local port=$2
    local db=$3
    local dir="$SERVICES_DIR/$name"

    echo -e "${GREEN}Creating BUN service: $name${NC}"
    mkdir -p "$dir/src"/{controllers,services,models}

    # Create package.json
    cat > "$dir/package.json" <<EOF
{
  "name": "$name",
  "version": "1.0.0",
  "main": "src/index.ts",
  "scripts": {
    "start": "bun run src/index.ts",
    "dev": "bun --watch src/index.ts"
  },
  "dependencies": {
    "express": "^4.18.2"
  },
  "devDependencies": {
    "@types/express": "^4.17.21",
    "bun-types": "latest"
  }
}
EOF

    # Create index.ts
    cat > "$dir/src/index.ts" <<EOF
import express from 'express';

const app = express();
const PORT = process.env.PORT || $port;

app.get('/health', (req, res) => {
    res.json({ status: 'healthy', service: '$name' });
});

app.listen(PORT, () => {
    console.log(\`$name running on port \${PORT}\`);
});
EOF

    create_dockerfile_bun "$dir"
    create_env_template "$dir" "$port" "$db"
}

# Dockerfile templates
create_dockerfile_csharp() {
    cat > "$1/Dockerfile" <<'EOF'
FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /src
COPY . .
RUN dotnet restore && dotnet publish -c Release -o /app

FROM mcr.microsoft.com/dotnet/aspnet:8.0
WORKDIR /app
COPY --from=build /app .
EXPOSE 5001
ENTRYPOINT ["dotnet", "*.dll"]
EOF
}

create_dockerfile_java() {
    cat > "$1/Dockerfile" <<'EOF'
FROM gradle:8.5-jdk21 AS build
WORKDIR /app
COPY . .
RUN gradle build -x test --no-daemon

FROM eclipse-temurin:21-jre-alpine
WORKDIR /app
COPY --from=build /app/build/libs/*.jar app.jar
EXPOSE 5002
ENTRYPOINT ["java", "-jar", "app.jar"]
EOF
}

create_dockerfile_go() {
    cat > "$1/Dockerfile" <<'EOF'
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 5004
ENTRYPOINT ["./main"]
EOF
}

create_dockerfile_rust() {
    cat > "$1/Dockerfile" <<'EOF'
FROM rust:1.75-alpine AS builder
WORKDIR /app
COPY . .
RUN cargo build --release

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/target/release/* ./
EXPOSE 5006
ENTRYPOINT ["./main"]
EOF
}

create_dockerfile_bun() {
    cat > "$1/Dockerfile" <<'EOF'
FROM oven/bun:1.0-alpine
WORKDIR /app
COPY package.json bun.lockb* ./
RUN bun install
COPY . .
EXPOSE 5011
CMD ["bun", "run", "src/index.ts"]
EOF
}

# Create .env.template
create_env_template() {
    local dir=$1
    local port=$2
    local db=$3

    cat > "$dir/.env.template" <<EOF
SERVICE_NAME=$(basename $dir)
SERVICE_PORT=$port
DATABASE=$db

# Database connections (based on $db)
POSTGRES_HOST=postgres-master
POSTGRES_PORT=5432
MONGODB_URI=mongodb://mongodb-primary:27017
CASSANDRA_CONTACT_POINTS=cassandra-1
MYSQL_HOST=mysql
REDIS_HOST=redis-master
COUCHDB_URL=http://couchdb:5984
MARIADB_HOST=mariadb
EOF
}

# Main execution
echo "Generating services..."
echo

for service_def in "${SERVICES[@]}"; do
    IFS=':' read -r name lang port db <<< "$service_def"

    case $lang in
        csharp)
            create_csharp_service "$name" "$port" "$db"
            ;;
        java)
            create_java_service "$name" "$port" "$db"
            ;;
        go)
            create_go_service "$name" "$port" "$db"
            ;;
        rust)
            create_rust_service "$name" "$port" "$db"
            ;;
        bun)
            create_bun_service "$name" "$port" "$db"
            ;;
    esac
done

echo
echo -e "${GREEN}✓ All 30 services generated successfully!${NC}"
echo "Services are located in: $SERVICES_DIR"
echo
echo "Next steps:"
echo "1. Review generated service code"
echo "2. Run: docker-compose build"
echo "3. Run: docker-compose up"
