# Quora API

This is an unofficial Quora API to get information about Quora.

## What is working right now

At this moment, only profile user information can be get from this public API.

## How works

URL to make the request : https://quora-api-node.herokuapp.com

### What should I send ?

The only information needed is the url's name from the profile whose information you want to retrieve. Send it through headers, as 'user'.

![Profile URL example](/img/profile_url.png)
![Postman header sending example](/img/postman_header_ex.PNG)

### End Points Available

* (GET) `/user` -> Returns user profile info.
![GET /user received info example](/img/send_example.PNG)




