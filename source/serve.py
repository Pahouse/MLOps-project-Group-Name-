"""Production entrypoint: start the MLflow scoring server and the FastAPI proxy, and keep running."""
import uvicorn
from fastapi import FastAPI

import deployment as d  # importing is safe: the pipeline is under __main__

# The handlers read this global to report health
d.scoring_proc = d._start_scoring_server()

app = FastAPI(docs_url=None, openapi_url=None, redoc_url=None)
app.get("/health")(d.health)
app.get("/version")(d.version_endpoint)
app.post("/invocations")(d.invocations)
app.post("/infer")(d.invocations)
app.get("/drift")(d.drift)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=d.port, log_level="info")