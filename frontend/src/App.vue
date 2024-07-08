<template>
  <div class="min-h-screen bg-gray-100">
    <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <SearchForm />
      <hr class="my-8" />
      <EmailList />
    </div>
  </div>
</template>

<script>
import SearchForm from "./components/SearchForm.vue";
import EmailList from "./components/EmailList.vue";

export default {
  components: {
    SearchForm,
    EmailList,
  },
};
</script>
