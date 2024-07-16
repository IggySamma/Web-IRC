My plan is to try build an irc chat like from back in the day.

This is a learning project for me to learn GO

My goals:

Make Web socket connection from scratch for channel to end users

Channel creators have elevated permissions to include:
Kick
Ban (by ip and Mac address)
Silence for time period
Shadow ban (ban without notification and chats blank always to end user banned)
Give other users customised privileges

Chats to be stored in memory so no history of chats

Possibly whisper functionality

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Hackathon DevLog:

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 1:

    Installed Go on my laptop and installed GO Lsp for VSCode
    Watched youtube tutorial for starting in Go that was linked in docs:
    https://www.youtube.com/watch?v=Q0sKAMal4WQ

    Started reading docs for how to make websocket connection.
    Will probably go with following package for websockets from Go themselves.
    https://pkg.go.dev/golang.org/x/net/websocket

    Current thought process for what I'm planning to build.

    Pseudo code/process flow for planned functionality:
        Client: User joins page.
        Client: Enters a username for chat.
        Server: Create websocket connection with client.
        Server: Websockets map of clients usernames and ip/mac address connected (for banning from channels)
        Client: Gets list of avaliable channels or to create new channel.
        Client: Can send message to channel.
        Server: Broadcasts message back to channel and everyone subscribed to it.
        
            Joining channel:
        Server: Check clients not banned.
        Server: Password check for channel, and time-out for join attempts after 5 failed attempts of password.
        Server: Subscribe user to channel.
        
            Creating server:
        Server: Create new channel 
        Server: Assign creator to admin priveledges for that channel only.
        Server: Subscribe the user to the channel.
        Client: Let users create password protected channels

        User features:
        Whisper: Message other users privatly not through a channel.
        Block: Block users from messaging you.

        Channel admin features:
        Ban: Ban user from channel.
        Kick: Kicks user from channel.
        ShadowBan: Bans user from messaging but can still view chat.
        Create role: Admins to be able to create custom roles for channel with different level of priveledges.
        Assign role: Assign created roles to user in chat.

        Client features:
        Total online users counter.
        Channel user count & user count.
        Chat for whispers.


-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 2:

Didn't have a lot of time today, but got a solid hour/2 in on reading about this.
Lots of fustration and confusion today with trying to learn a new language and trying something you haven't done before as well.
Learning and trying to understand the type system is completely new to me so trying to get the hang of it, very different from the javascript I'm more used to writing.

I read through below article to get better understanding on websockets in general as reading docs proved difficult.
https://yalantis.com/blog/how-to-build-websockets-in-go/

Based on other articles and even the golang developer Aleksandr Ryzhyi in the above article recommends using the gorrila websocket instead of the std library one, so will be going with below websocket library instead.
https://pkg.go.dev/github.com/gorilla/websocket#section-readme

Using this tutorial to try and create the websocket. (I tried reading the docs but I can't make heads or tails of it right now with not being able to completely understand how a type system works yet)
https://reintech.io/blog/building-websocket-servers-in-go-using-gorilla-websocket

I keep getting error that incorrect websocket connection is being made even though I wrote out the code the same as in the tutorial.
I believe the client side websocket is not being created correctly or with correct data, possibly need to create a http server first to server the html as I wonder if the browser doesn't have access yet to the index.html?

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -