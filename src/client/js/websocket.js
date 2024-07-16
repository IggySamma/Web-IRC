const socket = new WebSocket('ws://localhost:3000/ws');

socket.addEventListener('open', (event) => {
    socket.send("Hello World!");
});

socket.addEventListener('message', (event) => {
    console.log('Server message: ', event.data);
});