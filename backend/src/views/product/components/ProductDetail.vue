<template>
  <div class="createPost-container">
    <el-form ref="postForm" :model="postForm" :rules="rules" class="form-container">

      <sticky :class-name="'sub-navbar '+postForm.status">
        <el-button v-loading="loading" style="margin-left: 10px;" type="success" @click="submitForm">发布
        </el-button>
      </sticky>

      <div class="createPost-main-container">
        <el-row>
          <el-col :span="24">
            <el-row>
              <el-col :span="24">
                <el-form-item :label="$t('product.name')" label-width="80px"  >
                  <el-input v-model="postForm.name" :placeholder="$t('product.name')" type="text"  autosize />
                </el-form-item>
              </el-col>
            </el-row>

            <div class="postInfo-container">
              <el-row>
                <el-col :span="24">
                  <el-form-item :label="$t('product.number')" label-width="80px" >
                    <el-input v-model="postForm.number" :placeholder="$t('product.number')" type="text"  autosize />
                  </el-form-item>
                </el-col>
              </el-row>

              <el-row>
                <el-col :span="24">
                  <el-form-item :label="$t('product.category')" label-width="80px" >
                    <el-select v-model="postForm.category" :placeholder="$t('product.category')">
                      <el-option v-for="value in this.categories" :label="value" :value="value"></el-option>
                    </el-select>
                  </el-form-item>
                </el-col>
              </el-row>

              <el-row>
                <el-col :span="24">
                  <el-form-item :label="$t('product.price')" label-width="80px" >
                    <el-input-number v-model="postForm.price" :placeholder="$t('product.price')">
                      <template slot="append"> euros </template>
                    </el-input-number>
                  </el-form-item>
                </el-col>
              </el-row>

              <el-row>
                <el-col :span="24">
                  <el-form-item :label="$t('product.price_on_sale')" label-width="80px">
                    <el-input-number v-model="postForm.price_on_sale" placeholder="$t('product.price_on_sale')">
                      <template slot="append"> euros </template>
                    </el-input-number>
                  </el-form-item>
                </el-col>
              </el-row>

              <el-form-item :label="$t('product.onSale')" style="margin-bottom: 10px;" label-width="85px">
                <div class="box-item">
                  <el-switch v-model="postForm.on_sale"/>
                </div>
              </el-form-item>
            </div>
          </el-col>
        </el-row>

      </div>
    </el-form>

  </div>
</template>

<script>
import Tinymce from '@/components/Tinymce'
import Upload from '@/components/Upload/singleImage3'
import MDinput from '@/components/MDinput'
import Sticky from '@/components/Sticky' // 粘性header组件
import { validateURL } from '@/utils/validate'
import { fetchProduct, updateProduct, createProduct } from '@/api/product'
import { userSearch } from '@/api/remoteSearch'
import { CommentDropdown, PlatformDropdown, SourceUrlDropdown } from './Dropdown'

const defaultForm = {
  name: '',
  number: '',
  on_sale: false,
  price: 0,
  price_on_sale: 0
}

const categories = [
  'entry',
  'plat',
  'dessert',
  'drink',
]

export default {
  name: 'ProductDetail',
  components: { Tinymce, MDinput, Upload, Sticky, CommentDropdown, PlatformDropdown, SourceUrlDropdown },
  props: {
    isEdit: {
      type: Boolean,
      default: false
    }
  },
  data() {
    const validateRequire = (rule, value, callback) => {
      if (value === '') {
        this.$message({
          message: rule.field + '为必传项',
          typlengthe: 'error'
        })
        callback(new Error(rule.field + '为必传项'))
      } else {
        callback()
      }
    }
    const validateSourceUri = (rule, value, callback) => {
      if (value) {
        if (validateURL(value)) {
          callback()
        } else {
          this.$message({
            message: '外链url填写不正确',
            type: 'error'
          })
          callback(new Error('外链url填写不正确'))
        }
      } else {
        callback()
      }
    }
    return {
      postForm: Object.assign({}, defaultForm),
      loading: false,
      userListOptions: [],
      categories: categories,
      rules: {
        image_uri: [{ validator: validateRequire }],
        title: [{ validator: validateRequire }],
        content: [{ validator: validateRequire }],
        source_uri: [{ validator: validateSourceUri, trigger: 'blur' }]
      },
      tempRoute: {}
    }
  },
  computed: {
    lang() {
      return this.$store.getters.language
    }
  },
  created() {
    if (this.isEdit) {
      const id = this.$route.params && this.$route.params.id
      this.fetchData(id)
    } else {
      this.postForm = Object.assign({}, defaultForm)
    }

    // Why need to make a copy of this.$route here?
    // Because if you enter this product and quickly switch tag, may be in the execution of the setTagsViewTitle function, this.$route is no longer pointing to the current product
    // https://github.com/PanJiaChen/vue-element-admin/issues/1221
    this.tempRoute = Object.assign({}, this.$route)
  },
  methods: {
    fetchData(id) {
      fetchProduct(id).then(response => {
        this.postForm = response.data
      }).catch(err => {
        console.log(err)
      })
    },
    submitForm() {
      this.$refs.postForm.validate(valid => {
        if (valid) {
          this.postForm.price = parseInt(this.postForm.price)
          this.postForm.price_on_sale = parseInt(this.postForm.price_on_sale)
          if (this.$route.params.id) {
            updateProduct(this.$route.params.id, this.postForm).then(
              (res) => {
                this.loading = true
                let massageTitle = '成功'
                let message = '发布文章成功'
                let massageType = 'success'
                if (res.data.code === -1) {
                  massageType = 'error'
                  massageTitle = 'failed'
                  message = 'save failed'
                }
                this.$notify({
                  title: massageTitle,
                  message: message,
                  type: massageType,
                  duration: 2000
                })
                this.loading = false
              },
              () => {
                this.$notify({
                  title: 'failed',
                  message: 'save failed',
                  type: 'error',
                  duration: 2000
                })
              }
            );
          } else {
            createProduct(this.postForm).then(
              (res) => {
                this.loading = true
                let massageTitle = '成功'
                let message = '发布文章成功'
                let massageType = 'success'
                if (res.data.code === -1) {
                  massageType = 'error'
                  massageTitle = 'failed'
                  message = 'save failed'
                }
                this.$notify({
                  title: massageTitle,
                  message: message,
                  type: massageType,
                  duration: 2000
                })
                this.loading = false
                this.$router.push('/product/list')
              },
              () => {
                this.$notify({
                  title: 'failed',
                  message: 'save failed',
                  type: 'error',
                  duration: 2000
                })
              }
            );
          }


        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    draftForm() {
      if (this.postForm.content.length === 0 || this.postForm.title.length === 0) {
        this.$message({
          message: '请填写必要的标题和内容',
          type: 'warning'
        })
        return
      }
      this.$message({
        message: '保存成功',
        type: 'success',
        showClose: true,
        duration: 1000
      })
      this.postForm.status = 'draft'
    },
    getRemoteUserList(query) {
      userSearch(query).then(response => {
        if (!response.data.items) return
        this.userListOptions = response.data.items.map(v => v.name)
      })
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
@import "~@/styles/mixin.scss";
.createPost-container {
  position: relative;
  .createPost-main-container {
    padding: 40px 45px 20px 50px;
    .postInfo-container {
      position: relative;
      @include clearfix;
      margin-bottom: 10px;
      .postInfo-container-item {
        float: left;
      }
    }
    .editor-container {
      min-height: 500px;
      margin: 0 0 30px;
      .editor-upload-btn-container {
        text-align: right;
        margin-right: 10px;
        .editor-upload-btn {
          display: inline-block;
        }
      }
    }
  }
  .word-counter {
    width: 40px;
    position: absolute;
    right: -10px;
    top: 0px;
  }
}
</style>
