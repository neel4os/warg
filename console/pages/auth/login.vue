<template>
  <div class="signin">
    <SigninLogin></SigninLogin>
  </div>
</template>

<script lang="ts" setup>
  definePageMeta({
    layout: false,
  });
</script>

<style>
  .signin {
    height: 30%;
    width: 50%;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
  }
</style>
