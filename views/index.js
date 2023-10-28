import {createApp, onMounted, ref} from 'https://unpkg.com/vue@3/dist/vue.esm-browser.js'

const baseUrl = "http://localhost:1337/v1"
const url = "ws://localhost:1337/v1/chat"
let socket = new WebSocket(url);

const app = createApp({
    delimiters: ["{", "}"],
    setup() {
        const messages = ref([])
        const files = ref([])
        const humanFormat = window.humanFormat

        const newMessage = ref("")

        const waitingID = "waiting";
        const startToken = "START";
        const endToken = "END";
        const isConnected = ref(false);
        const shouldScroll = ref(false);
        const isWaitingForResponse = ref(false);


        function setupSocket() {
            socket.onopen = () => {
                isConnected.value = true;
                console.debug("[socket.onopen]: WebSocket connection established!");
            };

            socket.onmessage = (event) => {
                try {
                    handleIncomingMessage(JSON.parse(event.data));
                } catch (e) {
                    console.error("[socket.onmessage]: Error parsing JSON data:", e);
                }
            };
        }

        function sendMessage() {
            const message = newMessage.value
            shouldScroll.value = true;

            messages.value.push({
                id: "NO_HUMAN_MSG_ID",
                role: "human",
                time: new Date().toDateString(),
                text: message,
            });

            isWaitingForResponse.value = true;

            if (socket !== null && socket.readyState === WebSocket.OPEN) {
                socket.send(message);
                newMessage.value = ""
                console.debug("Message sent:", message);
            } else {
                console.warn("WebSocket is not connected or ready to send messages.");
            }
        }

        function handleIncomingMessage(data) {
            if (!data) {
                return;
            }

            // if the message is a buffer, start message token
            if (data.isBuffered) {
                if (data.token === startToken) {
                    console.debug("Starting buffer message from the server.");
                    messages.value.push({
                        id: waitingID,
                        role: "ai",
                        text: "...",
                        time: new Date().toDateString(),
                    });
                }

                const idx = messages.value.findIndex((x) => x.id === waitingID);
                if (idx > -1) {
                    if (data.token !== endToken) {
                        messages.value[idx].text = data.message;
                    } else {
                        messages.value[idx].id = "";
                        isWaitingForResponse.value = false;
                        shouldScroll.value = false;
                        console.debug(
                            "Received buffer end message from the server: ",
                            messages.value[idx].text
                        );
                        console.info("Received message:", data.message);
                    }
                }
            } else {
                messages.value.push({
                    id: "NO_AI_MSG_ID",
                    role: "ai",
                    time: new Date().toDateString(),
                    text: data.message,
                });

                isWaitingForResponse.value = false;
                shouldScroll.value = false;
            }
        }


        async function fetchConversations() {
            const response = await fetch(baseUrl + "/conversations");
            messages.value = await response.json();
        }

        async function fetchFiles() {
            const response = await fetch(baseUrl + "/documents");
            files.value = await response.json();
        }

        onMounted(() => {
            setupSocket()
            fetchConversations().catch(console.error)
            fetchFiles().catch(console.error)
        })

        return {
            files,
            sendMessage,
            isConnected,
            humanFormat,
            newMessage,
            handleIncomingMessage,
            messages
        }
    }
})

app.mount('#app')

