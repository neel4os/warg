<template>
  <div>
    <SigninEmailPrompt v-if="!showPrompt" @register="is_register" />
    <SigninRegisterPrompt v-if="showPrompt" @register="is_register"/>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from "vue";
  const showPrompt = ref(false);
  function is_register(is_register: boolean) {
    showPrompt.value = is_register;
  }
</script>

<style></style>
