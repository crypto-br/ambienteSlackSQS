from pyexpat.errors import messages
import urllib.parse
import boto3
import json
import time
import requests
from flask import Flask

slack_token = 'YourSlackToken'
slack_icon_emoji = ':see_no_evil:'
slack_user_name = 'YourUserName'
queue_name = 'alert'
aws_access_key_id="yourKey"
aws_secret_access_key="yourSecret"

HOST = "http://localhost.localstack.cloud:4566"

def post_message_to_slack(text, channel, blocks = None):
    return requests.post('https://slack.com/api/chat.postMessage', {
        'token': slack_token,
        'channel': channel,
        'text': text,
        'icon_emoji': slack_icon_emoji,
        'username': slack_user_name,
        'blocks': json.dumps(blocks) if blocks else None
    }).json()

app = Flask(__name__)
@app.route('/recivesqs', methods=["GET"])
def receive_message():
    sqs_client = boto3.client("sqs", endpoint_url=HOST, region_name="us-east-1", aws_access_key_id=aws_access_key_id, aws_secret_access_key=aws_secret_access_key)
    response = sqs_client.receive_message(
        QueueUrl=f"http://localhost.localstack.cloud:4566/000000000000/{queue_name}",
        MaxNumberOfMessages=1,
        WaitTimeSeconds=10,
    )
    for message in response.get("Messages", []):
        message_body = message["Body"]
        out = json.loads(message_body)
        canal =  out["Channel"]
        msg = out["Msg"]
        post_message_to_slack(msg, canal)

    return "recive msg"

if __name__ == '__main__':
    app.debug = True
    app.run(host = '0.0.0.0',port=8082)


