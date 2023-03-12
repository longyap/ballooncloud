<script setup>
import { RouterLink, RouterView } from 'vue-router'
import navbar from './components/Navbar.vue'
import footerterminal from './components/FooterTerminal.vue'
</script>

<template>
  <header>  
<navbar></navbar>
  </header>
<RouterView />
 <footer>
<footerterminal></footerterminal>
 </footer>  
</template>

<script>


</script>