<template>
  <!-- Content -->
  <div class="relative h-screen">
    <div class="py-10 lg:py-14">
      <ul class="mt-16 space-y-5">
        <div
          class="max-w-4xl py-2 px-4 sm:px-6 lg:px-8 mx-auto gap-x-2 sm:gap-x-4"
        >
          <!-- Chat Bubble -->
          <div v-for="(message, i) in messages" :key="i" class="py-2 sm:py-4">
            <div class="max-w-4xl mx-auto">
              <div class="gap-x-2 sm:gap-x-4">
                <div class="flex">
                  <!-- icon -->
                  <div>
                    <span v-if="message.role === 'ai'">
                      <img
                        class="inline-block h-10 w-10 rounded-lg"
                        :src="getUserAvatarPlaceholder('ai')"
                        alt="ai icon"
                      />
                    </span>
                    <span
                      v-else
                      class="flex-shrink-0 inline-flex items-center justify-center h-[2.375rem] w-[2.375rem] rounded-full bg-gray-600"
                    >
                      <img
                        class="inline-block h-10 w-10 rounded-lg"
                        :src="getUserAvatarPlaceholder('you', 'red')"
                        alt="human icon"
                      />
                    </span>
                  </div>

                  <!-- initials -->
                  <div>
                    <span class="text-2xl font-bold text-gray-800 mx-3">
                      {{ message.role }}
                    </span>
                    <span class="text-gray-500">
                      {{ moment(message.createdAt).fromNow() }}
                    </span>
                  </div>
                </div>
                <div class="grow mt-2 space-y-3">
                  <p
                    v-if="message.role === 'human'"
                    class="text-gray-800 dark:text-gray-200"
                  >
                    {{ message.text }}
                  </p>
                  <p
                    v-else
                    v-html="message.text"
                    class="text-gray-800 dark:text-gray-200"
                  ></p>
                </div>
              </div>
            </div>
          </div>
          <!-- End Chat Bubble -->
        </div>

        <!-- Chat Bubble -->
        <!--        <li class="max-w-4xl py-2 px-4 sm:px-6 lg:px-8 mx-auto flex gap-x-2 sm:gap-x-4">-->
        <!--          <svg class="flex-shrink-0 w-[2.375rem] h-[2.375rem] rounded-full" width="38" height="38" viewBox="0 0 38 38"-->
        <!--               fill="none" xmlns="http://www.w3.org/2000/svg">-->
        <!--            <rect width="38" height="38" rx="6" fill="#2563EB"/>-->
        <!--            <path-->
        <!--                d="M10 28V18.64C10 13.8683 14.0294 10 19 10C23.9706 10 28 13.8683 28 18.64C28 23.4117 23.9706 27.28 19 27.28H18.25"-->
        <!--                stroke="white" stroke-width="1.5"/>-->
        <!--            <path-->
        <!--                d="M13 28V18.7552C13 15.5104 15.6863 12.88 19 12.88C22.3137 12.88 25 15.5104 25 18.7552C25 22 22.3137 24.6304 19 24.6304H18.25"-->
        <!--                stroke="white" stroke-width="1.5"/>-->
        <!--            <ellipse cx="19" cy="18.6554" rx="3.75" ry="3.6" fill="white"/>-->
        <!--          </svg>-->

        <!--          <div class="grow max-w-[90%] md:max-w-2xl w-full space-y-3">-->
        <!--            &lt;!&ndash; Card &ndash;&gt;-->
        <!--            <div class="space-y-3">-->
        <!--              <p class="text-sm text-gray-800 dark:text-white">-->
        <!--                Preline UI is an open-source set of prebuilt UI components based on the utility-first Tailwind CSS-->
        <!--                framework.-->
        <!--              </p>-->
        <!--              <div class="space-y-1.5">-->
        <!--                <p class="text-sm text-gray-800 dark:text-white">-->
        <!--                  Here're some links to get started-->
        <!--                </p>-->
        <!--                <ul>-->
        <!--                  <li>-->
        <!--                    <a class="text-sm text-blue-600 decoration-2 hover:underline font-medium dark:text-blue-500 dark:hover:text-blue-400"-->
        <!--                       href="../docs/index.html">-->
        <!--                      Installation Guide-->
        <!--                    </a>-->
        <!--                  </li>-->
        <!--                  <li>-->
        <!--                    <a class="text-sm text-blue-600 decoration-2 hover:underline font-medium dark:text-blue-500 dark:hover:text-blue-400"-->
        <!--                       href="../docs/frameworks.html">-->
        <!--                      Framework Guides-->
        <!--                    </a>-->
        <!--                  </li>-->
        <!--                </ul>-->
        <!--              </div>-->
        <!--            </div>-->
        <!--            &lt;!&ndash; End Card &ndash;&gt;-->

        <!--            &lt;!&ndash; Button Group &ndash;&gt;-->
        <!--            <div>-->
        <!--              <div class="sm:flex sm:justify-between">-->
        <!--                <div>-->
        <!--                  <div class="inline-flex border border-gray-200 rounded-full p-0.5 dark:border-gray-700">-->
        <!--                    <button type="button"-->
        <!--                            class="inline-flex flex-shrink-0 justify-center items-center h-8 w-8 rounded-full text-gray-500 hover:bg-blue-100 hover:text-blue-800 focus:z-10 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all dark:hover:bg-blue-900 dark:hover:text-blue-200">-->
        <!--                      <svg class="h-4 w-4" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"-->
        <!--                           viewBox="0 0 16 16">-->
        <!--                        <path-->
        <!--                            d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>-->
        <!--                      </svg>-->
        <!--                    </button>-->
        <!--                    <button type="button"-->
        <!--                            class="inline-flex flex-shrink-0 justify-center items-center h-8 w-8 rounded-full text-gray-500 hover:bg-blue-100 hover:text-blue-800 focus:z-10 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all dark:hover:bg-blue-900 dark:hover:text-blue-200">-->
        <!--                      <svg class="h-4 w-4" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"-->
        <!--                           viewBox="0 0 16 16">-->
        <!--                        <path-->
        <!--                            d="M8.864 15.674c-.956.24-1.843-.484-1.908-1.42-.072-1.05-.23-2.015-.428-2.59-.125-.36-.479-1.012-1.04-1.638-.557-.624-1.282-1.179-2.131-1.41C2.685 8.432 2 7.85 2 7V3c0-.845.682-1.464 1.448-1.546 1.07-.113 1.564-.415 2.068-.723l.048-.029c.272-.166.578-.349.97-.484C6.931.08 7.395 0 8 0h3.5c.937 0 1.599.478 1.934 1.064.164.287.254.607.254.913 0 .152-.023.312-.077.464.201.262.38.577.488.9.11.33.172.762.004 1.15.069.13.12.268.159.403.077.27.113.567.113.856 0 .289-.036.586-.113.856-.035.12-.08.244-.138.363.394.571.418 1.2.234 1.733-.206.592-.682 1.1-1.2 1.272-.847.283-1.803.276-2.516.211a9.877 9.877 0 0 1-.443-.05 9.364 9.364 0 0 1-.062 4.51c-.138.508-.55.848-1.012.964l-.261.065zM11.5 1H8c-.51 0-.863.068-1.14.163-.281.097-.506.229-.776.393l-.04.025c-.555.338-1.198.73-2.49.868-.333.035-.554.29-.554.55V7c0 .255.226.543.62.65 1.095.3 1.977.997 2.614 1.709.635.71 1.064 1.475 1.238 1.977.243.7.407 1.768.482 2.85.025.362.36.595.667.518l.262-.065c.16-.04.258-.144.288-.255a8.34 8.34 0 0 0-.145-4.726.5.5 0 0 1 .595-.643h.003l.014.004.058.013a8.912 8.912 0 0 0 1.036.157c.663.06 1.457.054 2.11-.163.175-.059.45-.301.57-.651.107-.308.087-.67-.266-1.021L12.793 7l.353-.354c.043-.042.105-.14.154-.315.048-.167.075-.37.075-.581 0-.211-.027-.414-.075-.581-.05-.174-.111-.273-.154-.315l-.353-.354.353-.354c.047-.047.109-.176.005-.488a2.224 2.224 0 0 0-.505-.804l-.353-.354.353-.354c.006-.005.041-.05.041-.17a.866.866 0 0 0-.121-.415C12.4 1.272 12.063 1 11.5 1z"/>-->
        <!--                      </svg>-->
        <!--                    </button>-->
        <!--                  </div>-->
        <!--                  <button type="button"-->
        <!--                          class="py-2 px-3 inline-flex justify-center items-center gap-x-2 rounded-full border border-transparent text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2 transition-all text-sm dark:hover:bg-slate-800 dark:hover:text-gray-400 dark:hover:border-gray-900 dark:focus:ring-gray-900 dark:focus:ring-offset-gray-800">-->
        <!--                    <svg class="h-3.5 w-3.5" xmlns="http://www.w3.org/2000/svg" width="16" height="16"-->
        <!--                         fill="currentColor" viewBox="0 0 16 16">-->
        <!--                      <path-->
        <!--                          d="M4 1.5H3a2 2 0 0 0-2 2V14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3.5a2 2 0 0 0-2-2h-1v1h1a1 1 0 0 1 1 1V14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3.5a1 1 0 0 1 1-1h1v-1z"/>-->
        <!--                      <path-->
        <!--                          d="M9.5 1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5h3zm-3-1A1.5 1.5 0 0 0 5 1.5v1A1.5 1.5 0 0 0 6.5 4h3A1.5 1.5 0 0 0 11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3z"/>-->
        <!--                    </svg>-->
        <!--                    Copy-->
        <!--                  </button>-->
        <!--                  <button type="button"-->
        <!--                          class="py-2 px-3 inline-flex justify-center items-center gap-x-2 rounded-full border border-transparent text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2 transition-all text-sm dark:hover:bg-slate-800 dark:hover:text-gray-400 dark:hover:border-gray-900 dark:focus:ring-gray-900 dark:focus:ring-offset-gray-800">-->
        <!--                    <svg class="h-3.5 w-3.5" xmlns="http://www.w3.org/2000/svg" width="16" height="16"-->
        <!--                         fill="currentColor" viewBox="0 0 16 16">-->
        <!--                      <path-->
        <!--                          d="M13.5 1a1.5 1.5 0 1 0 0 3 1.5 1.5 0 0 0 0-3zM11 2.5a2.5 2.5 0 1 1 .603 1.628l-6.718 3.12a2.499 2.499 0 0 1 0 1.504l6.718 3.12a2.5 2.5 0 1 1-.488.876l-6.718-3.12a2.5 2.5 0 1 1 0-3.256l6.718-3.12A2.5 2.5 0 0 1 11 2.5zm-8.5 4a1.5 1.5 0 1 0 0 3 1.5 1.5 0 0 0 0-3zm11 5.5a1.5 1.5 0 1 0 0 3 1.5 1.5 0 0 0 0-3z"/>-->
        <!--                    </svg>-->
        <!--                    Share-->
        <!--                  </button>-->
        <!--                </div>-->

        <!--                <div class="mt-1 sm:mt-0">-->
        <!--                  <button type="button"-->
        <!--                          class="py-2 px-3 inline-flex justify-center items-center gap-x-2 rounded-full border border-transparent text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2 transition-all text-sm dark:hover:bg-slate-800 dark:hover:text-gray-400 dark:hover:border-gray-900 dark:focus:ring-gray-900 dark:focus:ring-offset-gray-800">-->
        <!--                    <svg class="h-3.5 w-3.5" xmlns="http://www.w3.org/2000/svg" width="16" height="16"-->
        <!--                         fill="currentColor" viewBox="0 0 16 16">-->
        <!--                      <path fill-rule="evenodd" d="M8 3a5 5 0 1 0 4.546 2.914.5.5 0 0 1 .908-.417A6 6 0 1 1 8 2v1z"/>-->
        <!--                      <path-->
        <!--                          d="M8 4.466V.534a.25.25 0 0 1 .41-.192l2.36 1.966c.12.1.12.284 0 .384L8.41 4.658A.25.25 0 0 1 8 4.466z"/>-->
        <!--                    </svg>-->
        <!--                    New answer-->
        <!--                  </button>-->
        <!--                </div>-->
        <!--              </div>-->
        <!--            </div>-->
        <!--            &lt;!&ndash; End Button Group &ndash;&gt;-->
        <!--          </div>-->
        <!--        </li>-->
        <!-- End Chat Bubble -->
      </ul>
    </div>

    <!-- Search -->
    <footer
      class="sticky bottom-0 z-10 bg-white border-t border-gray-200 pt-2 pb-3 sm:pt-4 sm:pb-6 dark:bg-slate-900 dark:border-gray-700"
    >
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center mb-3">
          <button
            type="button"
            class="inline-flex justify-center items-center gap-x-2 rounded-md font-medium text-gray-800 hover:text-blue-600 text-xs sm:text-sm dark:text-gray-200 dark:hover:text-blue-500"
          >
            <span v-if="isConnected">
              <span class="shadow-md bg-cyan-500 rounded">&nbsp;</span>
              Connected
            </span>
            <span v-else>
              <span class="shadow-md bg-red-500 rounded">&nbsp;</span>
              Disconnected
            </span>
          </button>

          <button
            v-if="isWaitingForResponse"
            type="button"
            class="py-1.5 px-2 inline-flex justify-center items-center gap-2 rounded-md border font-medium bg-white text-gray-700 shadow-sm align-middle hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-white focus:ring-blue-600 transition-all text-xs dark:bg-slate-900 dark:hover:bg-slate-800 dark:border-gray-700 dark:text-gray-400 dark:hover:text-white dark:focus:ring-offset-gray-800"
          >
            <svg
              class="w-3 h-3"
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              fill="currentColor"
              viewBox="0 0 16 16"
            >
              <path
                d="M5 3.5h6A1.5 1.5 0 0 1 12.5 5v6a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 11V5A1.5 1.5 0 0 1 5 3.5z"
              />
            </svg>
            Stop generating
          </button>
        </div>

        <!-- Input -->
        <div class="relative">
          <textarea
            @keyup.enter="sendMessage"
            v-model="newMessage"
            id="new_message"
            class="p-4 pb-12 block w-full border-gray-200 rounded-md text-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-slate-900 dark:border-gray-700 dark:text-gray-400"
            placeholder="Ask me anything..."
          ></textarea>

          <!-- Toolbar -->
          <div
            class="absolute bottom-px inset-x-px p-2 rounded-b-md bg-white dark:bg-slate-900"
          >
            <div class="flex justify-between items-center">
              <div></div>
              <!-- Button Group -->
              <div class="flex items-center gap-x-1">
                <!-- Send Button -->
                <button
                  type="button"
                  @click="sendMessage"
                  class="inline-flex flex-shrink-0 justify-center items-center h-8 w-8 rounded-md text-white bg-blue-600 hover:bg-blue-500 focus:z-10 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
                >
                  <svg
                    class="h-3.5 w-3.5"
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    fill="currentColor"
                    viewBox="0 0 16 16"
                  >
                    <path
                      d="M15.964.686a.5.5 0 0 0-.65-.65L.767 5.855H.766l-.452.18a.5.5 0 0 0-.082.887l.41.26.001.002 4.995 3.178 3.178 4.995.002.002.26.41a.5.5 0 0 0 .886-.083l6-15Zm-1.833 1.89L6.637 10.07l-.215-.338a.5.5 0 0 0-.154-.154l-.338-.215 7.494-7.494 1.178-.471-.47 1.178Z"
                    />
                  </svg>
                </button>
                <!-- End Send Button -->
              </div>
              <!-- End Button Group -->
            </div>
          </div>
          <!-- End Toolbar -->
        </div>
        <!-- End Input -->
      </div>
    </footer>
    <!-- End Search -->
  </div>
  <!-- End Content -->
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import moment from "moment";
import { baseUrl, url } from "../assets/constants";

// TODO: fix me
let socket = new WebSocket(url);

const messages = ref([]);

const newMessage = ref("");

const waitingID = "waiting";
const startToken = "START";
const endToken = "END";
const isConnected = ref(false);
const shouldScroll = ref(false);
const isWaitingForResponse = ref(false);

function getUserAvatarPlaceholder(
  fullName: string,
  color: string = "20c997"
): string {
  return `https://placehold.co/100/${color}/FFFFFF.png?text=${fullName}&font=montserrat`;
}

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
  const message = newMessage.value;
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
    newMessage.value = "";
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


onMounted(() => {
  setupSocket();
  fetchConversations().catch(console.error);
});
</script>
