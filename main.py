from openai import OpenAI
import os

from mentis.textual_ui import app


def get_models():
    with open("env/mentis.env", "r") as f:
        api_key = f.readline()

    client = OpenAI(
        api_key=api_key,
        base_url="https://dashscope.aliyuncs.com/compatible-mode/v1",
    )
    models = client.models.list()

    for model in models.data:
        print(model.model_dump_json())


if __name__ == "__main__":
    # get_models()
    app.run()
