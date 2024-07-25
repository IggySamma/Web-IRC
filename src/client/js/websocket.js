const socket = new WebSocket('ws:' + window.location.href.replace("http://","") + 'ws');
const chat = document.getElementById("liveChat")

socket.addEventListener('message', (event) => {
    insert("p",{},"",event.data)
    console.log('Server message: ', event.data);
});


function insert(type = "div", attributes = {}, classes = "message", message){
    if(message){
        let chat = document.getElementById("chat")
        const element = document.createElement(type)
        if (classes) element.classList = classes
        for (let key in attributes){
            element.setAttribute(key, attributes[key])
        }
        element.innerText = message
        chat.appendChild(element)
    }
    return
}

chat.addEventListener("submit", (event) => {
    event.preventDefault(); 
    let message = document.getElementById("message")
    socket.send(message.value)
    message.value = ""
})

