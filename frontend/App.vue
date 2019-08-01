<template>
  <el-container class="yadockeri-container" id="app">
    <el-header class="yadockeri-header" height="auto">
      <h1 class="title"><router-link to="/">Yadockeri</router-link></h1>
    </el-header>
    <el-main class="yadockeri-main">
      <router-view />
    </el-main>
  </el-container>
</template>

<script>
import 'ress/dist/ress.min.css'
export default {
  name: 'app',
}
</script>

<style lang="scss">
.yadockeri-container {
  min-height: 100vh;
}

.yadockeri-header {
  h1 {
    padding: 15px 0;
  }
}

.yadockeri-main {
  background-color: #f5f5f5;
}

body {
  font-family: 'Noto Sans', 'Noto Sans CJK JP', 'Takaoゴシック',
    'ヒラギノ角ゴ ProN W3', '-apple-system', 'BlinkMacSystemFont', 'Segoe UI',
    'Roboto', 'Helvetica Neue', 'Apple Color Emoji', 'Segoe UI Emoji',
    'Segoe UI Symbol', 'Noto Color Emoji', 'Noto Emoji';
}
</style>

<style lang="scss" scoped>
.title a {
  text-decoration: none;
  color: #409eff;
}
</style>
