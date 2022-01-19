<?php

header("Access-Control-Allow-Origin: *");
header("Content-Type: application/json; charset=UTF-8");
header("Access-Control-Allow-Methods: POST");
header("Access-Control-Max-Age: 3600");
header("Access-Control-Allow-Headers: Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With");

$uri = parse_url($_SERVER['REQUEST_URI'], PHP_URL_PATH);
$uri = explode('/', $uri);
$isEncodeReq = $uri[1] == 'encode';
$isDecodeReq = $uri[1] == 'decode';
// Our two endpoints are /encode and /decode
// everything else results in a 404 Not Found
if ($isDecodeReq == false && $isEncodeReq == false) {
    header("HTTP/1.1 404 Not Found");
    exit();
}

//This is the body of the POST request
$data = json_decode(file_get_contents('php://input'), true);

// For decode request, we should be receiving a string (jwt)
// And returning an object
if($isDecodeReq) {
    $decoded = base64_decode(str_replace('_', '/', str_replace('-','+',explode('.', $data)[1])));
    if(empty($decoded)) {
        echo json_encode('Invalid JWT');
        exit();
    }
    echo json_encode($decoded);
    exit();
}

// For an encode request we should be receiving an object
// And returning a string (jwt)

function base64UrlEncode($text)
{
    return str_replace(
        ['+', '/', '='],
        ['-', '_', ''],
        base64_encode($text)
    );
}

if($isEncodeReq) {

    $secret = 'secret';
    $header = ['typ' => 'JWT', 'alg' => 'HS256'];
    $headers_encoded = base64UrlEncode(json_encode($header));
    
    $payload_encoded = base64UrlEncode(json_encode($data));
    
    $signature = hash_hmac('SHA256', "$headers_encoded.$payload_encoded", $secret, true);
    $signature_encoded = base64UrlEncode($signature);
    
    $jwt = "$headers_encoded.$payload_encoded.$signature_encoded";
    
    echo json_encode($jwt);
    exit();
}


