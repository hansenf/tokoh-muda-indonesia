<template>
  <v-app>
    <div id="app">
      <!-- <div id="nav">
        <router-link to="/">Home</router-link> |
        <router-link to="/about">About</router-link> |
        <router-link to="/register">Register</router-link>
      </div> -->
      <router-view />
      <v-app-bar
        app
        absolute
        elevation="0"
        color="#FEF7EC"
        class="pt-3"
        v-if="!['Login', 'Register'].includes($route.name)"
      >
        <div class="d-flex container">
          <v-col cols="6" md="3" align-self="center">
            <img src="./assets/logos.svg" alt="logo" />
          </v-col>
          <v-col md="6" align-self="center" class="d-none d-md-flex">
            <v-tabs centered color="secondary" background-color="transparent">
              <v-tabs-slider color="yellow"></v-tabs-slider>
              <v-tab
                v-for="(item, index) in menus"
                :key="`menu-${index}`"
                :to="item.route"
                class="px-7"
              >
                {{ item.title }}
              </v-tab>
            </v-tabs>
          </v-col>
          <v-col cols="6" md="3" align-self="center" class="py-0 d-block px-0">
            <v-btn
              color="#FB6C37"
              class="float-right d-none d-md-inline-block px-7 py-2"
              outlined
              rounded
              @click="goPage('/login')"
            >
              Login
            </v-btn>
          </v-col>
        </div>
      </v-app-bar>
    </div>
  </v-app>
</template>
<script>
  export default {
    data: () => ({
      menus: [
        { title: "Home", icon: "mdi-home", route: "/" },
        { title: "About Us", icon: "mdi-account-group", route: "/about" },
        { title: "Contact Us", icon: "mdi-note", route: "/blogs" },
      ],
    }),
    props: {
      source: String,
    },
    methods: {
      goPage(url) {
        this.$router.push(url);
      },
    },
  };
</script>

<style lang="scss">
  #app {
    font-family: "Poppins", sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #2c3e50;
    // background: url("./assets/back.svg");
    // background-repeat: no-repeat;
    // background-position: center center;
    // background-size: cover;
  }

  #nav {
    padding: 30px;

    a {
      font-weight: bold;
      color: #2c3e50;

      &.router-link-exact-active {
        color: #42b983;
      }
    }
  }
</style>
