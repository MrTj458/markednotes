<template>
  <nav class="nav-bar">
    <ul class="nav-list">
      <li>
        <RouterLink :to="{ name: 'home' }" class="nav-item"
          >Marked Notes</RouterLink
        >
      </li>
    </ul>
    <ul class="nav-list">
      <li>
        <RouterLink :to="{ name: 'login' }" class="nav-item">Login</RouterLink>
      </li>
      <li>
        <RouterLink :to="{ name: 'register' }" class="nav-item">
          Register
        </RouterLink>
      </li>
    </ul>
  </nav>
</template>

<script setup></script>

<style>
.nav-bar {
  background-color: var(--orange-color);
  display: flex;
  align-items: center;
  width: 100%;
  justify-content: space-between;
  padding: 1rem 2rem;
}

.nav-list {
  list-style: none;
  display: flex;
  justify-content: center;
  gap: 2em;
}

.nav-item {
  font-size: 1.2em;
  text-decoration: none;
  color: var(--white-color);
  border-radius: 4px;
  padding: 0.25rem 0.5rem;
}
</style>
