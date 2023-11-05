<script lang="ts" setup>
import { ref, reactive, onMounted } from "vue";
import { WirelessRead, WirelessWrite } from "../../wailsjs/go/main/App.js";

const loading = ref(false);
const wireless = reactive({
  name: "",
  enable: false,
  password: "",
});

onMounted(async () => {
  let data = await WirelessRead();
  wireless.name = data.ssid;
  wireless.enable = data.enable;
  wireless.password = data.password;
});

async function submit() {
  loading.value = true;
  await WirelessWrite(wireless.enable, wireless.name, wireless.password);
  loading.value = false;
}
</script>

<template>
  <article>
    <hgroup>
      <h1>Wireless Settings</h1>
      <h2> Customize Wi-Fi settings to configure your network.</h2>
    </hgroup>
    <form @submit.prevent>
      <fieldset>
        <label for="switch">
          <input
            v-model="wireless.enable"
            type="checkbox"
            id="switch"
            name="switch"
            role="switch"
          />
          Enable
        </label>
      </fieldset>
      <label for="ssid">Network Name (SSID)</label>
      <input
        v-model="wireless.name"
        type="text"
        id="ssid"
        name="ssid"
        placeholder="Network Name (SSID)"
        aria-label="Network Name"
        required
      />
      <label for="password">Password</label>
      <input
        v-model="wireless.password"
        type="password"
        id="password"
        name="password"
        placeholder="Password"
        aria-label="Password"
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
