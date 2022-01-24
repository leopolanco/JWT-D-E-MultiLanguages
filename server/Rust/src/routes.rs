use rocket_contrib::json::Json;

#[post("/encode", format = "application/json")]
pub fn encode() -> Json<&'static str>  {
    println!("HIT ON /ENCODE");
    Json("Check on /encode")
}
#[post("/decode", format = "application/json")]
pub fn decode() -> Json<&'static str>  {
    println!("HIT ON /DECODE");
    Json("Check on /decode")
}