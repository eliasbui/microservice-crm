// Automation module for handling workflow automation tasks

use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct AutomationTask {
    pub id: String,
    pub task_type: TaskType,
    pub status: TaskStatus,
}

#[derive(Debug, Serialize, Deserialize)]
pub enum TaskType {
    EmailCampaign,
    LeadNurturing,
    DataSync,
    ReportGeneration,
}

#[derive(Debug, Serialize, Deserialize)]
pub enum TaskStatus {
    Pending,
    Running,
    Completed,
    Failed,
}
