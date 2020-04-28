<template>
  <div id="app">
    <button @click="click_btn">POST リクエストを送るよ</button>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "App",
  methods: {
    async click_btn() {
      await axios
        .post("http://localhost:8081/test", { text: "test preflight request" })
        .then((res) => console.log("OK!", res))
        .catch((err) => console.log(err));
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
