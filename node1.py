from flask import Flask

app = Flask(__name__)


@app.route('/')
def hello():
    return 'This is Node1!'


if __name__ == '__main__':
    app.run(port=5001)