<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商店用户ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="已删除:" prop="isDeleted">
          <el-switch v-model="formData.isDeleted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="创建时间:" prop="dateAdd">
          <el-date-picker v-model="formData.dateAdd" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="更新时间:" prop="dateUpdate">
          <el-date-picker v-model="formData.dateUpdate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="删除时间:" prop="dateDelete">
          <el-date-picker v-model="formData.dateDelete" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="模板详情id:" prop="logisticsId">
          <el-input v-model.number="formData.logisticsId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="模板详情id:" prop="detailId">
          <el-input v-model.number="formData.detailId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="provinceId字段:" prop="provinceId">
          <el-input v-model="formData.provinceId" :clearable="true"  placeholder="请输入provinceId字段" />
       </el-form-item>
        <el-form-item label="cityId字段:" prop="cityId">
          <el-input v-model="formData.cityId" :clearable="true"  placeholder="请输入cityId字段" />
       </el-form-item>
        <el-form-item label="regionId字段:" prop="regionId">
          <el-input v-model="formData.regionId" :clearable="true"  placeholder="请输入regionId字段" />
       </el-form-item>
        <el-form-item label="regionStr字段:" prop="regionStr">
          <el-input v-model="formData.regionStr" :clearable="true"  placeholder="请输入regionStr字段" />
       </el-form-item>
        <el-form-item label="addAmount字段:" prop="addAmount">
          <el-input-number v-model="formData.addAmount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="addNumber字段:" prop="addNumber">
          <el-input-number v-model="formData.addNumber" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="firstAmount字段:" prop="firstAmount">
          <el-input-number v-model="formData.firstAmount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="firstNumber字段:" prop="firstNumber">
          <el-input-number v-model="formData.firstNumber" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createBeeShopLogisticsException,
  updateBeeShopLogisticsException,
  findBeeShopLogisticsException
} from '@/api/bee/beeShopLogisticsException'

defineOptions({
    name: 'BeeShopLogisticsExceptionForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            id: undefined,
            userId: undefined,
            isDeleted: false,
            dateAdd: new Date(),
            dateUpdate: new Date(),
            dateDelete: new Date(),
            logisticsId: undefined,
            detailId: undefined,
            provinceId: '',
            cityId: '',
            regionId: '',
            regionStr: '',
            addAmount: 0,
            addNumber: 0,
            firstAmount: 0,
            firstNumber: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBeeShopLogisticsException({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createBeeShopLogisticsException(formData.value)
               break
             case 'update':
               res = await updateBeeShopLogisticsException(formData.value)
               break
             default:
               res = await createBeeShopLogisticsException(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
