# Golang Quiz App
A simple API to create a quiz, built on top of `Go 1.18`.

### Endpoints
* `/quizzes` - returns a list of quizzes
  * Method: `GET`
  * Response:
    ```
    "data": [
        {
            "id": 1,
            "quiz": {
                "question": "Apa nama ibu kota Indonesia?",
                "answers": [
                    {
                        "key": "A",
                        "value": "Jakarta"
                    },
                    {
                        "key": "B",
                        "value": "Semarang"
                    }
                ]
            }
        }
    ]
    ```


* `/quiz` - returns a random quiz
  * Method: `GET`
  * Response:
    ```
    "data": {
        "id": 1,
        "quiz": {
            "question": "Apa nama ibu kota Indonesia?",
            "answers": [
                {
                    "key": "A",
                    "value": "Jakarta"
                },
                {
                    "key": "B",
                    "value": "Semarang"
                }
            ]
        }
    }
    ```


* `/answer/:quizID` - returns a result of a specific quiz answer
    * Method: `GET`
    * Query Parameters:
      * `key`: only receive an uppercase alphabet of length 1
    * Response:
      ```
      "data": {
          "result": true
      }
      ```


* `/quiz` - creates a new quiz
    * Method: `POST`
    * Request Body:
      ```
      {
          "question": "Testing",
          "answers": [
              {
                  "key": "A",
                  "value": "Hello",
                  "isCorrect": true
              },
              {
                  "key": "B",
                  "value": "World",
                  "isCorrect": false
              }
          ]
      }
      ```
    * Response:
      ```
      "data": {
          "message": "successfully create a new quiz"
      }
      ```


* `/quiz/:quizID` - removes a specific quiz
    * Method: `DELETE`
    * Response:
      ```
      "data": {
          "message": "successfully delete a quiz"
      }
      ```