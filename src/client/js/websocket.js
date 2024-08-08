const socket = new WebSocket('ws:' + window.location.href.replace("http://","") + 'ws');
let user;
let channels;
const channel = document.getElementById("channels")
socket.addEventListener('message', (event) => {messageHandler(event.data)});

function messageHandler(message){
    console.log(message)
    if (message.includes("Error:")){
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
    } else if (message.startsWith("Join channel:")) {
        //let join = message.slice(14)
        hideBlock(channel, "hide")
        modal.show();
    }else {
        if((chat.scrollTop + chat.clientHeight) === chat.scrollHeight){
            insert("p",{},"",message, isUser(message))
            chat.scrollTop = chat.scrollHeight  
        } else {
            insert("p",{},"",message, isUser(message))
            popupTrigger();
        }
    }
}

function errorHandler(message){
    if (message.startsWith("Username Error:")){
        entryError.childNodes[1].innerText = message;
        entryError.style.display = "block";
    } else {
        console.log("Error message: " + message);
    }
}

function isUser(message){
    return message.slice(0, message.indexOf(":")) === user
}