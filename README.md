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
Lots of frustration and confusion today with trying to learn a new language and trying something you haven't done before as well.
Learning and trying to understand the type system is completely new to me so trying to get the hang of it, very different from the javascript I'm more used to writing.

I read through below article to get better understanding on websockets in general as reading docs proved difficult.
https://yalantis.com/blog/how-to-build-websockets-in-go/

Based on other articles and even the golang developer Aleksandr Ryzhyi in the above article recommends using the gorilla websocket instead of the std library one, so will be going with below websocket library instead.
https://pkg.go.dev/github.com/gorilla/websocket#section-readme

Using this tutorial to try and create the websocket. (I tried reading the docs but I can't make heads or tails of it right now with not being able to completely understand how a type system works yet)
https://reintech.io/blog/building-websocket-servers-in-go-using-gorilla-websocket

I keep getting error that incorrect websocket connection is being made even though I wrote out the code the same as in the tutorial.
I believe the client side websocket is not being created correctly or with correct data, possibly need to create a http server first to server the html as I wonder if the browser doesn't have access yet to the index.html?

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 3:

Hello World!

Everyone's first program, so only seems right.
Got the server and client to actually write and respond to each other.

My hunch yesterday was right about the server not having access to the index, I wasn't serving it correctly.
I believe I understand how and why routes are made now.
Was messing about with route handlers to get an idea and to confirm some pre-notions.

After that, I was playing about with the tcp packet to see what I can see on the packet and to see what a tcp packet data layer looks like.

Articles I was reading today to get ideas and understanding of what I'm actually doing with these libraries:
https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
https://blog.logrocket.com/routing-go-gorilla-mux/
https://www.alexedwards.net/blog/interfaces-explained

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 4:

Threw together a quick front-end to be able to send and receive message from webpage instead of a forced message on open.
Need to do server side parsing, malicious stuff and just empty {enter} keystrokes before returning.
Need to assign users names as well as a timestamp with message most likely to be able to keep track of chat times and who sent the message.
I think I might tackle the first instance of entering a username and keeping track of it on the connection otherwise will have to be re-write parser and stamp anyways.

For the website design I was thinking of opening a modal view for the chat itself after it's selected from list, and possibly keeping memory of chat client side if accidental close while still connected to server.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 5:

Didn't get much done, busy with life.
Read a bit about maps and realised I'd probably need locks on the maps.
Read the about mutex's, read/write locks and unlocks from the sync package.
Started building simple parser for setting username as first instance, will have message handler for all the different commands.
Not much, but  little by little.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 6:

Working on parsing to create username for username by IP and loading into map.
Started working on responses and re writing how messages are sent back to client but got cut short as work called for support and had to drop this for work instead.
Hopefully get more progress done tomorrow.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 7:

Re-organised functions a bit, re-thinking how I have the base server and message response to make it more easier to send replies rather than re-writting same codes.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -