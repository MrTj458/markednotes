<script setup>
import NavBar from "./components/NavBar.vue";
</script>

<template>
  <header>
    <NavBar />
  </header>
  <RouterView />
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Nunito:wght@200;300;400;600;700;800;900&display=swap");

:root {
  --gray-color: #403e41;
  --white-color: #e7e6f7;
  --blue-color: #39a2ae;
  --orange-color: #ef8354;
  --yellow-color: #dda448;
}

*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

* {
  font-family: "Nunito", sans-serif;
}

body {
  font-size: 16px;
  line-height: 1.6;
  background-color: var(--gray-color);
  color: var(--white-color);
}

a {
  color: var(--blue-color);
}
</style>
