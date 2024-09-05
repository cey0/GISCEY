<template>
  <div class="login-container">
    <h2>Login</h2>
    <form @submit.prevent="handleLogin">
      <div>
        <label for="username">Username:</label>
        <input type="text" v-model="username" id="username" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" v-model="password" id="password" required />
      </div>
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';
import { useRouter } from 'vue-router';

export default {
  data() {
    return {
      username: '',
      password: ''
    };
  },
  setup() {
    const router = useRouter();
    return { router };
  },
  methods: {
    async handleLogin() {
      try {
        const response = await axios.post('http://localhost:8080/api/login', {
          username: this.username,
          password: this.password
        });
        console.log('Login successful:', response.data);
        localStorage.setItem('token', response.data.token);
        this.router.push({ name: 'dash' });
      } catch (error) {
        console.error('Login failed:', error);
        // Tambahkan logika untuk menangani kesalahan login
      }
    }
  }
};
</script>

<style scoped>
.login-container {
  max-width: 300px;
  margin: 0 auto;
  padding: 1em;
  border: 1px solid #ccc;
  border-radius: 5px;
}
.login-container h2 {
  text-align: center;
}
.login-container div {
  margin-bottom: 1em;
}
.login-container label {
  display: block;
  margin-bottom: 0.5em;
}
.login-container input {
  width: 100%;
  padding: 0.5em;
  box-sizing: border-box;
}
.login-container button {
  width: 100%;
  padding: 0.5em;
  background-color: #007BFF;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
.login-container button:hover {
  background-color: #0056b3;
}
</style>