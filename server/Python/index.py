from flask import Flask
from flask_restful import Resource, Api, reqparse
from flask_cors import CORS
import jwt

app = Flask(__name__)
# Allow all access
CORS(app)
api = Api(app)

class Encode(Resource):
    def post(self):
        try:
            key = 'secret'
            received_object = reqparse.request.get_json()
            encoded_object = jwt.encode(received_object, key, algorithm='HS256').decode("utf-8")
            return encoded_object

        except Exception as e:
            print("Error on /encode: ", e)
            return 'An error occurred'
    pass

class Decode(Resource):
    def post(self):
        try:
            received_jwt = reqparse.request.get_json()
            # Not verifying the signature
            decoded_jwt = jwt.decode(received_jwt, verify=False)
            return str(decoded_jwt)

        except Exception as e:
            print("Error on /decode: ", e)
            return 'Invalid Token'
    pass

api.add_resource(Encode, '/encode') 
api.add_resource(Decode, '/decode') 

if __name__ == '__main__':
    app.run(host='127.0.0.1', port=8080)