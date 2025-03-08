## Running Ollama Third-Party Service 

### Choosing a Model 

You can get the model_id that ollama will launch from the [Ollama Library](https://ollama.com/library).

eg. https://ollama.com/library/llama3.2

### Getting the Host IP

#### Linux

Get your IP address
```sh
sudo apt install net-tools
ifconfig
```
Or you can try this way `$(hostname -I | awk '{print $1}')`

HOST_IP==$(hostname -I | awk '{print $1}') NO_PROXY=localhost LLM_ENDPOINT_PORT=9000 LLM_MODEL_ID="llama3.2:1B" docker compose up

### Ollama API

Once the Ollama server is running we can make API calls to the Ollama API

https://github.com/ollama/ollama/blob/main/docs/api.md 

## Download (Pull) a model

curl http://localhost:9000/api/pull -d '{
  "model": "llama3.2:1B"
}'

## Generate a Request

curl http://localhost:9000/api/generate -d '{
  "model": "llama3.2:1B",
  "prompt": "Why is the sky blue?"
}'

# Technical Uncertainty 

Q: Does bridge mode mean we can only access Ollama API with another model in the docker compose?

A: No, the host machine will be able to access it

Q: Which port is being mapped 8008->141414

A: In this case, 8008 is the port the host machine will access. The other port is the guest port (the port of the service inside the container)

Q: If we pass the LLM_MODEL_Id to the ollama server will it download the model when on start?

A: It does not appear so. The ollama CLI might be running multiple APIs so you need to call/pull api before trying to generate text

Q: Will the model be downloaded in the container? Does that mean the LLM model will be deleted when the container stops running?

A: The model will download into the container, and vanish when the container stops running. you need to mount a local drive and there is probably more work to be done.

Q: For LLM service which can text-generate, it suggests it will only work with TGI/vLLM, and all you have to do is have it running. Do TGI and vLLM have a standardized API, or is there code to detect which one is running? Do we have to really use Xeon or Gaudi processors?

vLLM, TGI (Text Generation Inference), and Ollama all offer APIs with OpenAI compatibility, so in theory they should be interchangeable.