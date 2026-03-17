pub struct Config {
    pub host: String,
    pub port: u16,
}

impl Config {
    pub fn new() -> Self {
        Self {
            host: "127.0.0.1".to_string(),
            port: 3000,
        }
    }

    pub fn from_env() -> Self {
        let host = std::env::var("HOST").unwrap_or_else(|_| "127.0.0.1".to_string());
        let port = std::env::var("PORT")
            .ok()
            .and_then(|p| p.parse().ok())
            .unwrap_or(3000);

        Self { host, port }
    }

    pub fn address(&self) -> String {
        format!("{}:{}", self.host, self.port)
    }
}
