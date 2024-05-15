<!-- login.vue -->
<template>
  <div class="auth-wrapper auth-v1">
    <div class="auth-inner">
      <v-card class="auth-card">
        <!-- logo -->
        <v-card-title class="d-flex align-center justify-center py-7">
          <router-link :to="{ name: 'pages-login' }" class="d-flex align-center">
            <v-img
              :src="require('@/assets/images/logos/logo-KBT.png').default"
              max-height="120px"
              max-width="120px"
              alt="logo"
              contain
              class="me-3"
            ></v-img>
          </router-link>
        </v-card-title>

        <!-- title -->
        <v-card-text>
          <p class="text-1xl font-weight-semibold text--primary mb-2 text-center">
            Welcome to E-KBT
          </p>
          <p class="mb-2 text-center">Kepuasan Penumpang adalah Kebahagian Kami</p>
        </v-card-text>
        <v-card-text>
          <v-form ref="form" @submit.prevent="login">
            <v-text-field
              v-model="email"
              outlined
              label="Email"
              placeholder="john@example.com"
              hide-details
              :error-messages="errors.email"
              class="mb-1"
            ></v-text-field>
            <label class="text-danger" v-if="errors.email" type="error" dismissible>
              Email tidak boleh kosong!
            </label>

            <v-text-field
              v-model="password"
              outlined
              :type="isPasswordVisible ? 'text' : 'password'"
              label="Password"
              class="mb-1 mt-3"
              :error-messages="errors.password"
              placeholder="********"
              :append-icon="
                isPasswordVisible ? icons.mdiEyeOffOutline : icons.mdiEyeOutline
              "
              hide-details
              @click:append="isPasswordVisible = !isPasswordVisible"
            ></v-text-field>
            <label class="text-danger" v-if="errors.password" type="error" dismissible>
              Password tidak boleh kosong!
            </label>

            <div class="d-flex align-center justify-space-between flex-wrap">
              <v-checkbox
                label="Ingat saya"
                hide-details
                class="me-3 mt-1"
                v-model="rememberMe"
              ></v-checkbox>

              <!-- forgot link -->
              <router-link :to="{ name: 'pages-forgot-password' }">
                Lupa Password?
              </router-link>
            </div>

            <v-btn type="submit" block color="primary" class="mt-6" :loading="isLoading">
              <template v-if="!isLoading"> Login </template>
              <template v-if="isLoading">
                <v-progress-circular
                  indeterminate
                  size="24"
                  color="white"
                ></v-progress-circular>
              </template>
            </v-btn>
          </v-form>
        </v-card-text>

        <!-- create new account  -->
        <v-card-text class="d-flex align-center justify-center flex-wrap mt-2">
          <span class="me-2"> Belum memiliki akun?</span>
          <router-link :to="{ name: 'pages-register' }"> Daftar sekarang</router-link>
        </v-card-text>
      </v-card>
    </div>
  </div>
</template>

<script>
import { mdiEyeOutline, mdiEyeOffOutline } from "@mdi/js";
import { ref } from "@vue/composition-api";
import axios from "axios";
import Swal from "sweetalert2";
import { mapActions } from 'vuex';

export default {
  setup() {
    const isPasswordVisible = ref(false);
    const email = ref("");
    const password = ref("");
    const rememberMe = ref(false);

    return {
      isPasswordVisible,
      email,
      password,
      rememberMe,
      icons: {
        mdiEyeOutline,
        mdiEyeOffOutline,
      },
    };
  },
  data() {
    return {
      errors: {},
      isLoading: false,
    };
  },
  methods: {
    ...mapActions(['setUserRole']), // Menggunakan setUserRole dari Vuex actions
    login() {
      this.isLoading = true;

      const data = {
        email: this.email,
        password: this.password,
        remember_me: this.rememberMe,
      };

      axios
        .post("http://localhost:8081/login", data)
        .then((response) => {
          if (response.data.message === "Login berhasil") {
            const token = response.data.access_token;
            localStorage.setItem("access_token", token);
            localStorage.setItem("expires_at", response.data.expires_at);
            var email = localStorage.getItem('email');

// Menampilkan email pada console log
            console.log("Role ID:", response.data.role_id);
            // Simpan email ke dalam localStorage
            localStorage.setItem('userEmail', response.data.email);
            localStorage.setItem('userRole', response.data.role_id);

            this.$store.dispatch('setUserRole', response.data.role_id);

            const userRole = localStorage.getItem('userRole');
            const userEmail = localStorage.getItem('userEmail');
            console.log('Email yang disimpan:', userEmail);
            console.log('ROLE_ID = ', response.data.role_id);
            this.$router.push("/dashboard");
          } else {
            Swal.fire({
              icon: "error",
              title: "Login failed",
              text: "Incorrect email or password!",
              confirmButtonText: "Ok",
              confirmButtonColor: "#d33",
            });
          }
        })
        .catch((error) => {
          Swal.fire({
            icon: "error",
            title: "Login failed",
            text: "An error occurred while logging in. Please try again later.",
            confirmButtonText: "Ok",
            confirmButtonColor: "#d33",
          });
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
    checkSavedUserData() {
      const userDataString = localStorage.getItem("user_data");
      if (userDataString) {
        const userData = JSON.parse(userDataString);
        const expirationDate = new Date(userData.expires_at);
        if (expirationDate > new Date()) {
          this.$store.dispatch("updateUserRole", userData.access_token);
          this.$router.push("/dashboard");
        } else {
          localStorage.removeItem("user_data");
        }
      }
    },
  },
  mounted() {
    this.checkSavedUserData();
  },
};
</script>

<style lang="scss">
@import "~@resources/sass/preset/pages/auth.scss";
</style>
