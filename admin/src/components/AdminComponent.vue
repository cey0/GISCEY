<template>
  <div class="admin-container">
    <div class="card">
      <h2>CRUD Restaurant</h2>
      <form @submit.prevent="submitRestaurant">
        <div>
          <label for="restaurantName">Nama Restoran:</label>
          <input type="text" v-model="restaurant.name" id="restaurantName" required />
        </div>
        <div>
          <label for="restaurantType">Tipe Restoran:</label>
          <select v-model="restaurant.type_id" id="restaurantType" required>
            <option v-for="type in types" :key="type.id" :value="type.id">{{ type.nama }}</option>
          </select>
        </div>
        <div>
          <label for="restaurantDescription">Deskripsi:</label>
          <textarea v-model="restaurant.deskripsi" id="restaurantDescription" required></textarea>
        </div>
        <div>
          <label for="restaurantLatitude">Latitude:</label>
          <input type="number" step="any" v-model="restaurant.latitude" id="restaurantLatitude" required />
        </div>
        <div>
          <label for="restaurantLongitude">Longitude:</label>
          <input type="number" step="any" v-model="restaurant.longitude" id="restaurantLongitude" required />
        </div>
        <button type="submit">Simpan</button>
      </form>
      <table>
        <thead>
          <tr>
            <th>Nama</th>
            <th>Tipe</th>
            <th>Deskripsi</th>
            <th>Latitude</th>
            <th>Longitude</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="restaurant in restaurants" :key="restaurant.id">
            <td>{{ restaurant.name }}</td>
            <td>{{ getTypeName(restaurant.type_id) }}</td>
            <td>{{ restaurant.deskripsi }}</td>
            <td>{{ restaurant.latitude }}</td>
            <td>{{ restaurant.longitude }}</td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="card">
      <h2>CRUD Type</h2>
      <form @submit.prevent="submitType" enctype="multipart/form-data">
        <div>
          <label for="typeName">Nama Tipe:</label>
          <input type="text" v-model="type.nama" id="typeName" required />
        </div>
        <div>
          <label for="typePhoto">Foto:</label>
          <input type="file" @change="handleFileUpload" id="typePhoto" required />
        </div>
        <button type="submit">Simpan</button>
      </form>
      <table>
        <thead>
          <tr>
            <th>Nama</th>
            <th>Foto</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="type in types" :key="type.id">
            <td>{{ type.nama }}</td>
            <td><img :src="type.foto_url" alt="Foto Tipe" width="50" /></td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      restaurant: {
        name: '',
        type_id: '',
        deskripsi: '',
        latitude: '',
        longitude: ''
      },
      type: {
        nama: '',
        foto: null
      },
      types: [], // Array untuk menyimpan data tipe restoran
      restaurants: [] // Array untuk menyimpan data restoran
    };
  },
  methods: {
    async submitRestaurant() {
      try {
        const response = await axios.post('http://localhost:8080/api/restaurants', this.restaurant);
        console.log('Restaurant data:', response.data);
        // Reset form
        this.restaurant = {
          name: '',
          type_id: '',
          deskripsi: '',
          latitude: '',
          longitude: ''
        };
        this.fetchRestaurants(); // Refresh data
      } catch (error) {
        console.error('Error submitting restaurant:', error);
      }
    },
    async submitType() {
      try {
        const formData = new FormData();
        formData.append('nama', this.type.nama);
        formData.append('foto', this.type.foto);

        const response = await axios.post('http://localhost:8080/api/types', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        });
        console.log('Type data:', response.data);
        // Reset form
        this.type = {
          nama: '',
          foto: null
        };
        this.fetchTypes(); // Refresh data
      } catch (error) {
        console.error('Error submitting type:', error);
      }
    },
    handleFileUpload(event) {
      this.type.foto = event.target.files[0];
    },
    async fetchTypes() {
      try {
        const response = await axios.get('http://localhost:8080/api/types');
        this.types = response.data.success;
      } catch (error) {
        console.error('Error fetching types:', error);
      }
    },
    async fetchRestaurants() {
      try {
        const response = await axios.get('http://localhost:8080/api/restaurants');
        this.restaurants = response.data.success;
      } catch (error) {
        console.error('Error fetching restaurants:', error);
      }
    },
    getTypeName(type_id) {
      const type = this.types.find(t => t.id === type_id);
      return type ? type.nama : 'Unknown';
    }
  },
  mounted() {
    this.fetchTypes(); // Ambil data tipe saat komponen dimuat
    this.fetchRestaurants(); // Ambil data restoran saat komponen dimuat
  }
};
</script>

<style scoped>
.admin-container {
  display: flex;
  gap: 20px;
}

.card {
  border: 1px solid #ccc;
  padding: 20px;
  border-radius: 8px;
  width: 45%;
  overflow: auto; /* Tambahkan ini */
}

.card h2 {
  margin-top: 0;
}

.card form div {
  margin-bottom: 10px;
}

.card form label {
  display: block;
  margin-bottom: 5px;
}

.card form input,
.card form select,
.card form textarea {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
}

.card form button {
  padding: 10px 15px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.card form button:hover {
  background-color: #0056b3;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

table, th, td {
  border: 1px solid #ccc;
}

th, td {
  padding: 10px;
  text-align: left;
}

img {
  max-width: 100%;
  height: auto;
}
</style>
