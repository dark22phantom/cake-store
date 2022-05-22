# cake-store

List of API:

1. Insert cake.
method: POST.
endpoint: /cakes.
request:
```JSON
{
    "title":"Lemon cheesecake",
    "description":"A cheesecake made of lemon",
    "image":"https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
    "rating":7
}
```JSON
response:
```JSON
{
    "responseCode": "200",
    "responseMessage": "Success",
    "data": {
        "title": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 7,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
    }
}
```JSON

2. List of cake.
method: GET. 
endpoint: /cakes.
request: -.
response:
```JSON
{
    "responseCode": "200",
    "responseMessage": "Success",
    "data": [
        {
            "id": 1,
            "string": "Lemon cheesecake",
            "description": "A cheesecake made of lemon",
            "rating": 7,
            "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
            "created_at": "2022-05-22 18:51:25",
            "updated_at": "2022-05-22 18:51:25"
        }
    ]
}
```JSON

3. Detail of cake.
method: GET. 
endpoint: /cakes/:id.
request: -.
response:
```JSON
{
    "responseCode": "200",
    "responseMessage": "Success",
    "data": {
        "id": 1,
        "string": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 7,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
        "created_at": "2022-05-22 18:51:25",
        "updated_at": "2022-05-22 18:51:25"
    }
}
```JSON

4. Update cake.
method: PATCH. 
endpoint: /cakes/:id.
request:
```JSON
{
    "title":"Lemon cheesecake 2",
    "description":"A cheesecake made of lemon",
    "image":"https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
    "rating":7
}
```JSON
response:
```JSON
{
    "responseCode": "200",
    "responseMessage": "Success",
    "data": {
        "id": 1,
        "string": "Lemon cheesecake 2",
        "description": "A cheesecake made of lemon",
        "rating": 7,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
        "created_at": "2022-05-22 18:51:25",
        "updated_at": "2022-05-22 18:54:32"
    }
}
```JSON

5. Delete cake.
method: DELETE. 
endpoint: /cakes/:id.
request: -.
response:
```JSON
{
    "responseCode": "200",
    "responseMessage": "Success",
    "data": ""
}
```JSON

