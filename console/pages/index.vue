<template>
  <div>
    <q-btn-group>
      <q-btn color="primary" size="xl" label="Sign IN" @click="clickHandler1" />
      <q-btn color="primary" size="xl" label="Sign Up" @click="clickHandler2" />
    </q-btn-group>
  </div>
</template>


<script lang="ts" setup>
  import { useRouter } from 'vue-router';

  const router = useRouter();

  function clickHandler1() {
    // Your existing clickHandler1 code
  }

  function clickHandler2() {
    router.push('/signup');
  }

  definePageMeta({
    layout: "non-default",
  });
</script>
