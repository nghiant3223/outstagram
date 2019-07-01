export default (function () {
    let instance;
    let eventType2Handler = {};
    
    function open(params) {
        if (instance !== undefined) {
            throw new Error("Attempt to open WebSocket connection fails");
        }
        
        instance = new WebSocket(`${window.location.protocol === "https:" ? "wss" : "ws"}://${window.location.host}/ws${params !== undefined ? `?${[Object.keys(params).map(param => `${param}=${params[param]}`)].join("&")}` : ""}`);

        instance.onopen = function () {
            console.log("Connect succesfully");
        }

        instance.onclose = function () {
            console.log("Disconnect from server");
        }

        instance.onerror = function (err) {
            console.log("Connect fail: " + err);
        }

        instance.onmessage = function (event) {
            const data = JSON.parse(event.data);

            if (eventType2Handler[data.type] !== undefined) {
                eventType2Handler[data.type].forEach((handler) => handler(data));
            }
        }
    }

    function close() {
        if (instance === undefined) {
            throw new Error("Attempt to close null WebSocket connection fails");
        }

        instance.close();
        instance = undefined;
    }

    function emit(eventType, data) {
        instance.send(JSON.stringify({ type: eventType, data: data }));
    }

    function on(eventName, f) {
        if (eventType2Handler[eventName] === undefined) {
            eventType2Handler[eventName] = [];
        }

        eventType2Handler[eventName].push(f);
    }

    function removeListener(eventName, f) {
        if (eventType2Handler[eventName] !== undefined) {
            eventType2Handler[eventName] = eventType2Handler[eventName].filter(handler => handler !== f);
        }
    }

    return { open, close, emit, on, removeListener };
})();