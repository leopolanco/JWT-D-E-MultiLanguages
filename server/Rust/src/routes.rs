use std::collections::HashMap;
use rocket_contrib::json::Json;
use serde_json::Value;
use jsonwebtoken::{encode as jwt_encode, dangerous_insecure_decode as jwt_decode, Header, EncodingKey};

#[post("/encode", format = "application/json", data="<payload>")]
pub fn encode(payload: Json<HashMap<&str, Value>>) -> Json<String>  {
    let parsed_object = payload.into_inner();
    let token = jwt_encode(&Header::default(), &parsed_object, &EncodingKey::from_secret("secret".as_ref())).unwrap();
    Json(token)
}

#[post("/decode", format = "application/json", data="<received_jwt>")]
pub fn decode(received_jwt: Json<String>) -> Json<String>  {
    let string_jwt = received_jwt.into_inner();

    let parsed_jwt = jwt_decode::<Value>(&string_jwt);
    match parsed_jwt {
        Ok(decoded_jwt) => {
            let parsed_string_jwt = serde_json::to_string(&decoded_jwt.claims).unwrap();
            Json(parsed_string_jwt)
        },
        Err(e) => {
            println!("Error when parsing jwt {}", e);
            Json("Invalid Token".to_owned())
        }
    }
}