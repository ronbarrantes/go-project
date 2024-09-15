    // there is a ton of thing that have to happen, I think
    // We are going to create an api that when someone hits an endpoint with a qr code
    // it will increase a counter in the QR database
    // inside the database we have users and the amount of times they have scanned the qr code
    // the qr code will contain information about the store and the endpoint
    // when the qr code is read, it will hit an endpoint and pass the user's id as a parameter
    // the endpoint will then increase the counter in the database
    // for the time being the user's id will be created at random in the frontend and stored in local storage
    // the frontend will check if the user has an id, if not it will create one and store it in local storage
    // the qr code will represent a coffee shop and the endpoint will be the coffee shop's endpoint
    // there should be a table for users
    // there should be a table for coffee shops
    // there should be a table to relate users to coffee shops and how many times they have scanned the qr code
    // a coffee shop should be able to update their id without affecting the user's points

    //// ------ from here on down more business and less technical stuff
    // a user should be able to scan the qr code multiple times from multiple restaurants
    // a qr code can be made on the fly by a coffee shop
    // a coffee shop should be able to see users and their points
    // a coffee shop should be able to reset a user's points
    // a coffee shop should be able to add points to a user
    // a user should be able to see their own points
    // after n number of points (specified by the coffee shop) a user should be able to redeem a reward (a coffee for example)
    // a user should be able to see their rewards
    // a user should be able to see when they have redeemed a reward
    // a store can manage itself
