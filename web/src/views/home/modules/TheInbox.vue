<template>
  <div class="mx-auto px-2 sm:px-5 my-4">
    <div v-for="item in items" :key="item.id">
      <TheInboxCard :inbox="item" />
    </div>
    <div v-if="hasMore" @click="loadMore">加载更多</div>
  </div>
</template>
<script lang="ts" setup>
import { onMounted, onUnmounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useInboxStore } from '@/stores'
import { fetchMarkInboxRead } from '@/service/api'
import TheInboxCard from '@/components/advanced/TheInboxCard.vue'

const inboxStore = useInboxStore()
const { loadMore } = inboxStore
const { items, hasMore } = storeToRefs(inboxStore)

let timer: ReturnType<typeof setInterval>

onMounted(async () => {
  // 用户停留超过 1 秒则更新消息为已读
  timer = setInterval(() => {
    if (items.value.length > 0) {
      items.value.forEach((item) => {
        if (!item.read) {
          fetchMarkInboxRead(item.id).then(() => {
            item.read = true
          })
        }
      })
    }
  }, 1500)
})

onUnmounted(() => {
  clearInterval(timer)
})
</script>
<style scoped></style>
