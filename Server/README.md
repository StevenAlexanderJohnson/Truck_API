# The API

As explained earlier, the API creates a safe and streamlined way of interacting with the database. In your case, you should never expose this 
API. What I mean by that is you should never share the URL for this API. If you do then people with malintent could try to find a way into 
the API. I am no security expert, but I did implement authentication.


## Authentication

The API uses cookie authentication. How this works is that the user signs in using a username and password, and is then assigned a JWT 
(JavaScript Web Token). The JWT is then returned to the user in the form of a Cookie (small bit of data that is stored in the user's browser). 
Every time the user uses the website and sends a request to the API, the cookie is sent along the request which the API uses verify that the 
user is who they say they are.

It is important to note that the JWT is not encrypted but it is **SIGNED**. What that means is that anybody can take this cookie and read its 
contents, but any attempt to change the cookie's value will result in the cookie being rejected and signing out the user. (This would be a great 
place to log the event) This is slightly technical so hit me up if you want to know further. 
Seeing that the JWT is not encrypted, you should **NEVER** keep anything sensative in the cookie, i.e. passwords or secrets.


## The Code

This code is written in Go, a C like language that is memory safe (don't worry about it). It compiles to binary which means that it can run on 
any machine without having to have the language installed (unlike Python, C#, or JavaScript). This is useful because most servers are run on 
Linux but gives you the flexability to use Windows on your server if you want to. 
It is very easy to learn thanks to not having to manage memory (like you had to do in C/C++ class in college), and has very basic concepts. 
[https://go.dev](https://go.dev) has a workbook that you can work through that explains how to code in this language and should have you up 
to speed once you complete. You probably will never have to learn how to code in Go but if you want to there it is. 

This code uses a few external libraries which is why most of the server code looks like gibberish. I don't recommend trying to add endpoints 
yourself and have your techies do it. It also helps to have Github Copilot. It is an AI that was trained using hundreds of thousands of 
repositories like this one and uses that knowledge to write code for you. It's very useful but also expensive.