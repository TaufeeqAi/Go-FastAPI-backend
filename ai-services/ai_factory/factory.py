import os
from langchain_groq import ChatGroq
# We'll add other providers here later

def create_llm(provider: str = "groq", api_key: str = None):
    if provider == "groq":
        if not api_key:
            api_key = os.getenv("GROQ_API_KEY")
        return ChatGroq(model="llama3-70b-8192", api_key=api_key)
    # Add other providers here, e.g., 'gemini', 'anthropic'
    else:
        raise ValueError(f"Unsupported LLM provider: {provider}")