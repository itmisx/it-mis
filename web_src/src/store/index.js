import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    route: "",
  },
  getters: {
    route: (state) => {
      return state.route;
    },
  },
  mutations: {
    setRoute: (state, r) => {
      state.route = r;
    },
  },
  actions: {},
  modules: {},
});
