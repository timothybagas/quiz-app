# Golang Quiz App
A simple API to create a quiz, built on top of `Go 1.18`.

### Endpoints
* `/quizzes` - returns a list of quizzes
  * Method: `GET`


* `/quiz` - returns a random quiz
  * Method: `GET`


* `/answer/:quizID` - returns a result of a specific quiz answer
    * Method: `GET`
    * Query Parameters:
      * `key`: only receive an uppercase alphabet of length 1


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


* `/quiz/:quizID` - removes a specific quiz
    * Method: `DELETE`