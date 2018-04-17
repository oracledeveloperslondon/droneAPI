/**
 * messages module
 */
define(['ojs/ojcore', 'knockout', 'ojs/ojbutton', 'factories/WebsocketFactory', 'ojs/ojinputtext', 'ojs/ojformlayout'
], function (oj, ko, io, WebsocketFactory) {
    /**
     * The view model for the messages view template
     */

    function MessagesViewModel() {

        var self = this;
        self.messageArray = ko.observableArray();
        self.yawAmount = ko.observable("0");
        self.rollAmount = ko.observable("0");
        self.pitchAmount = ko.observable("0");
        self.gazAmount = ko.observable("0");

        var i = 0;
        var websocket = WebsocketFactory.getWebsocket();

        self.addMessage = function (event) {

            var array = self.messageArray();
            
            if (event.data) {
                var message = JSON.parse(event.data);
                i = i + 1;
                array.push({
                    "id": i,
                    "yaw": message.movement.yaw,
                    "roll": message.movement.roll,
                    "pitch": message.movement.pitch,
                    "gaz": message.movement.gaz
                });
            }

            self.messageArray(array);

        };

        websocket.onmessage = function (event) {
            self.addMessage(event);
        };

        self.emitMessage = function() {

            var command = {
                "yaw": self.yawAmount(),
                "roll": self.rollAmount(),
                "pitch": self.pitchAmount(),
                "gaz": self.gazAmount()
            };

            websocket.send(JSON.stringify(command));

        };

    };
    
    return new MessagesViewModel();

});
