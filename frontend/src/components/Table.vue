<template>
  <v-card>
    <v-card-title>
      Hit API List
      <v-spacer></v-spacer>
      <v-text-field
        v-model="search"
        append-icon="mdi-magnify"
        label="Search for Host or Method"
        single-line
        hide-details
      ></v-text-field>
    </v-card-title>
    <v-data-table   
      :headers="headers"
      :items="desserts"
      :search="search"
    >
    <template v-slot:item.time="{ item }">
      {{ intToDateTime(item.time)  }} 
    </template>
    </v-data-table>
  </v-card>
</template>

<script>
import axios from 'axios';

  export default {
    data () {
      return {
        search: '',
        headers: [
          {
            text: 'Method',
            align: 'start',
            value: 'method'
          },
          { text: 'Host', value: 'host' },
          { text: 'Datetime', value: 'time', sortable: false, filterable: false },
          { text: 'URL', value: 'url', sortable: false, filterable: false },
          { text: 'Response Code', value: 'response', sortable: false, filterable: false },
          { text: 'Referer', value: 'referer', sortable: false, filterable: false },
        ],
        desserts: [],
      }
    },mounted() {
            // this is the url of the back-end spitting out json
        axios.get("http://localhost:8090/api/test")
            .then(response => {
                console.log(response)
                this.desserts = response.data
        })
    }, 
    methods : {

        intToDateTime(time) {
            var d = new Date(parseInt(time) * 1000)

            var date = d.getFullYear() + "-" +
                        ("00" + (d.getMonth() + 1)).slice(-2) + "-" + 
                        ("00" + d.getDate()).slice(-2) + " " +
                        ("00" + d.getHours()).slice(-2) + ":" +
                        ("00" + d.getMinutes()).slice(-2) + ":" +
                        ("00" + d.getSeconds()).slice(-2);

            return date
        }
    }

  }
</script>