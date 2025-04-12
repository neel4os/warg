<template>
  <div class="q-pa-md fit row justify-center q-gutter-md;">
    <q-card class="q-pa-md; q-gutter-md; q-ma-md; col-8">
      <q-card-section>
        <div class="text-h6 q-py-md">Account Registration</div>
        <div class="text-h7 q-py-md">Enter your Account name and user data</div>
      </q-card-section>
      <q-card-section>
        <q-input
          color="grey-3"
          label-color="orange"
          class="q-w-50 shadow-2"
          outlined
          v-model="account_name"
          label="Account Name"
        >
        </q-input>
        <div class="row q-mt-md">
          <q-input
            color="grey-3"
            label-color="orange"
            class="q-w-50 shadow-2 col-6"
            outlined
            v-model="first_name"
            label="First Name"
          >
          </q-input>
          <q-input
            color="grey-3"
            label-color="orange"
            class="q-w-50 shadow-2 col-6 q-pl-sm"
            outlined
            v-model="last_name"
            label="Last Name"
          >
          </q-input>
        </div>
        <div class="q-pt-md">
          <q-input
            color="grey-3"
            label-color="orange"
            class="q-w-50 shadow-2"
            outlined
            v-model="email"
            label="Email"
          >
          </q-input>
        </div>
      </q-card-section>
      <q-card-actions align="around">
        <q-btn
          no-caps
          outline
          padding="xs lg"
          color="grey"
          @click="emitterSignin"
          >Sign In</q-btn
        >
        <q-btn no-caps padding="xs lg" color="primary" @click="submitForm"
          >Register</q-btn
        >
      </q-card-actions>
    </q-card>
  </div>
</template>

<script lang="ts" setup>
  import { ref, defineEmits } from "vue";
  import { useQuasar } from "quasar";
  const $q = useQuasar();
  const account_name = ref("");
  const first_name = ref("");
  const last_name = ref("");
  const email = ref("");
  const buildPayload = () => {
    return {
      account_name: account_name.value,
      first_name: first_name.value,
      last_name: last_name.value,
      email: email.value,
    };
  };
  const emit = defineEmits(["register"]);
  const submitForm = async () => {
    console.log(buildPayload());
    try {
      const response = await onboardingService(buildPayload());
      console.log("Registration successful:", response);
    } catch (error) {
      console.error("Registration failed:", error);
    }
  };
  const onboardingService = async (data: {
    account_name: string;
    first_name: string;
    last_name: string;
    email: string;
  }) => {
    const serviceUrl = `${location.protocol}://${location.host}/onboard`;
    console.log(serviceUrl);
    try {
      const response = await fetch(serviceUrl, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      if (!response.ok) {
        throw new Error(response.statusText);
      }
      return await response.json();
    } catch (error) {
      if (error instanceof Error) {
        showNotification($q, "Registration failed", error.message, "negative");
        console.log(`Registration failed: ${error.message}`);
      } else {
        showNotification(
          $q,
          "Registration failed",
          "Unknown Error",
          "negative"
        );
        console.log("Registration failed: An unknown error occurred.");
      }
    }
  };

  function emitterSignin() {
    emit("register", false);
  }
</script>
<style></style>
