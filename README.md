# Tech Stack
<ul>
<li>Golang</li>
<li>gorm</li>
</ul>
 
## Instructions

### Overview

Cryptocurrencies are on the rise, and we want to get in on the action. Let's build a bot that watches the prices of certain coins and their prices, and places trades for us when they hit certain levels. 


### Core

#### Top Coin Service -> get and store top 3 coins in our db
This module collect and store the top 3 coins in the market in our DB
![Alt text](/git_images/p1.png "test locally" )
#### Trade coin Service -> Analize and place order is the requirements are meet for a given coing
This module is historical data for a given coin, and if the current price is lowest that market avg price; it Submit an error 
![Alt text](/git_images/p2.png "test locally" )
 

 
### APP features
- GET and store the top 3 cryptocurrency coins 
- COMPARE crypto
- PLACE order if crypto price is lower than the average price
- STORE order info in our db

#### Logs 
![Alt text](/git_images/p3.png "test locally" )

### Extra |  Crypto bot build on Python 
- https://github.com/jbenavidez/crypto-bot
 