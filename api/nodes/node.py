from flask import Flask
import os

app = Flask(__name__)
name = os.getenv("NAME")
port = int(os.getenv("PORT"))

@app.route('/')
def hello():
    return f'This is {name}!'


if __name__ == '__main__':
    app.run(host='0.0.0.0',port=port)