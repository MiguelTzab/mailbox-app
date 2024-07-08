<template>
  <div class="mx-auto bg-white shadow-md rounded-lg overflow-hidden">
    <DataTable
      v-model:itemsPerPage="itemsPerPage"
      :config="tableConfig"
      :loading="store.loading"
      :list="store.emails"
      :headers="headers"
      :totalItems="store.totalEmails"
      @onPaginationChange="
        (page, limit) => store.searchBy(undefined, page, limit)
      "
    >
      <template v-slot:[`col-subject`]="{ item }">
        {{ truncate(item.subject, 40) }}
      </template>
      <template v-slot:[`col-to`]="{ item }">
        {{ truncate(item.to, 60, ",") }}
      </template>
    </DataTable>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { truncate } from "@/common/utils";
import { useEmailStore } from "../stores/emails";
import DataTable from "./DataTable.vue";

const store = useEmailStore();
const itemsPerPage = ref(10);
const tableConfig = {
  rowActions: () => {
    return [
      {
        text: "See",
        value: "view",
      },
    ];
  },
  actions: (type, item) => {
    const actionsMap = {
      view: () => showItem(item),
    };
    actionsMap[type]();
  },
};
const headers = [
  {
    text: "subject",
    value: "subject",
  },
  {
    text: "from",
    value: "from",
  },
  {
    text: "to",
    value: "to",
  },
];

const showItem = (item) => {
  console.log(item);
};
</script>
