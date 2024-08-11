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
    Not much, but little by little.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 6:

    Working on parsing to create username for username by IP and loading into map.
    Started working on responses and re writing how messages are sent back to client but got cut short as work called for support and had to drop this for work instead.
    Hopefully get more progress done tomorrow.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 7:

    Re-organised functions a bit, re-thinking how I have the base server and message response to make it easier to send replies rather than re-writing same codes.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 8:

    Learned how GO modules and packages work through experimentation.
    Great write up by Alex Edwards:
    https://www.alexedwards.net/blog/an-introduction-to-packages-imports-and-modules

    Learn about structs work and I think I understand them a lot more now.
    Created separate package from main just to learn
    Re-organised my functions to make it easier to reply to clients.
    I believe I'm doing the message handler wrong but ran out of time today to test my theory that it's going to echo all users rather than single user, but need to verify.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 9:

    Re-done front end a bit to get more functionality.
    Prompt is initial for username then after success it'll open modal for chat room.
    Usernames checks added and working for duplicates.
    Username updates only does global message to show user updated from this to this instead of everyone seeing user’s response as well.
    Currently working on the popover in chat if users scrolled up, want to make it that user can just click popover like on twitch or youtube and it scrolls to bottom.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 10:

    Got the popover working as I wanted.
    Worked a bit on readability when user receives its own message.
    Hiding original username input after modal shows, will be used to hide to show channel list instead after.
    For channels list I was thinking the following:
        Create map like user connections. 
        Map will hold channel name then the head pointer to linked list.
        Linked list will hold all the users connected to the channel along with their permissions for the channel.
    If connecting to new channel will clear modal and make it fresh, is closing and re-opening the same channel modal won't be cleared so will hold the history and should hold all new messages that came in when modal was closed.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 11:

    Created doubly linked list for the channels. (First time creating a linked list and was easier than I thought)
    Need to map the linked list to the channels next and will be used for active users in channels, which users to broadcast to and check privileges for channel commands.
    
-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 12:

    Started working on linking the channels with linked list for users.
    Didn't have much time today to do coding.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 13:

    Fixed nil pointer dereference from Day 12 as I didn't test that code since I was so rushed.
    Trying to get the channels to work, might read some articles of how people are handling this as I might be going about it the wrong way.


-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 14:

    Got channels working, now part of server struct.
    Redone the reply so can be to individual and then new ReplyAll to reply to all connections.
    Fixed messages to be taken in as strings instead of array of chars 
    Same thing in principal but it's easier to put the message in string() at the message handler instead of writing out []byte(message) everytime for Reply() and ReplyAll().
    Need to fix GetChannels(), need to create array, push to array then return the array instead of the loop I have now since it's only sending single message.
    I thought it would return like a recursive function but obviously I haven't called or made it like a correct recursive function.
    Main function is now working, after fixing GetChannels() I will need to create front end to take the channels display them, on select it will join to that channel.
    Once that's working, it'll be foundation to get everything else in place as will just be different methods attached to the channels and some message parsing which I might fo a switch statement for instead of if statements to make it more readable. 
    Need to look into switch statements in GoLang if that's the case.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 15:

    Fixed GetChannels().
    Updated server side message parsing to do prefix instead of contains as works can be used in regular messages.
    Front end channels now being added as buttons.
    Buttons send message with channel name.
    Need to now change to only send messages to users in the channels selected.
    Parse for user privileges to join channel.
    Need to add functions for users to add channels.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 16:

    Core functionallity is working ! :D
    Hashed passwords for channels stored and working for connecting
    Channels now working correctly.
    Each channel replies only to it's own channels.
    Something broke for error receiving on front end need to double check why it's not displaying the errors anymore.
    Need to remove users from channels list after they leave, as it'll still be there right now even if user disconnects.
    Need to make modal close be channels list instead of username login as wont allow use to use /Username to update to the same since name already exists.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -

Day 17:

    Last day:
        Updated ui added sidebar.
        Pulls active users in chat.
        Got domain.
        Going to cross-compile to raspberry pie to host.
        Will port forward pie.
    
    Remarks:
        Most features other than core features not working.
        Ran out of time.
        Average coding time 2 hours per day.

-   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -   -