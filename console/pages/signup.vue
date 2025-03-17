<style></style>
<template>
  <div class="q-pa-md" style="max-width: 300px; margin: auto">
    <h5>Customer Sign Up</h5>
    <q-form @submit="submitForm" class="q-gutter-md">
      <q-input v-model="form.customername" label="Customer Name" />
      <q-input v-model="form.firstname" label="First Name" />
      <q-input v-model="form.lastname" label="Last Name" />
      <q-input v-model="form.email" label="Email" />
      <q-btn type="submit" label="Sign Up" color="primary" class="q-mt-md" />
    </q-form>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from "vue";

  definePageMeta({
    layout: "non-default",
  });

  const form = ref({
    customername: "",
    firstname: "",
    lastname: "",
    email: "",
  });

  const submitForm = async () => {
    console.log(form.value);
    
    // try {
    //   const response = await onboardingService(form.value);
    //   console.log("Registration successful:", response);
    // } catch (error) {
    //   console.error("Registration failed:", error);
    // }
  };

  const onboardingService = async (data: {
    username: string;
    email: string;
    password: string;
  }) => {
    const serviceUrl = `${import.meta.env.VITE_ONBOARDING_SERVICE_URL}/onboard`;
    try {
      const response = await fetch(serviceUrl, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return await response.json();
    } catch (error) {
      throw new Error(`Registration failed: ${error.message}`);
    }
  };
</script>

<style scoped>
  form {
    display: flex;
    flex-direction: column;
    max-width: 300px;
    margin: auto;
  }

  div {
    margin-bottom: 10px;
  }

  label {
    margin-bottom: 5px;
  }

  input {
    padding: 8px;
    font-size: 14px;
  }

  button {
    padding: 10px;
    font-size: 16px;
    cursor: pointer;
  }
</style>
