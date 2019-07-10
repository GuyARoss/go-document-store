# In Memory File Store
This in memory file store provies fast retrieval times for files by persisting them in memory.

## Installation
For running this project locally run the `run.sh` script within the root directory 

## Usage
### File Creation
To create a file, make a POST request to the default route `/` with an id parameter of the file name. 

__E.g. of JSON file__:
```curl
curl -X POST -H 'Content-Type: application/json' -d '{
    "data": "some data",
    "type": "this is a test",
}' http://localhost:4203?id=test
```

### File retrieval
To retrieve a file, make a GET request to the default route `/` with the id paramater of the file name.

__E.g. of file retrieval__: 
```curl
curl http://locahost:4203?id=test
```