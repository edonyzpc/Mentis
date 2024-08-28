from setuptools import setup, find_packages

setup(
    name="Mentis",
    version="0.0.1",
    packages=find_packages(exclude=["tests"]),
    package_data={"": ["*.tcss"]},
    include_package_data=True,
    description="A LLM-based app for attack surface management.",
    author="edonyzpc",
    author_email="edonyzpc@edony.ink",
    url="https://github.com/edonyzpc/Mentis",
    classifiers=[
        "Development Status :: 3 - Alpha",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
        "Programming Language :: Python :: 3.12",
        "License :: OSI Approved :: MIT License",
        "Topic :: Scientific/Engineering :: Artificial Intelligence",
    ],
    entry_points={
        "console_scripts": [
            "mentis = mentis.textual_ui.app:run",
        ],
        "mentis_function": [
            "google_web_search = mentis.runtime.function_calling.functions:google_web_search"
        ],
    },
    keywords="attack-surface,gpt,chat,llm,chatgpt,langchain,openai,sap,textual,terminal",
    install_requires=[
        "httpx",
        "humanize",
        "langchain",
        "langchain-core",
        "polars",
        "pydantic",
        "pyperclip",
        "rich",
        "shortuuid",
        "textual",
        "tiktoken",
        "toolong",
        "pyyaml",
        "setuptools",
    ],
    extras_require={
        "openai": ["langchain-openai"],
        "google": ["langchain-google-genai"],
        "sap": ["generative-ai-hub-sdk"],
        "anthropic": ["langchain-anthropic"],
        "qwen": [
            "langchain-community",
            "dashscope",
        ],
        "all": [
            "langchain-openai",
            "langchain-google-genai",
            "generative-ai-hub-sdk",
            "langchain-anthropic",
            "langchain-community",
            "dashscope",
        ],
    },
)
