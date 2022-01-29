/**
 * Install dependencies before compiling:
 * $ sudo apt install libulfius-dev uwsc
 */

#include <stdio.h>
#include <string.h>

// Using ulfius as web framework
#include <ulfius.h>
#include <jansson.h>


#define PORT 8080

char *strconcat(int num_args, ...) {
    int strsize = 0;
    va_list ap;
    va_start(ap, num_args);
    for (int i = 0; i < num_args; i++) 
        strsize += strlen(va_arg(ap, char*));

    char *res = malloc(strsize+1);
    strsize = 0;
    va_start(ap, num_args);
    for (int i = 0; i < num_args; i++) {
        char *s = va_arg(ap, char*);
        strcpy(res+strsize, s);
        strsize += strlen(s);
    }
    va_end(ap);
    res[strsize] = '\0';

    return res;
}


//Preflight response headers
int callback_options (const struct _u_request * request, struct _u_response * response, void * user_data) {
  u_map_put(response->map_header, "Access-Control-Allow-Origin", "*");
  u_map_put(response->map_header, "Access-Control-Allow-Methods", "POST");
  u_map_put(response->map_header, "Content-Type", "application/json");
  u_map_put(response->map_header, "Access-Control-Allow-Headers", "Origin, Content-Type");

  return U_CALLBACK_COMPLETE;
}

int callback_decode (const struct _u_request * request, struct _u_response * response, void * user_data) {
  u_map_put(response->map_header, "Access-Control-Allow-Origin", "*");
  u_map_put(response->map_header, "Content-Type", "application/json");


  json_t * json_received_jwt = ulfius_get_json_body_request(request, NULL);

  //Change third parameter to decoded jwt
  char *jwt_as_string = strconcat(3, "\"", json_string_value(json_received_jwt), "\"");

  ulfius_set_string_body_response(response, 200, jwt_as_string);
  json_decref(json_received_jwt);
  return U_CALLBACK_CONTINUE;
}

int callback_encode (const struct _u_request * request, struct _u_response * response, void * user_data) {
  u_map_put(response->map_header, "Access-Control-Allow-Origin", "*");
  u_map_put(response->map_header, "Content-Type", "application/json");

  json_t * json_received_payload = ulfius_get_json_body_request(request, NULL);


  //Change third parameter to encoded payload
  char *jwt_as_string = strconcat(3, "\"", json_string_value(json_received_payload), "\"");

  ulfius_set_string_body_response(response, 200, json_string_value(json_received_payload));
  json_decref(json_received_payload);
  return U_CALLBACK_CONTINUE;
  return U_CALLBACK_CONTINUE;
}



/**
 * main function
 */
int main(void) {
  struct _u_instance instance;

  // Initialize instance with the port number
  if (ulfius_init_instance(&instance, PORT, NULL, NULL) != U_OK) {
    fprintf(stderr, "Error ulfius_init_instance, abort\n");
    return(1);
  }

  ulfius_add_endpoint_by_val(&instance, "OPTIONS", NULL, "*", 0, &callback_options, NULL);
  // Endpoint list declaration
  ulfius_add_endpoint_by_val(&instance, "POST", "/encode", NULL, 0, &callback_encode, NULL);
  ulfius_add_endpoint_by_val(&instance, "POST", "/decode", NULL, 0, &callback_decode, NULL);

  // Start the framework
  if (ulfius_start_framework(&instance) == U_OK) {
    printf("Started server on port %d\n", instance.port);
    getchar();
  } else {
    fprintf(stderr, "Error starting framework\n");
  }
  printf("End framework\n");

  ulfius_stop_framework(&instance);
  ulfius_clean_instance(&instance);

  return 0;
}