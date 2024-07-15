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