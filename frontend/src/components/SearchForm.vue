<template>
  <form class="mx-auto" @submit.prevent="search">
    <label
      for="search-input"
      class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white"
      >Search</label
    >
    <div class="relative">
      <input
        v-model="query"
        type="search"
        id="search-input"
        class="block w-full p-4 ps-10 text-sm border border-gray-300 rounded-full"
        placeholder="Search..."
        @input="debouncedSearch"
        required
      />
      <button
        type="submit"
        class="text-white absolute end-2.5 bottom-2.5 bg-blue-500 hover:bg-blue-800 font-medium rounded-full text-sm px-4 py-2"
      >
        Search
      </button>
    </div>
  </form>
</template>

<script setup>
import { ref } from "vue";
import { useEmailStore } from "../stores/emails";
import { debounce } from "@/common/utils";

const { searchBy } = useEmailStore();
const query = ref("");

const search = () => {
  searchBy(query.value);
};

const debouncedSearch = debounce(() => {
  search();
}, 250);
</script>
