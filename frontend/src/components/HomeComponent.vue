<template>
  <div class="flex justify-center items-center h-full">
    <div
      class="grid gap-4 w-full h-full"
      :class="{
        'grid-cols-1': types.length === 1,
        'grid-cols-2': types.length === 2,
        'grid-cols-3': types.length === 3,
        'grid-cols-4': types.length === 4,
        'grid-cols-5': types.length === 5,
      }"
    >
      <div
        v-for="type in types"
        :key="type.ID"
        class="box bg-gray-100 p-4 rounded-lg shadow-md text-center flex flex-col justify-center items-center h-full cursor-pointer"
        @click="goToMap(type.ID)"
      >
        <img
          :src="getImagePath(type.foto)"
          alt="Type Image"
          class="w-24 h-24 object-cover "
        />
        <h3 class="mt-2 text-xl font-bold">{{ type.nama }}</h3>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import { useRouter } from 'vue-router';

export default {
  name: "TypeBoxes",
  setup() {
    const router = useRouter();
    return { router };
  },
  data() {
    return {
      types: [],
    };
  },
  mounted() {
    this.fetchTypes();
  },
  methods: {
    async fetchTypes() {
      try {
        const response = await axios.get("http://localhost:8080/types");
        this.types = response.data.success;
      } catch (error) {
        console.error("Error fetching types:", error);
      }
    },
    getImagePath(filename) {
      return `http://localhost:8080/fotoTypes/${filename}`;
    },
    goToMap(typeId) {
    this.router.push({ name: 'maps', query: { typeId } });
  },
  },
};
</script>

<style scoped>
.grid > div {
  height: 100%; /* Ensure each grid item takes full height of the grid */
}
.box {
  background-color: #7B66FF;
}
</style>
