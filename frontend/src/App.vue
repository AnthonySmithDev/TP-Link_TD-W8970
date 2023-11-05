<script lang="ts" setup>
</script>

<template>
  <div class="center">
    <div class="element">
      <router-view></router-view>
    </div>
  </div>
</template>

<style scoped>
.center {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
.element {
  width: 400px;
}
</style>
