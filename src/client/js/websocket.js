const socket = new WebSocket('ws:' + window.location.href.replace("http://","") + 'ws');
let user;
let channels;
let previous;
let jchannel;
const channel = document.getElementById("channels")
socket.addEventListener('message', (event) => {messageHandler(event.data)});

function messageHandler(message){
    previous = message;
    console.log(message)
    if (message.includes("Error")){
        errorHandler(message)
    } else if (message.includes("Username set as:")){
        user = message.slice(17)
        socket.send("/Request channels")
    } else if (message.startsWith("Channels:")){
        clearChannels();
        channels = message.slice(10)
        channels = channels.split(",")
        hideBlock(entry, "hide");
        insertChannels(channels);
        hideBlock(channel, "show");
    } else if (message.startsWith("Enter password")) {
        //let join = message.slice(14)
        hideBlock(channel, "hide")
        receiveMessage(message)
        modal.show();
        
    } else if (message.startsWith("Joined")) {
        jchannel = message.slice(7)
        hideBlock(channel, "hide")
        receiveMessage(message)
        document.getElementById("channelName").innerText = jchannel;
        modal.show();
    } else if (message.startsWith("Users:")){
        let users = message.slice(6);
        users = users.split(",");
        for(let i = 0; i < users.length; i++){
            insert("button", 
                {"data-bs-toggle":"offcanvas", "data-bs-target":"#chatSidebar","value": users[i],"onclick":"insertOption('"+  users[i] + "')"},
                "btn btn-toggle d-inline-flex align-items-center rounded",
                "Users:" + users[i],
                false
            )
        }
    } else {
        receiveMessage(message)
    }
}

function errorHandler(message){
    console.log(entryError.childNodes[1])
    if (message.startsWith("Username Error:")){
        entryError.childNodes[1].innerText = message;
        entryError.style.display = "block";
    } else {
        console.log("Error message: " + message);
        entryError.childNodes[1].innerText = message;
        entryError.style.display = "block";
    }
}

function isUser(message){
    return message.slice(0, message.indexOf(":")) === user
}

function receiveMessage(message){
    message = message.replace("/Channel:" + jchannel + ":", "")
    if((chat.scrollTop + chat.clientHeight) === chat.scrollHeight){
        insert("p",{},"",message, isUser(message))
        chat.scrollTop = chat.scrollHeight  
    } else {
        insert("p",{},"",message, isUser(message))
        popupTrigger();
    }
}

function getUsers(){
    for(let i = ulUserList.childNodes.length -1; i > 0; i--){
        console.log(ulUserList.childNodes[i])
        ulUserList.childNodes[i].remove()
    }
    socket.send("/Users:Channel:" + jchannel)
}