const socket = new WebSocket('ws:' + window.location.href.replace("http://","") + 'ws');
const modal = new bootstrap.Modal(document.getElementById("modalView"), {
    backdrop: "static",
    focus: true,
    keyboard: false
})
const entry = document.forms["entryUsername"];
const chatForm = document.forms["liveChat"];
const chat = document.getElementById("chat")
let user;
let popup = bootstrap.Popover.getOrCreateInstance("#message")
popup.setContent({'.popover-body': 'New messages avaliable'})
popup.show()

socket.addEventListener('message', (event) => {
    console.log('Server message: ', event.data);
    
    if (event.data.includes("Error:")){
        let errorBlock = document.getElementById("error");
        errorBlock.childNodes[1].innerText = event.data;
        errorBlock.style.display = "block";
    } else if (event.data === "Success"){
        let errorBlock = document.getElementById("error");
        errorBlock.style.display = "none";
        chat.firstChild.remove()
        modal.show();
    } else {
        if((chat.scrollTop + chat.clientHeight) === chat.scrollHeight){
            insert("p",{},"",event.data)
            chat.scrollTop = chat.scrollHeight  
        } else {
            insert("p",{},"",event.data)
            popup.show();
        }
    }
});


function insert(type = "div", attributes = {}, classes = "message", message){
    if(message){
        let chat = document.getElementById("chat");
        const element = document.createElement(type);
        if (classes) element.classList = classes
        for (let key in attributes){
            element.setAttribute(key, attributes[key])
        }
        element.innerText = message;
        chat.appendChild(element);
    }
    return
}

chat.addEventListener("scroll", () => {
    if((chat.scrollTop + chat.clientHeight) === chat.scrollHeight){
        popup.hide();
    }
    popup.show();
})



chatForm.addEventListener("submit", (event) => {
    event.preventDefault(); 
    let message = document.getElementById("message");
    socket.send(message.value);
    message.value = "";
})

entry.addEventListener("submit", (event) => {
    event.preventDefault();
    let username = document.getElementById("username");
    socket.send("Username: " + username.value);
    user = username.value
    username.value = "";
})
