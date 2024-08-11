const modal = new bootstrap.Modal(document.getElementById("modalView"), {
    backdrop: "static",
    focus: true,
    keyboard: false
});
const popupElement = document.getElementById("message")

let popup = new bootstrap.Popover(popupElement, {
    container: '.modal-body',
    content: 'Scroll down to see newest messages',
    trigger: 'manual',
    delay: { "show": 500, "hide": 500},
});

const setUser = document.forms["entryUsername"];
const chatForm = document.forms["liveChat"];
const chat = document.getElementById("chat");
const entry = document.getElementById("entry");
const entryError = document.getElementById("error");
const getModal = document.getElementById("modalView");
const channelForm = document.getElementById("channelsForm");
const createChannel = document.forms["createNewChannel"]
let optionsButtons = document.querySelectorAll("#optionsBtn");
let chatOptions;
const userList = new bootstrap.Collapse('#chatUserList', {
    toggle: false
})
const ulUserList = document.getElementById("ulUserList")

chat.addEventListener("scroll", () => { popupTrigger() } );

function insert(type = "div", attributes = {}, classes = "message", message, isUser){
    const element = document.createElement(type);
    if (classes) element.classList = classes
    for (let key in attributes){
        element.setAttribute(key, attributes[key])
    };

    if (message.startsWith("Users:")) {
        let users = message.slice(6);
        element.innerText = users;
        let li = document.createElement('li')
        li.appendChild(element);
        ulUserList.appendChild(li)
        return
    } else if(isUser) {
        element.innerText = "---> " + message;
    } else {
        element.innerText = message;
    };

    if(message.startsWith("Channel:")){
        element.innerText = message.slice(9);
        channelForm.appendChild(element);
    } else if (message){
        let chat = document.getElementById("chat");
        chat.appendChild(element);
    };
    return
};

function popupTrigger(){
    let isShown = document.querySelector("[id^='popover']")
    if((chat.scrollTop + chat.clientHeight) === chat.scrollHeight && isShown !== null ){
        popup.hide();
    } else if((chat.scrollTop + chat.clientHeight) !== chat.scrollHeight && isShown === null){
        popup.show();
    };
};

chatForm.addEventListener("submit", (event) => {
    event.preventDefault(); 
    let message = document.getElementById("message");
    console.log(message.value.trim() !== 0)
    if (message.value.trim().length !== 0) {
        if(previous.includes("Enter password")){
            socket.send("/Password:/Channel:" + message.value + ":" + jchannel)
        } else {
            socket.send("/Channel:" + jchannel + ":" + message.value);
        }
    }
    message.value = "";
});

setUser.addEventListener("submit", (event) => {
    event.preventDefault();
    let username = document.getElementById("username");
    socket.send("/Username: " + username.value);
    user = username.value
    username.value = "";
});

function hideBlock(id, display){
    if (display === 'show'){
        id.style.display = 'block';
    } else {
        id.style.display = 'none'
    };
};

getModal.addEventListener('shown.bs.modal', () => {
    if(entryError.style.display !== 'none'){
        entryError.style.display = "none";
    };

    if (entry.style.display !== 'none'){
        entry.style.display = 'none';
    };
});

getModal.addEventListener('hide.bs.modal', () => {   
    /*if (entry.style.display === 'none'){
        entry.style.display = 'block';
    };*/
    hideBlock(channel, "show");
    socket.send("/Disconnect:" + jchannel)
    clearChat();
});

window.addEventListener("onbeforeunload", () => {
    socket.send("/Disconnect:" + jchannel)
})

function insertChannels(channels){
    for(let i = 0; i < channels.length; i++){
        insert("button", {"id": "channelButton"} ,"btn btn-primary m-2", "Channel: " + channels[i], false);
    }
    let buttons = document.querySelectorAll('#channelButton');
    for(let i = 0; i < buttons.length; i++){
        buttons[i].addEventListener("click", (event) => {
            event.preventDefault();
            jchannel = event.srcElement.innerText;
            socket.send("/Join: " + event.srcElement.innerText);
        })
    };
};


function clearChannels(){
    let buttons = document.querySelectorAll('#channelButton');
    buttons.forEach(button => {
        button.remove();
    });
};

function clearChat(){
    let chat = document.querySelectorAll('p')
    chat.forEach(line => {
        if(line.id != "errorMessage"){
            line.remove();
        }
    });
};

function insertOption(option){
    console.log(option);
    if (option.startsWith("/")){
        chatOptions = option;
        getUsers()
        userList.show();
    } else {
        chatOptions = chatOptions + "Channel:" + jchannel + ":" + option;
        socket.send(chatOptions);
        chatOptions = "";
    }
}

function toggleCreate(){
    hideBlock(channelForm,"hide")
    hideBlock(createChannel, "show")
}

createChannel.addEventListener("submit", (event) => {
    event.preventDefault();
    let newChannelName = document.getElementById("newChannelName");
    let newChannelPass = document.getElementById("newChannelPassword");
    socket.send("/Create:" + newChannelName + ":" + newChannelPass);
    hideBlock(channelForm,"show")
    hideBlock(createChannel, "hide")
})