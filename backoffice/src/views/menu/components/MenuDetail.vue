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
                <el-form-item :label="$t('menu.name')" label-width="80px" >
                  <el-input v-model="postForm.name" type="text"  autosize placeholder="product name"/>
                </el-form-item>
              </el-col>
            </el-row>
            <div class="postInfo-container">
              <el-row>
                <el-col :span="24">
                  <el-form-item :label="$t('menu.name')" label-width="80px" >
                    <el-input-number v-model="postForm.price" type="number" placeholder="price">
                      <template slot="append"> euros </template>
                    </el-input-number>
                  </el-form-item>
                </el-col>
              </el-row>

              <el-row>
                <el-form-item :label="$t('menu.on_sale')" style="margin-bottom: 10px;" label-width="85px">
                  <div class="box-item">
                    <el-switch v-model="postForm.on_sale"/>
                  </div>
                </el-form-item>
              </el-row>

              <el-row>
                <el-col :span="24">
                  <el-form-item :label="$t('menu.price_on_sale')" label-width="80px" >
                    <el-input-number v-model="postForm.price_on_sale" type="number" placeholder="on sale price">
                      <template slot="append"> euros </template>
                    </el-input-number>
                  </el-form-item>
                </el-col>
              </el-row>

              <el-row>
                <el-form-item :label="$t('menu.entry')" label-width="80px" class="postInfo-container-item">
                <el-button type="primary" round @click="addEntry()">{{$t('menu.add_entry')}}</el-button>
                <el-row v-for="entry, index in postForm.entries" :key="entry._id">
                  <span>{{entry.product.name}}</span>
                  <el-col :span="24">
                    <el-form-item :span="12" :label="$t('menu.product_code')" label-width="80px" class="postInfo-container-item" >
                      <el-select v-model="entry.product" :remote-method="getRemoteEntryList" :placeholder="$t('search')" filterable remote>
                        <el-option v-for="(item) in entries" :key="item._id" :label="item.name" :value="item"/>
                      </el-select>
                    </el-form-item>

                    <el-form-item :label="$t('menu.supp_price')" :span="12"  label-width="80px" class="postInfo-container-item">
                      <el-input-number v-model="entry.extra_price" :placeholder="$t('menu.supp_price')">
                        <template slot="append"> euros </template>
                      </el-input-number>
                    </el-form-item>
                    <el-button type="danger" icon="el-icon-delete" circle @click="deleteEntry(index)"></el-button>
                  </el-col>
                </el-row>
                </el-form-item>
              </el-row>

              <el-row>
              <el-form-item :label="$t('menu.plat')" label-width="80px" class="postInfo-container-item">
                <el-button type="primary" round @click="addPlat()">{{$t('menu.add_plat')}}</el-button>
              <el-row v-for="plat, index in postForm.plats" :key="plat._id">
                <span>{{plat.product.name}}</span>
                <el-col :span="24">
                  <el-form-item :label="$t('menu.product_code')" label-width="80px"class="postInfo-container-item">
                    <el-select v-model="plat.product" :remote-method="getRemotePlatList" :placeholder="$t('search')" autocomplete filterable remote >
                      <el-option v-for="(item) in plats" :key="item._id" :label="item.name" :value="item"/>
                    </el-select>
                  </el-form-item>

                  <el-form-item :label="$t('menu.supp_price')" label-width="80px" class="postInfo-container-item">
                    <el-input-number :placeholder="$t('menu.supp_price')" v-model="plat.extra_price">
                      <template slot="append"> euros </template>
                    </el-input-number>
                  </el-form-item>
                  <el-button type="danger" icon="el-icon-delete" circle @click="deletePlat(index)"></el-button>

                </el-col>
              </el-row>
              </el-form-item>
              </el-row>

              <el-row>
                <el-form-item :label="$t('menu.dessert')" label-width="80px" class="postInfo-container-item">
               <el-button type="primary" round @click="addDessert()">{{$t('menu.add_dessert')}}</el-button>
                <el-row v-for="dessert, index in postForm.desserts" :key="dessert._id">
                  <span>{{dessert.product.name}}</span>
                  <el-col :span="24">
                    <el-form-item label="code:" label-width="80px" class="postInfo-container-item">
                      <el-select v-model="dessert.product" :remote-method="getRemoteDessertList" :placeholder="$t('search')" filterable remote>
                        <el-option v-for="item in desserts" :key="item._id" :label="item.name" :value="item"/>
                      </el-select>
                      <!--<el-input v-model="dessert.name" :value="item.name" type="hidden" ></el-input>-->
                    </el-form-item>

                    <el-form-item :label="$t('menu.supp_price')" label-width="80px" class="postInfo-container-item">
                      <el-input-number :placeholder="$t('menu.supp_price')" v-model="dessert.extra_price" >
                        <template slot="append"> euros </template>
                      </el-input-number>
                      <el-button type="danger" icon="el-icon-delete" circle @click="deleteDessert(index)"></el-button>

                    </el-form-item>
                  </el-col>
                </el-row>
                </el-form-item>
              </el-row>

              <el-row>
                <el-form-item :label="$t('menu.drink')" label-width="80px" class="postInfo-container-item">
                  <el-button type="primary" round @click="addDrink()">{{$t('menu.drink')}}</el-button>

                  <el-row  v-for="drink, index in postForm.drinks" :key="drink._id">
                    <el-col :span="24">
                      <span>{{drink.product.name}}</span>
                      <el-form-item label="code:" label-width="80px" class="postInfo-container-item">
                        <el-select v-model="drink.product" :remote-method="getRemoteDrinkList" :placeholder="$t('search')" filterable remote>
                          <el-option v-for="(item) in drinks" :key="item._id" :label="item.name" :value="item"/>
                        </el-select>
                      </el-form-item>

                      <el-form-item :label="$t('menu.supp_price')" label-width="80px" class="postInfo-container-item">
                        <el-input-number :placeholder="$t('menu.supp_price')" v-model="drink.extra_price" >
                          <template slot="append"> euros </template>
                        </el-input-number>
                        <el-button type="danger" icon="el-icon-delete" circle @click="deleteDrink(index)"></el-button>
                      </el-form-item>
                    </el-col>
                  </el-row>
                </el-form-item>
              </el-row>

            </div>
          </el-col>
        </el-row>

      </div>
    </el-form>

  </div>
</template>

<script>
import _ from 'lodash'
import Tinymce from '@/components/Tinymce'
import Upload from '@/components/Upload/singleImage3'
import MDinput from '@/components/MDinput'
import Sticky from '@/components/Sticky' // 粘性header组件
import { validateURL } from '@/utils/validate'
import { getProducts } from '@/api/product'
import { fetchMenu, updateMenu, createMenu } from '@/api/menu'

const defaultPlat = {
  code: '',
  extra_price: 0
}

const defaultForm = {
  name: '',
  entries: [],
  plats: [],
  desserts: [],
  drinks: [],
  on_sale: false,
  price: 0,
  price_on_sale: 0
}

export default {
  name: 'ProductDetail',
  components: { Tinymce, MDinput, Upload, Sticky },
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
      entries: [],
      plats: [],
      desserts: [],
      drinks: [],
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
    addEntry() {
      this.postForm.entries = [...this.postForm.entries, {
        code: '',
        extra_price: 0
      }]
    },
    addPlat() {
      this.postForm.plats = [...this.postForm.plats, {
        code: '',
        extra_price: 0
      }]
    },
    addDessert() {
      this.postForm.desserts = [...this.postForm.desserts, {
        code: '',
        extra_price: 0
      }]
    },
    addDrink() {
      this.postForm.drinks = [...this.postForm.drinks, {
        code: '',
        extra_price: 0
      }]
    },
    deleteEntry(index) {
      delete this.postForm.entries[index]
      this.postForm.entries = _.filter(this.postForm.entries, (v) => (v !== undefined))
    },
    deletePlat(index) {
      delete this.postForm.plats[index]
      this.postForm.plats = _.filter(this.postForm.plats, (v) => (v !== undefined))
    },
    deleteDessert(index) {
      delete this.postForm.desserts[index]
      this.postForm.desserts = _.filter(this.postForm.desserts, (v) => (v !== undefined))
    },
    deleteDrink(index) {
      delete this.postForm.drinks[index]
      this.postForm.drinks = _.filter(this.postForm.drinks, (v) => (v !== undefined))
    },
    fetchData(id) {
      fetchMenu(id).then(response => {
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
          this.postForm.entries.map(
            (v) => {
              v.extra_price = parseInt(v.extra_price)
              return v
            }
          )
          this.postForm.plats.map(
            (v) => {
              v.extra_price = parseInt(v.extra_price)
              return v
            }
          )
          this.postForm.desserts.map(
            (v) => {
              v.extra_price = parseInt(v.extra_price)
              return v
            }
          )
          this.postForm.drinks.map(
            (v) => {
              v.extra_price = parseInt(v.extra_price)
              return v
            }
          )
          if (this.$route.params.id) {
            updateMenu(this.$route.params.id, this.postForm).then(
              (res) => {
                this.loading = true
                let massageTitle = '成功'
                let message = '发布文章成功'
                let massageType = 'success'
                if (res.data.code !== 0) {
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
            createMenu(this.postForm).then(
              (res) => {
                this.loading = true
                let massageTitle = '成功'
                let message = '发布文章成功'
                let massageType = 'success'
                if (res.data.code !== 0) {
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
                this.$router.push('/menu/list')
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
    getRemoteEntryList(query) {
      getProducts({category: 'entry', query: query}).then(response => {
        if (!response.data.items) return
        this.entries = response.data.items
      })
    },
    getRemotePlatList(query) {
      getProducts({category: 'plat', query: query}).then(response => {
        if (!response.data.items) return
        this.plats = response.data.items
      })
    },
    getRemoteDessertList(query) {
      getProducts({'category': 'dessert', query: query}).then(response => {
        if (!response.data.items) return
        this.desserts = response.data.items
      })
    },
    getRemoteDrinkList(query) {
      getProducts({'category': 'drink', query: query}).then(response => {
        if (!response.data.items) return
        this.drinks = response.data.items
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
