
//======== function for send data from client to server ============
function send() {
    var msg = document.getElementById("message").value
    var sender = document.getElementById("username").value
    console.log("username"+ sender+ "and message"+ msg)
    let message = {
        Message :msg,
        From : sender
    }
    var ws = new WebSocket("ws://" + window.location.host + "/Socket");
    ws.onopen = function() {
        ws.send(JSON.stringify(message))
    }
    $("#chathistory").append(`<div class="col-sm-12 " >
            <div class="card" style="width:400px">
                <div class="card-body">
                    <h4 class="card-title">`+sender+`</h4>
                    <p class="card-text">`+ msg+`</p>
            </div>
           </div>
`);
    ws.onmessage = function(event) {
        var m = JSON.parse(event.data);
        console.log(m);
        console.debug("Received message", m.Message);
        $("#chathistory").append(`<div class="col-sm-12 "  style="color: darkblue">
            <div class="card" style="width:400px">
                <div class="card-body">
                    <h4 class="card-title">`+m.From+`</h4>
                    <p class="card-text">`+ m.Message+`</p>
            </div>
           </div>
`);
    }
    ws.onerror = function(event) {
        console.debug(event)
    }
}
//GetAllMessagesRooms

function getAllRooms() {
    // console.log("username"+ sender+ "and message"+ msg)
    let message = {
        Client :"string",
        Messages :"[]string"
    }
    var ws = new WebSocket("ws://" + window.location.host + "/GetAllMessagesRooms");

    ws.onmessage = function(event) {
        var m = JSON.parse(event.data);
        console.log(m);
        // console.debug("Received message", m.Messages);
for (var i = 0 ; i < m.length; i++){
    $("#chatrooms").append(`<div class="col-sm-12 "  style="color: darkblue">
            <div class="card" style="width:400px">
                <div class="card-body">
                    <h4 class="card-title">`+m[i].Client+`</h4>
                   
`);
    for(var j = 0 ;j< m[i].Messages.length; j++){
        $("#chatrooms").append(`<p class="card-text">`+ m[i].Messages[j]+`</p>
            </div>
            </div>`);
    }
}

    };
}

