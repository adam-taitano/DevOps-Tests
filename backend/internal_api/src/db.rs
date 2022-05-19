use sqlx::postgres::{PgPoolOptions, PgPool};
use crate::models::user::User;
use dotenv;

#[derive(Clone)]
pub struct Db {
    pub host: String,
    pub db_name: String,
    pub pool: Option<PgPool>
}

impl Db {
    pub async fn new(un:String, pw:String, h:String, db:String) -> Self {
        Self {
            host: h.clone(),
            db_name: db.clone(),
            pool: Db::pool(un, pw, h, db).await
        }
    }

    pub async fn from_env() -> Self {
        dotenv::dotenv().ok();
        let username = dotenv::var("DB_USERNAME");
        let password = dotenv::var("DB_PASSWORD");
        let host = dotenv::var("HOST");
        let db_name = dotenv::var("DB_NAME");

        match (username, password, host, db_name) {
            (Ok(un), Ok(pw), Ok(h), Ok(db)) => {
                Db::new(un, pw, h, db).await
            }
            (_, _, _, _) => {
                warn!("Postgres DB environment variables not set. Defaulting to localhost/postgres");
                Db::new("postgres".to_string(),
                        "postgres".to_string(),
                        "127.0.0.1".to_string(),
                        "postgres".to_string()).await
            }
        }
    }

    pub async fn pool(username:String, password:String, host:String, db:String) -> Option<PgPool> {
        let pgpool = PgPoolOptions::new()
            .max_connections(5)
            .connect(format!("postgres://{}:{}@{}/{}", username, password, host, db).as_str())
            .await;

        match pgpool {
            Ok(pool) => {
                info!("Connected to Postgres DB {}/{}", host, db);
                Some(pool)
            }
            Err(err) => {
                warn!("Database connection error: {}", err);
                None
            }
        }
    }

    pub async fn get_users(self) -> Option<Vec<User>> {
        match self.pool {
            Some(pool) => {
                match sqlx::query_as::<_, User>("SELECT * FROM users").fetch_all(&pool).await {
                    Ok(rows) => {Some(rows)}
                    Err(err) => {
                        warn!("Database query error: {}", err);
                        None
                    }
                }
            }
            None => {
                warn!("No database connections exist.");
                None
            }
        }
    }
}