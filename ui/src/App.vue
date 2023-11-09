<template>
  <div class="container mx-auto mt-10">
    <!-- Title -->
    <div class="max-w-4xl px-4 sm:px-6 lg:px-8 mx-auto text-center">
      <!--TODO: logo here -->

      <h1 class="text-3xl font-bold text-gray-800 sm:text-4xl dark:text-white">
        Welcome to Private GPT in Go
      </h1>
      <p class="mt-3 text-gray-600 dark:text-gray-400">
        Interact with your documents using the power of GPT, 100% privately, no
        data leaks
      </p>
    </div>
    <!-- End Title -->

    <div class="border-b border-gray-200 dark:border-gray-700">
      <nav
        class="-mb-0.5 flex justify-center space-x-6"
        aria-label="Tabs"
        role="tablist"
      >
        <button
          type="button"
          class="hs-tab-active:font-semibold hs-tab-active:border-blue-600 hs-tab-active:text-blue-600 py-4 px-1 inline-flex items-center gap-2 border-b-[3px] border-transparent text-sm whitespace-nowrap text-gray-500 hover:text-blue-600 active"
          id="horizontal-alignment-item-1"
          data-hs-tab="#horizontal-alignment-1"
          aria-controls="horizontal-alignment-1"
          role="tab"
        >
          Chat
        </button>
        <button
          type="button"
          class="hs-tab-active:font-semibold hs-tab-active:border-blue-600 hs-tab-active:text-blue-600 py-4 px-1 inline-flex items-center gap-2 border-b-[3px] border-transparent text-sm whitespace-nowrap text-gray-500 hover:text-blue-600"
          id="horizontal-alignment-item-2"
          data-hs-tab="#horizontal-alignment-2"
          aria-controls="horizontal-alignment-2"
          role="tab"
        >
          Files
        </button>
      </nav>
    </div>

    <div class="mt-3">
      <div
        id="horizontal-alignment-1"
        role="tabpanel"
        aria-labelledby="horizontal-alignment-item-1"
      >
        <Chat />
      </div>
      <div
        id="horizontal-alignment-2"
        class="hidden"
        role="tabpanel"
        aria-labelledby="horizontal-alignment-item-2"
      >
        <Files />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Chat from "./components/Chat.vue";
import Files from "./components/Files.vue";
</script>
