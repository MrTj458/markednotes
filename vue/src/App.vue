<script setup>
import NotificationViewer from "./components/NotificationViewer.vue";
</script>

<template>
  <NotificationViewer />
  <RouterView />
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Nunito:wght@200;300;400;600;700;800;900&display=swap");

:root {
  --light-gray-color: #59585a;
  --gray-color: #403e41;
  --white-color: #e7e6f7;
  --blue-color: #39a2ae;
  --orange-color: #ef8354;
  --yellow-color: #dda448;
}

*,
*:before,
*:after {
  font-family: "Nunito", sans-serif;
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

html {
  height: 100%;
}

body {
  font-size: 16px;
  line-height: 1.6;
  background-color: var(--gray-color);
  color: var(--white-color);
  height: 100%;
}

#app {
  height: 100%;
}

a {
  color: var(--blue-color);
}
</style>
