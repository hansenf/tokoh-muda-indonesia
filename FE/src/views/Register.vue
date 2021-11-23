<template>
  <div class="register">
    <v-container fluid class="fill-height pt-5">
      <v-row align="center" justify="center" class="d-flex flex-column">
        <img src="../assets/logos.svg" alt="logo" />
        <v-col cols="12" sm="6" md="4" class="mt-5 rounded-xl pa-2">
          <v-card class="elevation-0">
            <v-card-text>
              <h3 class="text-center display-1 font-weight-bold black--text">
                Ayo Daftar Sekarang
              </h3>
              <p class="text-center body-2 mt-3">
                Daftarkan dirimu dan bergabung dengan kami
              </p>
              <v-form>
                <v-col cols="12" class="pa-0 text-left">
                  <label>Username</label>
                </v-col>
                <v-text-field
                  name="username"
                  type="text"
                  color="#fb6c37"
                  outlined
                  dense
                  height="20px"
                />
                <v-col cols="12" class="pa-0 text-left">
                  <label>Phone Number Hp/Email</label>
                </v-col>
                <v-text-field
                  name="hp"
                  type="text"
                  color="#fb6c37"
                  dense
                  outlined
                  height="20px"
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
                  height="20px"
                />
                <v-col cols="12" class="pa-0 text-left">
                  <label>Confirm Password</label>
                </v-col>

                <v-text-field
                  id="password"
                  name="confirm password"
                  type="password"
                  color="#fb6c37"
                  dense
                  outlined
                />

                <!-- <p class="teal--text text-left">Password cocok</p> -->
                <v-checkbox
                  v-model="checkbox"
                  :rules="[(v) => !!v || 'You must agree to continue!']"
                  label="I accept Term & Condition and Privacy Policies"
                  required
                  class="subtitle-1 mt-0"
                ></v-checkbox>
                <div>
                  <v-btn block color="#fb6c37" class="white--text"
                    >Register</v-btn
                  >
                </div>
              </v-form>
              <div class="mt-5 secondary--text">
                Kamu sudah punya akun?
                <router-link
                  class="text-decoration-none secondary--text"
                  to="/login"
                >
                  <span style="color: #fb6c37; text-decoration: none"
                    >Login disini</span
                  >
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
    name: "Register",
  };
</script>

<style>
  * {
    margin: 0;
    padding: 0;
  }
  .register {
    background: url("../assets/back.svg");
    background-repeat: no-repeat;
    background-position: center center;
    background-size: cover;
    height: 100%;
  }
</style>
