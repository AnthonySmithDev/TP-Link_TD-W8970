<script lang="ts" setup>
import { ref } from "vue";
import { SystemRead } from "../../wailsjs/go/main/App.js";

const loading = ref(false);
const invalid = ref("");
const password = ref("");

async function submit() {
  invalid.value = "";
  loading.value = true;
  const pwd = await SystemRead();
  if (pwd == password.value) {
    invalid.value = "false";
  } else {
    invalid.value = "true";
  }
  loading.value = false;
}

function isInvalid(value: string) {
  if (value == "") {
    return undefined;
  } else if (value == "true") {
    return true;
  } else {
    return false;
  }
}
</script>

<template>
  <article>
    <hgroup>
      <h1>Sign in</h1>
      <h2>A minimalist sign in for router settings</h2>
    </hgroup>
    <form @submit.prevent>
      <label for="password">Password</label>
      <input
        v-model="password"
        :aria-invalid="isInvalid(invalid)"
        type="password"
        id="password"
        name="password"
        aria-label="Password"
        autocomplete="current-password"
        required
      />
      <button v-if="loading" aria-busy="true" class="secondary">
        Please wait…
      </button>
      <button v-else @click="submit" type="submit">Save</button>
    </form>
  </article>
</template>

<style></style>
