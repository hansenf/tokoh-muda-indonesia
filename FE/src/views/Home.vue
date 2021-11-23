<template>
  <div class="home">
    <v-container>
      <v-row align="center" justify="center">
        <v-col xs="12" md="6">
          <h1 class="left--text display-3 font-weight-bold head">
            Komunitas <span style="color: #ffc700">Belajar</span> Bagi Pemuda
            dan Mahasiswa <span style="color: #fb6c37">Indonesia</span>
          </h1>
          <h5 class="mt-5 font-weight-regular">
            Kami membantu meningkatkan self-development anda untuk mengurangi
            jumlah pengangguran pemuda di indonesia. Ayok bergabung bersama kami
            agar anda tidak menganggur lagi.
          </h5>
          <v-btn
            color="#FB6C37"
            rounded
            class="white--text mt-5 py-5 px-8 text-capitalize"
            @click="goPage('/register')"
          >
            Daftar Disini
          </v-btn>
          <v-row class="mt-5">
            <v-col md="3">
              <h4>10K+</h4>
              <p>Member Count</p>
            </v-col>
            <v-col md="3">
              <h4>200+</h4>
              <p>University</p>
            </v-col>
            <v-col md="3">
              <h4>120+</h4>
              <p>Practice Kit</p>
            </v-col>
            <v-col md="3">
              <h4>130+</h4>
              <p>Challenges</p>
            </v-col>
          </v-row>
        </v-col>
        <v-col xs="12" md="6">
          <!-- <img src="../assets/mhs.svg" alt="" height="100%" width="100%" /> -->
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
  // @ is an alias to /src

  export default {
    name: "Home",
    methods: {
      goPage(url) {
        this.$router.push(url);
      },
    },
  };
</script>
<style lang="scss" scoped>
  .home {
    padding-top: 7rem;
    // height: 100vh;
    height: 700px;
    width: 100vw;
    background: url("../assets/kabeh.svg");
    background-repeat: no-repeat;
    background-position: center center;
    background-size: cover;

    .head {
      width: 110%;
      line-height: 70px;
    }
    h4 {
      color: #ff5107;
      font-size: 40px;
    }
    p {
      font-size: 16px;
    }
  }
</style>
