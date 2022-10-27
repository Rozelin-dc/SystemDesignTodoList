<script lang="ts" setup>
import { AxiosError } from 'axios'
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useMe } from '@/store/me'
import { showErrorMessage } from '@/util/showErrorMessage'

const meStore = useMe()

const confirmDelete = async () => {
  await ElMessageBox.confirm(
    'アカウント削除は取り消せません。削除しますか',
    '確認',
    {
      confirmButtonText: 'はい',
      cancelButtonText: 'いいえ',
      type: 'warning'
    }
  )
    .then(() => deleteAccount())
    // eslint-disable-next-line @typescript-eslint/no-empty-function
    .catch(() => {})
}

const inputData = reactive({
  password: ''
})
const loading = ref(false)
const router = useRouter()
const deleteAccount = async () => {
  if (!meStore.getMe) {
    return
  }

  try {
    loading.value = true
    await meStore.deleteMe(inputData.password)
    ElMessage({
      message: 'アカウントを削除しました',
      type: 'success'
    })
    router.push({ name: 'Login' })
  } catch (e: any) {
    const err: AxiosError = e
    showErrorMessage(err)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="update-user-name-container">
    <div class="description">
      アカウント削除は取り消せません。アカウントを削除するとタスク情報なども削除され復元は出来ません。注意してください。
    </div>
    <el-form :model="inputData" label-position="top">
      <el-form-item prop="password" label="パスワードを入力して削除">
        <el-input v-model="inputData.password" type="password" show-password />
      </el-form-item>
    </el-form>

    <el-button type="danger" :loading="loading" @click="confirmDelete" round>
      削除
    </el-button>
  </div>
</template>

<style lang="scss" scoped>
.update-user-name-container {
  padding: 10px 5px;

  .description {
    margin-bottom: 10px;
    color: var(--el-text-color-regular);
  }
}
</style>
