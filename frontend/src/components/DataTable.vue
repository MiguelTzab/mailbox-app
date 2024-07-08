<template>
  <div class="card">
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead>
          <div v-if="loading" class="w-full h-2 bg-blue-200">
            <div class="h-full bg-blue-500 animate-pulse"></div>
          </div>
          <tr>
            <th
              v-for="header in computedHeaders"
              :key="header.text"
              @click="header.sortable && sortColumn(header.value)"
              class="px-6 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer"
            >
              <svg
                v-if="header.sortable"
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5 text-gray-500"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M6 10l4-4 4 4H6z"
                  clip-rule="evenodd"
                />
              </svg>
              {{ header.text }}
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="item in list" :key="item.id" class="hover:bg-gray-100">
            <td
              v-for="header in computedHeaders"
              :key="header.value"
              class="px-6 py-4 whitespace-nowrap text-sm text-gray-500"
            >
              <slot :name="getSlotColName(header.value)" :item="item">
                {{ dotNotation(header.value, item) }}
              </slot>
            </td>
            <td
              class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium"
            >
              <div class="flex items-center justify-end">
                <button
                  v-for="action in actionsToShow(item)"
                  :key="action.value"
                  @click="handleActionClick(action.value, item)"
                  :class="['mr-2', action.buttonClass]"
                  class="px-3 py-1 rounded"
                >
                  {{ action.text }}
                </button>
              </div>
            </td>
          </tr>
          <tr v-if="list.length === 0">
            <td
              :colspan="computedHeaders.length"
              class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-500 text-left"
            >
              There are no records for your search parameters. Please try
              another search.
            </td>
          </tr>
        </tbody>
      </table>
      <div class="flex justify-between items-center p-4" v-if="totalItems > 0">
        <div>
          <span>Items per page: </span>
          <select v-model="itemsPerPage" class="border rounded px-2 py-1">
            <option
              v-for="option in rowsPerPageItems"
              :key="option"
              :value="option"
            >
              {{ option }}
            </option>
          </select>
        </div>
        <div>
          <button
            @click="goTo(page - 1)"
            :disabled="page === 1"
            class="mr-2 bg-gray-300 px-3 py-1 rounded"
          >
            Previous
          </button>
          <span>Page {{ page }} of {{ computedTotalPages }}</span>
          <button
            @click="goTo(page + 1)"
            :disabled="page === computedTotalPages"
            class="ml-2 bg-gray-300 px-3 py-1 rounded"
          >
            Next
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, defineProps, defineEmits } from "vue";
import { dotNotation } from "../common/utils";

const props = defineProps({
  config: {
    type: Object,
    default: () => ({
      rowActions: null,
      actions: null,
    }),
  },
  headers: {
    type: Array,
    default: () => [],
  },
  list: {
    type: Array,
    default: () => [],
  },
  totalItems: {
    type: Number,
    default: 0,
  },
  loading: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(["onPaginationChange"]);
const itemsPerPage = ref(10);
const rowsPerPageItems = ref([10, 25, 50]);
const page = ref(1);

const handleActionClick = (type, item) => {
  props.config.actions(type, item);
};

const goTo = (to) => {
  page.value = to;
};

const actionsToShow = (item) => {
  if (props.config.rowActions && item) {
    return props.config.rowActions(item);
  }
  return [];
};

const getSlotColName = (column) => `col-${column}`;

const computedHeaders = computed(() => {
  if (props.config.actions && props.config.actions.length) {
    return [
      ...props.headers,
      { text: "Acciones", value: "actions", sortable: false },
    ];
  }

  return props.headers;
});

const computedTotalPages = computed(() =>
  Math.ceil(props.totalItems / itemsPerPage.value)
);

const notify = () => {
  emit("onPaginationChange", page.value, itemsPerPage.value);
};

watch([page, itemsPerPage], notify);
</script>
