I needed a php library who did everything I needed to do for a project but the issue was that I didn't know how to write in PHP,
So I made this repo to start creating a simple rest api with PHP, but then I ended up extending this project to multiple langues, just in case I ever need to use a library from them:

PHP ✅

Go ✅

Python ✅

Rust ✅

C

C++ 

C#

Java 8

Assembly (For fun)


Each api will work around encoding a decoding a jwt

The first endpoint "/encode" will receive a json object and return a jwt

The second endpoint called "/decode" will receive a jwt and return its payload


Algorithm for both is the same

- Retrieve json from body
- Parse it
- Transform it
- Check and catch errors
- Convert to json
- Send result
