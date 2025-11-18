use actix_web::{web, App, HttpResponse, HttpServer, Responder};
use serde::{Deserialize, Serialize};
use std::env;

mod automation;
mod workflow;

#[derive(Debug, Serialize, Deserialize)]
struct HealthResponse {
    status: String,
    service: String,
}

#[derive(Debug, Serialize, Deserialize)]
struct WorkflowRequest {
    workflow_type: String,
    customer_id: Option<i64>,
    trigger: String,
    parameters: serde_json::Value,
}

#[derive(Debug, Serialize, Deserialize)]
struct WorkflowResponse {
    workflow_id: String,
    status: String,
    message: String,
}

async fn health() -> impl Responder {
    HttpResponse::Ok().json(HealthResponse {
        status: "healthy".to_string(),
        service: "automation-service".to_string(),
    })
}

async fn ready() -> impl Responder {
    HttpResponse::Ok().json(HealthResponse {
        status: "ready".to_string(),
        service: "automation-service".to_string(),
    })
}

async fn live() -> impl Responder {
    HttpResponse::Ok().json(HealthResponse {
        status: "alive".to_string(),
        service: "automation-service".to_string(),
    })
}

async fn create_workflow(req: web::Json<WorkflowRequest>) -> impl Responder {
    log::info!("Creating workflow: {:?}", req.workflow_type);

    let workflow_id = uuid::Uuid::new_v4().to_string();

    HttpResponse::Ok().json(WorkflowResponse {
        workflow_id,
        status: "created".to_string(),
        message: format!("Workflow {} created successfully", req.workflow_type),
    })
}

async fn execute_workflow(workflow_id: web::Path<String>) -> impl Responder {
    log::info!("Executing workflow: {}", workflow_id);

    HttpResponse::Ok().json(WorkflowResponse {
        workflow_id: workflow_id.to_string(),
        status: "executing".to_string(),
        message: "Workflow execution started".to_string(),
    })
}

async fn get_workflow_status(workflow_id: web::Path<String>) -> impl Responder {
    log::info!("Getting workflow status: {}", workflow_id);

    HttpResponse::Ok().json(WorkflowResponse {
        workflow_id: workflow_id.to_string(),
        status: "running".to_string(),
        message: "Workflow is running".to_string(),
    })
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    env_logger::init_from_env(env_logger::Env::new().default_filter_or("info"));
    dotenv::dotenv().ok();

    let port = env::var("AUTOMATION_SERVICE_PORT").unwrap_or_else(|_| "8083".to_string());
    let bind_address = format!("0.0.0.0:{}", port);

    log::info!("Starting Automation Service on {}", bind_address);

    HttpServer::new(|| {
        App::new()
            .route("/health", web::get().to(health))
            .route("/ready", web::get().to(ready))
            .route("/live", web::get().to(live))
            .service(
                web::scope("/api/workflows")
                    .route("", web::post().to(create_workflow))
                    .route("/{workflow_id}/execute", web::post().to(execute_workflow))
                    .route("/{workflow_id}/status", web::get().to(get_workflow_status))
            )
    })
    .bind(&bind_address)?
    .run()
    .await
}
