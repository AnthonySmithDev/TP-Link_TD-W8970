<script lang="ts" setup>
import { ref, reactive } from "vue";
import { SystemRead, SystemWrite } from "../../wailsjs/go/main/App.js";

const loading = ref(false);
const password = reactive({
  old: "",
  new: "",
  conf: "",
});
const invalid = reactive({
  old: "",
  new: "",
  conf: "",
});

async function submit() {
  invalid.old = "";
  invalid.new = "";
  invalid.conf = "";

  if (password.new != password.conf) {
    invalid.conf = "true";
    return;
  }

  loading.value = true;
  const pwd = await SystemRead();
  if (pwd == password.old) {
    await SystemWrite(password.new);
    invalid.old = "false";
    invalid.new = "false";
    invalid.conf = "false";
  } else {
    invalid.old = "true";
  }
  loading.value = false;
}

function isInvalid(value: string) {
  if (value == "") {
    return undefined
  } else if (value == "true") {
    return true
  } else {
    return false
  }
}
</script>

<template>
  <article>
    <form @submit.prevent aria-invalid="true">
      <label for="old_password">Old Password</label>
      <input
        v-model="password.old"
        :aria-invalid="isInvalid(invalid.old)"
        type="password"
        id="old_password"
        name="old_password"
        aria-label="Old Password"
        autocomplete="current-password"
        required
      />
      <label for="new_password">New Password</label>
      <input
        v-model="password.new"
        :aria-invalid="isInvalid(invalid.new)"
        type="password"
        id="new_password"
        name="new_password"
        aria-label="New Password"
        autocomplete="current-password"
        required
      />
      <label for="confirm_password">Confirm Password</label>
      <input
        v-model="password.conf"
        :aria-invalid="isInvalid(invalid.conf)"
        type="password"
        id="confirm_password"
        name="password"
        aria-label="Confirm Password"
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
