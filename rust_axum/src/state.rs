use crate::models::User;
use std::sync::Arc;
use tokio::sync::RwLock;

#[derive(Clone)]
pub struct AppState {
    pub users: Arc<RwLock<Vec<User>>>,
}

impl AppState {
    pub fn new() -> Self {
        Self {
            users: Arc::new(RwLock::new(vec![])),
        }
    }
}
