<template>
  <div class="grid grid-cols-4 gap-4">
    <input
      type="file"
      id="fileInput"
      ref="fileInput"
      @change="handleFileUpload"
      style="display: none"
    />

    <label for="fileInput"
      class="flex bg-slate-50 flex-col border shadow-sm rounded-xl dark:bg-gray-800 dark:border-gray-700 dark:shadow-slate-700/[.7]"
    >
      <div class="p-4 md:p-5">
        <svg
          width="30px"
          height="30px"
          stroke-width="2"
          viewBox="0 0 24 24"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
          color="#585555"
        >
          <path
            d="M8 12H12M16 12H12M12 12V8M12 12V16"
            stroke="#585555"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          ></path>
          <path
            d="M12 22C17.5228 22 22 17.5228 22 12C22 6.47715 17.5228 2 12 2C6.47715 2 2 6.47715 2 12C2 17.5228 6.47715 22 12 22Z"
            stroke="#585555"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          ></path>
        </svg>
        <h3 class="text-lg font-bold text-gray-800 dark:text-white">
          upload new file
        </h3>
      </div>
    </label>

    <div
      v-for="file in files"
      :key="file.id"
      class="flex flex-col bg-white border shadow-sm rounded-xl dark:bg-gray-800 dark:border-gray-700 dark:shadow-slate-700/[.7]"
    >
      <div class="p-4 md:p-5">
        <h3 class="text-lg font-bold text-gray-800 dark:text-white">
          Card title
        </h3>

        <div class="flex justify-between">
          <div><span>1.2mb</span> . pdf</div>
          <button>
            <svg
              width="16px"
              height="16px"
              viewBox="0 0 24 24"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
              color="#f50f0f"
              stroke-width="2"
            >
              <path
                d="M20 9L18.005 20.3463C17.8369 21.3026 17.0062 22 16.0353 22H7.96474C6.99379 22 6.1631 21.3026 5.99496 20.3463L4 9"
                fill="#f50f0f"
              ></path>
              <path
                d="M20 9L18.005 20.3463C17.8369 21.3026 17.0062 22 16.0353 22H7.96474C6.99379 22 6.1631 21.3026 5.99496 20.3463L4 9H20Z"
                stroke="#f50f0f"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              ></path>
              <path
                d="M21 6H15.375M3 6H8.625M8.625 6V4C8.625 2.89543 9.52043 2 10.625 2H13.375C14.4796 2 15.375 2.89543 15.375 4V6M8.625 6H15.375"
                stroke="#f50f0f"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              ></path>
            </svg>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { baseUrl } from "../assets/constants";

const files = ref([]);
const fileInput = ref(null);

async function uploadFile() {
  const selectedFile = fileInput.value.files[0];

  if (selectedFile) {
    const formData = new FormData();
    formData.append("file", selectedFile);

    try {
      const response = await fetch(baseUrl + "/upload", {
        method: "POST",
        body: formData,
      });

      if (response.ok) {
        fetchFiles();
        fileInput.value.value = "";
      } else {
        console.error("File upload failed");
      }
    } catch (error) {
      console.error(error);
    }
  }
}

function handleFileUpload() {
  const selectedFile = fileInput.value.files[0];
  
  console.log(selectedFile)
}

async function fetchFiles() {
  const response = await fetch(baseUrl + "/documents");
  files.value = await response.json();
}

onMounted(() => {
  fetchFiles().catch(console.error);
});
</script>
