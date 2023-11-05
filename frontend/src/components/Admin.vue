<script lang="ts" setup>
</script>

<template>
  <div class="content">
    <article class="navbar">
      <router-link role="button" to="/wireless">Wireless</router-link>
      <router-link role="button" to="/system">System</router-link>
      <router-link role="button" to="/">Logout</router-link>
    </article>
    <router-view></router-view>
  </div>
</template>

<style scoped>
.navbar {
  display: flex;
  justify-content: space-around;
}
.content {
  height: 85vh;
}
</style>
