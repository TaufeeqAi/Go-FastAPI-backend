from fastapi import FastAPI

app = FastAPI()

@app.get("/ai/hello")
def read_root():
    return {"message": "Hello from the AI service!"}