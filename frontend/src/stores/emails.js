import { defineStore } from "pinia";
import axiosInstance from "../axios";

const SEARCH_COLS = ["From", "To", "Subject", "MessageId"];

export const useEmailStore = defineStore("emails", {
  state: () => ({
    emails: [],
    loading: false,
    error: null,
    query: "",
    page: 1,
    limit: 10,
    totalEmails: 0,
  }),
  actions: {
    async searchBy(query, page = 1, limit = 10) {
      if (query) {
        this.query = query;
      }

      this.error = null;
      try {
        this.loading = true;
        const response = await axiosInstance.get("/emails", {
          params: {
            q: this.query,
            page: page,
            limit: limit,
            cols: SEARCH_COLS,
          },
        });
        this.emails = response.data.items;
        this.page = response.data.page;
        this.limit = response.data.page_size;
        this.totalEmails = response.data.total_results;
      } catch (error) {
        this.error = error;
      } finally {
        this.loading = false;
      }
    },
  },
});
