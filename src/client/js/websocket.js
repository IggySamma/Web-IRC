const socket = new WebSocket('ws:' + window.location.href.replace("http://","") + 'ws');
const modal = new bootstrap.Modal(document.getElementById("modalView"), {
    backdrop: "static",
    focus: true,
    keyboard: false
})
const setUser = document.forms["entryUsername"];
const chatForm = document.forms["liveChat"];
const chat = document.getElementById("chat")
const entry = document.getElementById("entry")
const entryError = document.getElementById("error");
const getModal = document.getElementById("modalView")

let user;
const popupElement = document.getElementById("message")
let popup = new bootstrap.Popover(popupElement, {
    container: '.modal-body',
    content: 'Scroll down to see newest messages',
    trigger: 'manual',
    delay: { "show": 500, "hide": 500},
})


socket.addEventListener('message', (event) => {messageHandler(event.data)});

function insert(type = "div", attributes = {}, classes = "message", message, isUser){
    if(message){
        let chat = document.getElementById("chat");
        const element = document.createElement(type);
        if (classes) element.classList = classes
        for (let key in attributes){
            element.setAttribute(key, attributes[key])
        }
        if (isUser) {
            element.innerText = "---> " + message;
        } else {
            element.innerText = message;
        }
        chat.appendChild(element);
    }
    return
}

chat.addEventListener("scroll", () => { popupTrigger() } )

function popupTrigger(){
    let isShown = document.querySelector("[id^='popover']")
    if((chat.scrollTop + chat.clientHeight) === chat.scrollHeight && isShown !== null ){
        popup.hide();
    } else if((chat.scrollTop + chat.clientHeight) !== chat.scrollHeight && isShown === null){
        popup.show();
    } 
}

chatForm.addEventListener("submit", (event) => {
    event.preventDefault(); 
    let message = document.getElementById("message");
    socket.send(message.value);
    message.value = "";
})

setUser.addEventListener("submit", (event) => {
    event.preventDefault();
    let username = document.getElementById("username");
    socket.send("Username: " + username.value);
    user = username.value
    username.value = "";
})

function messageHandler(message){
    if (message.includes("Error:")){
        entryError.childNodes[1].innerText = message;
        entryError.style.display = "block";
    } else if (message.includes("Username set as:")){
        user = message.slice(17)
        modal.show();
    } else {
        if((chat.scrollTop + chat.clientHeight) === chat.scrollHeight){
            insert("p",{},"",message, isUser(message))
            chat.scrollTop = chat.scrollHeight  
        } else {
            insert("p",{},"",message, isUser(message))
            popupTrigger();
        }
    }
}

function isUser(message){
    return message.slice(0, message.indexOf(":")) === user
}

getModal.addEventListener('shown.bs.modal', () => {
    if(entryError.style.display !== 'none'){
        entryError.style.display = "none";
    }

    if (entry.style.display !== 'none'){
        entry.style.display = 'none';
    }
})

getModal.addEventListener('hide.bs.modal', () => {   
    if (entry.style.display === 'none'){
        entry.style.display = 'block';
    }
})