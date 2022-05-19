#[macro_use]
extern crate log;
mod models;
mod logger;
mod api;
mod config;
mod db;

use actix_web::{get, App, web, HttpResponse, HttpServer, Responder};
use std::io;

struct SessionData {
    db: db::Db
}


#[get("/")] // test landing page
async fn test() -> impl Responder {
    info!("GET /");
    HttpResponse::Ok().body("Backend running.")
}

#[actix_web::main]
async fn main() -> io::Result<()> {
    match logger::init() { // initialize logger for info/debugging
        Err(err) => {
            println!("Unable to start logger due to {}", err);
        }
        _ => {
            info!("Logger initialized");
        }
    }

    let config = config::Config::new();
    info!("Running backend at {}:{}", config.host, config.port); // log server start
    
    // start db connection
    let db:db::Db = db::Db::from_env().await;

    // start server at HOST:PORT, persisting Db connection
    HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(SessionData {db: db.clone()}))
            .service(test)
            .service(api::config())
    })
    .bind((config.host, config.port))?
    .run()
    .await
}