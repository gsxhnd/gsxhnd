use axum::{
    Json,
    extract::{Path, State},
    http::StatusCode,
    response::IntoResponse,
};
use chrono::Utc;
use tracing::{info, warn};

use crate::models::{CreateUserRequest, User};
use crate::state::AppState;

pub async fn get_users(State(state): State<AppState>) -> impl IntoResponse {
    info!("Fetching all users");
    let users = state.users.read().await;
    (StatusCode::OK, Json(users.to_vec()))
}

pub async fn create_user(
    State(state): State<AppState>,
    Json(payload): Json<CreateUserRequest>,
) -> impl IntoResponse {
    let user = User {
        id: rand::random::<u32>() % 10000,
        name: payload.name.clone(),
        email: payload.email.clone(),
        created_at: Utc::now().to_rfc3339(),
    };

    info!("Creating user: {} ({})", payload.name, payload.email);

    let mut users = state.users.write().await;
    users.push(user.clone());

    (StatusCode::CREATED, Json(user))
}

pub async fn get_user(Path(id): Path<u32>, State(state): State<AppState>) -> impl IntoResponse {
    info!("Fetching user with id: {}", id);

    let users = state.users.read().await;
    match users.iter().find(|u| u.id == id) {
        Some(user) => (StatusCode::OK, Json(Some(user.clone()))).into_response(),
        None => {
            warn!("User not found: {}", id);
            (StatusCode::NOT_FOUND, Json::<Option<User>>(None)).into_response()
        }
    }
}

pub async fn delete_user(Path(id): Path<u32>, State(state): State<AppState>) -> impl IntoResponse {
    info!("Deleting user with id: {}", id);

    let mut users = state.users.write().await;
    let initial_len = users.len();
    users.retain(|u| u.id != id);

    if users.len() < initial_len {
        (StatusCode::NO_CONTENT, "User deleted successfully").into_response()
    } else {
        warn!("User not found for deletion: {}", id);
        (StatusCode::NOT_FOUND, "User not found").into_response()
    }
}
