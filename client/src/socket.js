export default (function () {
    let instance;

    function get() {
        if (!instance) {
            instance = new WebSocket(`${window.location.protocol === "https:" ? "wss" : "ws"}://${window.location.host}/ws`);

            instance.onopen = () => {
                console.log("Connect succesfully");
            }

            instance.onerror = (err) => {
                console.log("Connect fail: " + err);
            }
        }

        return instance;
    }

    function close() {
        if (instance) {
            instance.close();
            instance = undefined;
        }
    }

    function send(message) {
        instance.send(message);
    }

    return { get, close, send };
})();