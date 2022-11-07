# transactionAPIs
A RESTful API for transactions statistics. Its main use case is to calculate realtime statistics for the last 60 seconds of transactions.

## APIs
### 1. POST /transactions
Called every time a transaction is made. It is also the sole input of this rest API.
#### Parameters
Passed in the request body (JSON)
* `amount`-Transaction amount;.
* `timestamp` -transaction time in the ISO 8601 format YYYY-MM-
DDThh:mm:ss.sssZ in the UTC timezone (this is not the current timestamp)

#### Example request body
```
{
    "amount": 12.3343,
    "timestamp": "2018-07-17T09:59:51.312Z"
}
```
#### Response:
* 201 – In case of success
* 204 – If the transaction is older than 60 seconds
* 400 – If the JSON is invalid
* 422 – If any of the fields are not parsable or the transaction date is in the future

### 2. GET /statistics
#### Sample Response:
```
{
    "sum": "1000.00",
    "avg": "100.53",
    "max": "200000.49",
    "min": "50.23",
    "count": 10
}
```
#### Return parameters:
* sum – total sum of transaction value in the last 60 seconds
* avg – average amount of transaction value in the last 60 seconds
* max –single highest transaction value in the last 60 seconds
* min – single lowest transaction value in the last 60 seconds
* count – total number of transactions that happened in the last 60 seconds

### 3. DELETE /transactions
* This endpoint causes all existing transactions to be deleted
* The endpoint should accept an empty request body and return a 204 status code.




