use serde::{Serialize, Deserialize};
use sqlx::{FromRow};

#[derive(Serialize, Deserialize, FromRow)]
pub struct User {
    pub id: i32,
    pub username: String,
}

impl User {
    pub fn new(id:i32, username: String) -> Self {
        return Self {
            id: id,
            username: username
        }
    }
}

pub fn test_users() -> [User; 3] {
    [User::new(1, "testuser1".to_string()),
    User::new(2, "testuser2".to_string()),
    User::new(3, "testuser3".to_string())]
}