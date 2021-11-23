<template>
  <div class="login">
    <v-container fluid class="fill-height pt-5">
      <v-row align="center" justify="center" class="d-flex flex-column">
        <img src="../assets/logos.svg" alt="logo" />
        <v-col cols="12" sm="6" md="4" class="mt-5">
          <v-card class="elevation-0 rounded-xl pa-2">
            <v-card-text>
              <h3 class="text-center display-1 font-weight-bold black--text">
                Welcome Back!
              </h3>
              <p class="text-center body-2 mt-3">
                Masukan number handphone/email dan passwordmu untuk masuk ke
                halaman dashboard
              </p>
              <v-form>
                <v-col cols="12" class="pa-0 text-left">
                  <label>Phone Number Hp/Email</label>
                </v-col>
                <v-text-field
                  name="hp"
                  type="text"
                  color="#fb6c37"
                  dense
                  outlined
                />
                <v-col cols="12" class="pa-0 text-left">
                  <label>Password</label>
                </v-col>
                <v-text-field
                  id="password"
                  name="password"
                  type="password"
                  color="#fb6c37"
                  dense
                  outlined
                />
                <v-checkbox class="mt-0" label="Remember Me"></v-checkbox>
                <div>
                  <v-btn block color="#fb6c37" class="white--text">Login</v-btn>
                </div>
              </v-form>
              <div class="mt-5 secondary--text">
                Kamu belum punya akun?
                <router-link
                  class="text-decoration-none secondary--text"
                  to="/register"
                >
                  <span style="color: #fb6c37">Register disini</span>
                </router-link>
              </div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
  export default {
    name: "Login",
  };
</script>

<style>
  * {
    margin: 0;
    padding: 0;
  }
  .login {
    background: url("../assets/back.svg");
    background-repeat: no-repeat;
    background-position: center center;
    background-size: cover;
    height: 100vh;
  }
</style>
