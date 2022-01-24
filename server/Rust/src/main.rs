#![feature(plugin, decl_macro, proc_macro_hygiene)]
#![allow(proc_macro_derive_resolution_fallback, unused_attributes)]

#[macro_use]
extern crate rocket;
extern crate rocket_contrib;
use rocket::http::Method;
use rocket_cors::{
    AllowedOrigins,
    Cors, CorsOptions,
};
use routes::*;

mod routes;

fn main() {
    rocket().launch();
}

fn rocket() -> rocket::Rocket {
    rocket::ignite()
        .mount(
            "/",
            routes![encode, decode],
        )
        .attach(make_cors())
}

fn make_cors() -> Cors {
    let allowed_origins = AllowedOrigins::all();

    CorsOptions {
        allowed_origins,
        allowed_methods: vec![Method::Post].into_iter().map(From::from).collect(),
        allow_credentials: false,
        ..Default::default()
    }
        .to_cors()
        .expect("error while building CORS")
}
