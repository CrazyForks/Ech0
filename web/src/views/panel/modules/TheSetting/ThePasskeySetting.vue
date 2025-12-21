<template>
  <PanelCard>
    <div class="w-full">
      <div class="flex flex-row items-center justify-between mb-3">
        <h1 class="text-[var(--text-color-600)] font-bold text-lg">Passkey（多设备）</h1>
      </div>

      <div class="text-[var(--text-color-next-500)] text-sm mb-3">
        使用 Passkey（WebAuthn）可在不同设备上无密码登录；删除某设备后，该设备将无法再登录。
      </div>

      <!-- 绑定 -->
      <div class="flex flex-row items-center justify-start gap-2 mb-4">
        <BaseInput v-model="newDeviceName" type="text" placeholder="设备名称（可选）" class="w-full py-1!" />
        <BaseButton
          class="rounded-md px-3 w-20 h-9"
          :disabled="busy || !supported"
          @click="handleBind"
        >
          绑定
        </BaseButton>
      </div>

      <div v-if="!supported" class="text-[var(--text-color-next-500)] text-sm mb-3">
        当前浏览器不支持 Passkey / WebAuthn。
      </div>

      <!-- 多设备管理 -->

      <div class="text-[var(--text-color-next-500)] font-semibold mb-2">已绑定设备</div>
      <div v-if="devices.length === 0" class="text-[var(--text-color-next-500)] text-sm">
        暂无设备
      </div>
      <div v-else class="mt-2 overflow-x-auto border border-[var(--border-color-300)] rounded-lg">
        <table class="min-w-full divide-y divide-[var(--divide-color-200)]">
          <thead>
            <tr class="bg-[var(--bg-color-50)] opacity-70">
              <th
                class="px-3 py-2 text-left text-sm font-semibold text-[var(--text-color-next-600)]"
              >
                设备名称
              </th>
              <th
                class="px-3 py-2 text-left text-sm font-semibold text-[var(--text-color-next-600)]"
              >
                AAGUID
              </th>
              <th
                class="px-3 py-2 text-left text-sm font-semibold text-[var(--text-color-next-600)]"
              >
                时间
              </th>
              <th
                class="px-3 py-2 text-right text-sm font-semibold text-[var(--text-color-next-600)]"
              >
                操作
              </th>
            </tr>
          </thead>
          <tbody class="divide-y divide-[var(--divide-color-100)] text-nowrap">
            <tr v-for="d in devices" :key="d.id">
              <td class="px-3 py-2 text-sm text-[var(--text-color-next-700)] font-semibold">
                {{ d.device_name || 'Passkey' }}
              </td>
              <td class="px-3 py-2 text-sm text-[var(--text-color-next-500)]">
                {{ d.aaguid || '未知' }}
              </td>
              <td class="px-3 py-2 text-xs text-[var(--text-color-next-500)]">
                <div>最近使用：{{ formatTime(d.last_used_at) }}</div>
                <div>创建：{{ formatTime(d.created_at) }}</div>
              </td>
              <td class="px-3 py-2 text-right">
                <div class="flex flex-row items-center justify-end gap-2">
                  <BaseButton
                    class="rounded-md px-2 h-8 text-xs"
                    :disabled="busy"
                    @click="promptRename(d)"
                  >
                    改名
                  </BaseButton>
                  <BaseButton
                    class="rounded-md px-2 h-8 text-xs"
                    :disabled="busy"
                    @click="handleDelete(d.id)"
                  >
                    删除
                  </BaseButton>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </PanelCard>
</template>
<script setup lang="ts">
import { onMounted, ref } from 'vue'
import PanelCard from '@/layout/PanelCard.vue'
import BaseInput from '@/components/common/BaseInput.vue'
import BaseButton from '@/components/common/BaseButton.vue'
import {
  fetchDeletePasskeyDevice,
  fetchPasskeyDevices,
  fetchPasskeyRegisterBegin,
  fetchPasskeyRegisterFinish,
  fetchUpdatePasskeyDeviceName,
} from '@/service/api'
import { theToast } from '@/utils/toast'

const supported = !!(window.PublicKeyCredential && navigator.credentials)
const busy = ref(false)
const newDeviceName = ref<string>('My Passkey')
const devices = ref<App.Api.Auth.PasskeyDevice[]>([])

type Base64urlString = string

type CredentialDescriptorJSON = {
  type: PublicKeyCredentialType
  id: Base64urlString
  transports?: AuthenticatorTransport[]
}

type UserEntityJSON = {
  id: Base64urlString
  name: string
  displayName: string
}

type CreationOptionsJSON = Omit<PublicKeyCredentialCreationOptions, 'challenge' | 'user' | 'excludeCredentials'> & {
  challenge: Base64urlString
  user: UserEntityJSON
  excludeCredentials?: CredentialDescriptorJSON[]
}

function base64urlToUint8Array(input: string) {
  const base64 = input.replace(/-/g, '+').replace(/_/g, '/')
  const pad = base64.length % 4 === 0 ? '' : '='.repeat(4 - (base64.length % 4))
  const binary = atob(base64 + pad)
  const bytes = new Uint8Array(binary.length)
  for (let i = 0; i < binary.length; i++) bytes[i] = binary.charCodeAt(i)
  return bytes
}

function uint8ArrayToBase64url(bytes: ArrayBuffer | Uint8Array) {
  const u8 = bytes instanceof Uint8Array ? bytes : new Uint8Array(bytes)
  let binary = ''
  for (let i = 0; i < u8.length; i++) binary += String.fromCharCode(u8[i]!)
  const base64 = btoa(binary)
  return base64.replace(/\+/g, '-').replace(/\//g, '_').replace(/=+$/g, '')
}

function assertCreationOptionsJSON(raw: unknown): CreationOptionsJSON {
  if (!raw || typeof raw !== 'object') throw new Error('服务端返回的 publicKey 不合法')
  return raw as CreationOptionsJSON
}

function normalizeCreationOptions(raw: unknown): PublicKeyCredentialCreationOptions {
  const o = assertCreationOptionsJSON(raw)
  const { challenge, user, excludeCredentials, ...rest } = o
  const exclude = Array.isArray(excludeCredentials)
    ? excludeCredentials.map((c) => ({
        ...c,
        id: base64urlToUint8Array(c.id),
      }))
    : undefined

  return {
    ...rest,
    challenge: base64urlToUint8Array(challenge),
    user: {
      ...user,
      id: base64urlToUint8Array(user.id),
    },
    ...(exclude ? { excludeCredentials: exclude } : {}),
  }
}

function credentialToJSON(cred: PublicKeyCredential) {
  if (!cred) return null
  const obj: Record<string, unknown> = {
    id: cred.id,
    rawId: uint8ArrayToBase64url(cred.rawId),
    type: cred.type,
    clientExtensionResults: cred.getClientExtensionResults?.() ?? {},
  }

  const response: Record<string, unknown> = {}
  response.clientDataJSON = uint8ArrayToBase64url(cred.response.clientDataJSON)

  // 注册（attestation）
  if ('attestationObject' in cred.response) {
    const r = cred.response as AuthenticatorAttestationResponse
    response.attestationObject = uint8ArrayToBase64url(r.attestationObject)
  }

  // 登录（assertion）——这里暂时不会用到，但保持通用
  if ('authenticatorData' in cred.response) {
    const r = cred.response as AuthenticatorAssertionResponse
    response.authenticatorData = uint8ArrayToBase64url(r.authenticatorData)
    response.signature = uint8ArrayToBase64url(r.signature)
    if (r.userHandle && r.userHandle.byteLength > 0) {
      response.userHandle = uint8ArrayToBase64url(r.userHandle)
    }
  }

  obj.response = response
  return obj
}

function formatTime(v: string) {
  if (!v) return '暂无'
  const d = new Date(v)
  if (Number.isNaN(d.getTime())) return v
  return d.toLocaleString()
}

async function refresh() {
  const res = await fetchPasskeyDevices()
  if (res.code === 1) devices.value = res.data ?? []
}

async function handleBind() {
  if (!supported) return
  busy.value = true
  try {
    const begin = await fetchPasskeyRegisterBegin(newDeviceName.value || 'Passkey')
    if (begin.code !== 1) return

    const options = normalizeCreationOptions(begin.data.publicKey)
    const created = await navigator.credentials.create({ publicKey: options })
    if (!created) throw new Error('创建凭证失败')
    const cred = created as PublicKeyCredential

    const finish = await fetchPasskeyRegisterFinish(begin.data.nonce, credentialToJSON(cred))
    if (finish.code !== 1) return

    theToast.success('绑定成功')
    await refresh()
  } catch (e: unknown) {
    const msg = e instanceof Error ? e.message : '绑定失败'
    theToast.error(msg)
  } finally {
    busy.value = false
  }
}

async function handleDelete(id: number) {
  busy.value = true
  try {
    const res = await fetchDeletePasskeyDevice(id)
    if (res.code !== 1) return
    theToast.success('已删除')
    await refresh()
  } finally {
    busy.value = false
  }
}

async function promptRename(d: App.Api.Auth.PasskeyDevice) {
  const name = window.prompt('新的设备名称', d.device_name || 'Passkey')
  if (!name) return
  busy.value = true
  try {
    const res = await fetchUpdatePasskeyDeviceName(d.id, name)
    if (res.code !== 1) return
    theToast.success('已更新')
    await refresh()
  } finally {
    busy.value = false
  }
}

onMounted(() => {
  refresh()
})
</script>