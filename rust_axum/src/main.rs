use rust_axum::{AppState, Config, create_router};
use tracing::info;

#[tokio::main]
async fn main() {
    // 初始化日志
    tracing_subscriber::fmt::init();
    info!("Starting Rust Axum server...");

    // 加载配置
    let config = Config::from_env();
    info!("Server configuration: {}", config.address());

    // 创建应用状态
    let app_state = AppState::new();

    // 创建路由
    let app = create_router(app_state);

    // 启动服务器
    let listener = tokio::net::TcpListener::bind(&config.address())
        .await
        .expect(&format!("Failed to bind port {}", config.port));

    info!("🚀 Server running on http://{}", config.address());

    axum::serve(listener, app).await.expect("Server error");
}
