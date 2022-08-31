1. How the computer and other devise communicate one and another via internet.

- computer making message to another devise responding based on your message.
- sending and receiving message.

2. How the it happen ?

- computer and other devises standardize with set of protocols.
- it defines how should be communicate one and another.

3. One of these protocol is TCP / IP.

- TCP Transmission control protocol.
- IP Internet protocol
- These are defining how the computer and other devise communicate one and another.

4. IP address ?

- IP is the address for every devise on the internet.

5. What's port ?

- Assume IP as home port as particular person on the home.

6. What's HTTP ?

- HTTP hyper text transfer protocol.
- Envelope out side we have sender address and receiver address which means IP.
- but inside we have http request which contains.
- all the request parameter are bundled by http.
- example :

```
GET / HTTP/1.1
HOST: www.google.com / name="justins"
....
    > get request method it represent web page.
    > http/1.1 represent the http version.
    > host represent first parameter and using this "/ " you can use 2'parameter.
HTTP/1.1  200  ok
Content-Type: text/html
....
    > http/1.1 represent the http version.
    > state code 200 ok
    > state code 404 not exist.
```

6. status code

- 404 not found
- 403 access to the requested resource is forbidden. The server understood the request, but will not fulfill it.
- 500 Internal Server Error.
- 200 successful
