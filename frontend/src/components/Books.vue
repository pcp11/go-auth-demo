<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">Books</h1>

        <hr />

        <div class="filters text-center">
          <!-- display filter options for genre -->
          <span
            class="filter me-1"
            :class="{ active: currentFilter === '0' }"
            @click="setFilter('0')"
          >
            ALL
          </span>
          <span
            class="filter me-1"
            :class="{ active: currentFilter === '7' }"
            @click="setFilter('7')"
          >
            CLASSIC
          </span>
          <span
            class="filter me-1"
            :class="{ active: currentFilter === '2' }"
            @click="setFilter('2')"
          >
            FANTASY
          </span>
          <span
            class="filter me-1"
            :class="{ active: currentFilter === '6' }"
            @click="setFilter('6')"
          >
            HORROR
          </span>
          <span
            class="filter me-1"
            :class="{ active: currentFilter === '4' }"
            @click="setFilter('4')"
          >
            THRILLER
          </span>
          <span
            class="filter me-1"
            :class="{ active: currentFilter === '1' }"
            @click="setFilter('1')"
          >
            SCIENCE FICTION
          </span>
        </div>
        <!-- display filter options for genre -->

        <hr />

        <div v-if="!bookStore.isLoading">
          <div class="card-group">
            <transition-group
              class="p-3 d-flex flex-row justify-content-center flex-wrap"
              tag="div"
              appear
              name="books"
            >
              <div v-for="b in bookStore.books" :key="b.id">
                <div
                  class="card m-4"
                  style="width: 22rem"
                  v-if="b.subjects.includes(currentFilter) || currentFilter === ''"
                >
                  <router-link :to="`/books/${b.id}`">
                    <img
                      :src="b.formats['image/jpeg']"
                      class="card-img-top"
                      :alt="`cover for ${b.title}`"
                    />
                  </router-link>
                  <div class="card-body text-center">
                    <h6 class="card-title">{{ b.title }}</h6>
                    <span class="book-author">{{ b.authors[0].name }}</span
                    ><br />
                    <small class="text-muted book-genre" v-for="(g, index) in b.subjects">
                      <em class="me-1"
                        >{{ g }}<template v-if="index !== b.subjects.length - 1">,</template></em
                      >
                    </small>
                  </div>
                </div>
              </div>
            </transition-group>
          </div>
        </div>
        <p v-else>Loading ...</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onBeforeMount } from 'vue'
import { useBookStore } from '@/stores/books'

const bookStore = useBookStore()
const currentFilter = ref<string>('')

const emit = defineEmits(['error'])

onBeforeMount(() => {
  bookStore.fetchBooks()
})

const setFilter = (filter: string) => {
  currentFilter.value = filter
}
</script>

<style scoped>
.filters {
  height: 2.5em;
}

.filter {
  padding: 6px 6px;
  cursor: pointer;
  border-radius: 6px;
  transition: all 0.35s;
  border: 1px solid silver;
}

.filter.active {
  background: lightgreen;
}

.filter:hover {
  background: lightgray;
}

.book-author,
.book-genre {
  font-size: 0.8em;
}

/* transition styles */
.books-move {
  transition: all 500ms ease-in-out 50ms;
}

.books-enter-active {
  transition: all 500ms ease-in-out;
}

.books-leave-active {
  transition: all 500ms ease-in;
}

.books-enter,
.books-leave-to {
  opacity: 0;
}
</style>
