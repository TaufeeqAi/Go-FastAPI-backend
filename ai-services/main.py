from fastapi import FastAPI
from pydantic import BaseModel
from langchain.schema import HumanMessage
from dotenv import load_dotenv
from ai_factory.factory import create_llm

load_dotenv()

app = FastAPI()

# Initialize the LLM using the factory
llm = create_llm(provider="groq")

class AskRequest(BaseModel):
    prompt: str

@app.post("/ai/ask")
def ask(request: AskRequest):
    message = HumanMessage(content=request.prompt)
    response = llm.invoke([message])
    return {"response": response.content}
