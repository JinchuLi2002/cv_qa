<template>
    <div>
      <h1>Received Number: {{ number }}</h1>
      <button @click="fetchNumber">Fetch Number</button>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        number: null
      };
    },
    methods: {
      async fetchNumber() {
        try {
          const response = await axios.get('http://localhost:5001/api/send');
          this.number = response.data.data;
        } catch (error) {
          console.error('Error fetching the number:', error);
        }
      }
    }
  };
  </script>
  