<template>
  <div class="home">
    <Nav></Nav>
    <Center></Center>
  </div>
</template>

<script>
import Nav from "@/views/home/Nav.vue";
import Center from "@/views/home/Page.vue";
export default {
  name: "home",
  components: {
    Nav,
    Center,
  },
};
</script>
<style type="scss" scoped>
.home {
  display: flex;
  display: -webkit-flex;
  flex-direction: row;
  height: 100%;
  overflow: hidden;
}
</style>
