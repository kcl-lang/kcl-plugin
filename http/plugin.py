import requests


def get(url: str):
    resp = requests.get(url)
    return {
        "status": resp.status_code,
        "body": str(resp.content),
    }
