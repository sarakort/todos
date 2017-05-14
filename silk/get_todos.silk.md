# Users

## GET /todos

* Content-Type: "application/json; charset=utf-8"
* Accept: "application/json"

===
### Expected response
* Status: `200`
* Content-Type: "application/json; charset=utf-8"
```json
[
  {
    "ID": 1,
    "Checked": false,
    "Message": "Homework 1"
  },
  {
    "ID": 2,
    "Checked": false,
    "Message": "Homework 2"
  }
]
```