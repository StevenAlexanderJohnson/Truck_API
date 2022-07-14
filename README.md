# Hello Maria, Welcome to the QR Code Demo.

In this repo you will find two folders, the first being that of which holds the code for the website and the other holding the API.

These two solutions run seperately but interact together to form a complete product.

## Let's run down what each component does.

### The Website

The first component to cover is the website. The website is what the user will use to interact with the API. It is purely for looks and 
has absolutely no direct control of the database. The purpose of the website is to create a user friendly experience. They will not need 
to know SQL or have any technical knowledge to operate. That being said the website **DOES** interact with the database indirectly, 
thanks to the API.


## The API

The API, *Application Programming Interface*, is a protective layer between the user and the database. It's job is to collect data from the database and
return it in a streamlined way. The user will not need to know SQL to interact with it, but without the website they will need to have some technical knowledge 
to be able to use it.Unlike the website the database **DOES** interact with the database directly. That is why it is important to have the user authenticate 
in order to useit. By authenticating you can add logging to the API to see who changed what when. That is not in the demo but it is possible and in production 
recommended.The API also implements ROLE authentication. What that means is that you can assign users roles, i.e. *Engineer*, *TechSupport*, *Developer*, and 
the lockcertain functionality from people without the correct roles. This is useful if you have an endpoint, *url within the API*, that allows the user to modify
data. You do not want the average Joe to change data, only managers, engineers, and auditers.

The API has predetermined operations that it will execute upon request. That means that the SQL is already written and upon request it will execute that 
SQL and return the results. This is nice because it takes away the need to know SQL and makes the changes and queries to the database consistent.


## The Database

You know what a database is. It is essentially a massive file containing data. The user will *NEVER* directly interact with the database. The purpose of 
having the website and the API is to allow a user to interact with the database in a safe and streamlined way. If done correctly there is no chance that 
the user will ever accidently delete or alter data that they did not mean to delete/alter.


## How the Demo Works

How the demo works is that the user should be able to scan a QR Code, which contains a url to the website. When the user clicks on the url or navigates to the
website. If the user is not logged in yet, they will be directed to the login page and if not they will go straight to the page in the QR Code. Once there the page will
collect data from the API and will then display it along side the QR Code that they used to get there.
If you are using the website without the QR Code, you can use the website to get the QR Code which you can download by clicking it.


## That's about it.

For more specifics on how to run, go into each file and read the README.