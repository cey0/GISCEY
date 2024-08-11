<template>
  <div>
    <input v-model="searchQuery" @input="searchRestaurants" placeholder="Search restaurants" />
    <div id="map" style="height: 500px;"></div>
  </div>
</template>

<script>
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';

export default {
  name: 'MapComponent',
  setup() {
    const map = ref(null);
    const restaurants = ref([]);
    const markers = ref([]);
    const searchQuery = ref('');
    const userLocation = ref(null);
    const route = useRoute();
    const typeId = ref(route.query.typeId);

    onMounted(() => {
      initializeMap();
      fetchRestaurants();
    });

    const initializeMap = () => {
      map.value = L.map('map').setView([-6.394124940533345, 106.79166622329677], 13);

      L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; OpenStreetMap contributors',
      }).addTo(map.value);

      // Get user location
      if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(
          (position) => {
            userLocation.value = L.latLng(position.coords.latitude, position.coords.longitude);
            map.value.setView(userLocation.value, 13);
            L.marker(userLocation.value).addTo(map.value).bindPopup('You are here').openPopup();
            // Fetch restaurants again to update distances
            fetchRestaurants();
          },
          (error) => {
            console.error('Error getting user location:', error);
          }
        );
      }
    };

    const fetchRestaurants = async () => {
      try {
        const response = await fetch('http://localhost:8080/restaurants');
        const data = await response.json();

        // Clear previous markers
        markers.value.forEach(marker => marker.remove());
        markers.value = [];

        if (data && Array.isArray(data.succes)) {
          restaurants.value = data.succes;

          // Filter by typeId if provided
          if (typeId.value) {
            restaurants.value = restaurants.value.filter(restaurant => restaurant.type_id === parseInt(typeId.value));
          }

          if (restaurants.value.length === 0) {
            showNotification('No restaurants found for this type.');
          }

          restaurants.value.forEach((restaurant) => {
            const latLng = L.latLng(restaurant.latitude, restaurant.longitude);
            const marker = L.marker(latLng).addTo(map.value);

            // Check if userLocation is available before calculating distance
            let distanceText = '';
            if (userLocation.value) {
              const distance = userLocation.value.distanceTo(latLng) / 1000; // Distance in kilometers
              distanceText = `<br>Distance: ${distance.toFixed(2)} km`;
            }

            marker.bindPopup(`<b>${restaurant.name}</b><br>${restaurant.deskripsi}${distanceText}`).openPopup();
            markers.value.push(marker);
          });
        } else {
          console.error('Unexpected data format:', data);
        }
      } catch (error) {
        console.error('Error fetching restaurants:', error);
      }
    };

    const searchRestaurants = () => {
      const query = searchQuery.value.toLowerCase();
      markers.value.forEach(marker => {
        const popupContent = marker.getPopup().getContent().toLowerCase();
        if (popupContent.includes(query)) {
          marker.addTo(map.value);
        } else {
          map.value.removeLayer(marker);
        }
      });
    };

    const showNotification = (message) => {
      const notification = document.createElement('div');
      notification.textContent = message;
      notification.style.position = 'fixed';
      notification.style.bottom = '20px';
      notification.style.right = '20px';
      notification.style.backgroundColor = '#333';
      notification.style.color = '#fff';
      notification.style.padding = '10px';
      notification.style.borderRadius = '5px';
      notification.style.zIndex = '1000';
      notification.style.boxShadow = '0 2px 5px rgba(0,0,0,0.3)';
      document.body.appendChild(notification);

      setTimeout(() => {
        document.body.removeChild(notification);
      }, 3000);
    };

    return {
      searchQuery,
      searchRestaurants
    };
  }
};
</script>

<style>
#map {
  width: 100%;
  height: 100%;
}
input {
  margin: 10px;
  padding: 5px;
}
</style>
