<template>
  <div class="container">
    <div class="row mt-4">
      <div class="col-md-2">
        <img
          v-if="ready"
          class="img-fluid img-thumbnail"
          :src="book?.formats['image/jpeg']"
          alt="cover"
        />
      </div>
      <div class="col-md-10">
        <template v-if="ready">
          <h3 class="mt-3">
            {{ book?.title }}
          </h3>
          <hr />
          <p>
            <strong> Author: </strong>
            {{ book?.authors[0].name }}
            <br />
            <strong> Subjects: </strong>
            {{ book?.subjects.join(', ') }}
          </p>
        </template>
        <p v-else>Loading ...</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onBeforeMount } from 'vue'
import { useRoute } from 'vue-router'
import type { Book } from 'src/services/api'

const API_URL: string = import.meta.env.VITE_API_URL

const route = useRoute()
const bookId = route.params.id
const ready = ref<Boolean>(false)
const book = ref<Book>()

const emit = defineEmits(['error'])

onBeforeMount(() => {
  fetch(`${API_URL}/books/${bookId}`)
    .then((response) => response.json())
    .then((data) => {
      if (data.error) {
        emit('error', data.message)
      } else {
        book.value = data.data
        ready.value = true
      }
    })
})
</script>
