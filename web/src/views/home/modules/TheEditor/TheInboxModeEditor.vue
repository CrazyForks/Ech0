<template>
  <div>
    <h2 class="text-[var(--text-color-500)] font-bold mb-1">收件箱模式</h2>

    <!-- 当前共有 -->
    <div class="text-[var(--text-color-400)] my-2 font-serif">
      <p>当前共有 {{ items.length }} 条消息</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useInboxStore } from '@/stores'

const inboxStore = useInboxStore()
const { items } = storeToRefs(inboxStore)
</script>
