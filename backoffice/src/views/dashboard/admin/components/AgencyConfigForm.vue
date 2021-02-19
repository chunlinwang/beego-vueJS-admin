<template>
  <div class="createPost-container background-white">
    <el-form ref="agencyConfigForm" :model="agencyConfigForm" class="form-container">
      <sticky :class-name="'sub-navbar '+agencyConfigForm.status">
        <el-tag type="success">STETTING</el-tag>
        <el-button v-loading="loading" style="margin-left: 10px;" type="success" @click="submitForm">Save</el-button>
      </sticky>

      <div class="createPost-main-container">
        <!--<el-row>-->
        <!--<Warning/>-->
        <!--</el-row>-->
        <el-row>
          <el-form-item :label="$t('Address')" style="margin-bottom: 10px;" label-width="85px">
            <el-input
              :rows="1"
              v-model="agencyConfigForm.address"
              type="text"
              class="article-textarea"
              autosize
              placeholder="Address"/>
          </el-form-item>
        </el-row>

        <el-row>
          <el-col :span="12">
            <el-form-item :label="$t('City')" style="margin-bottom: 10px;" label-width="85px">
              <el-input
                :rows="1"
                v-model="agencyConfigForm.city"
                type="text"
                class="article-textarea"
                autosize
                placeholder="City"/>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item :label="$t('ZipCode')" style="margin-bottom: 10px;" label-width="85px">
              <el-input
                :rows="1"
                v-model="agencyConfigForm.zip_code"
                type="text"
                class="article-textarea"
                autosize
                placeholder="ZipCode"/>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="12">
            <el-form-item :label="$t('Phone')" style="margin-bottom: 10px;" label-width="85px">
              <el-input
                :rows="1"
                v-model="agencyConfigForm.phone"
                type="text"
                class="article-textarea"
                autosize
                placeholder="Phone"/>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item :label="$t('Email')" style="margin-bottom: 10px;" label-width="85px">
              <el-input
                :rows="1"
                v-model="agencyConfigForm.email"
                type="text"
                class="article-textarea"
                autosize
                placeholder="Email"/>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row v-for="(schedule, index) in agencyConfigForm.schedules">
          <el-col v-if="(schedule.day) % 2 != 0" :span="24">
            <span class="el-form-item__label">{{ $t(Weeks[(schedule.day-1)/2]) }}: </span>
          </el-col>
          <el-col :span="4">
            <span v-if="(schedule.day) % 2 == 0" class="el-form-item__label">{{ $t("openTime.PM") }}</span>
            <span v-else class="el-form-item__label">{{ $t("openTime.AM") }}</span>
          </el-col>

          <el-time-picker
            v-model="agencyConfigForm.schedules[index].begin_hour"
            :picker-options="{
              start: '10:00',
              step: '00:15',
              end: '23:30'
            }"
            format="HH:mm"
            type="time"
            placeholder="select time"/>
          -
          <el-time-picker
            v-model="agencyConfigForm.schedules[index].end_hour"
            :picker-options="{
              start: '10:00',
              step: '00:15',
              end: '23:30'
            }"
            type="time"
            format="HH:mm"
            placeholder="select time"/>
        </el-row>
        <el-row>
          <el-form-item :label="$t('Open')" style="margin-bottom: 10px;" label-width="85px">
            <div class="box-item">
              <el-switch v-model="agencyConfigForm.is_open"/>
            </div>
          </el-form-item>
        </el-row>
      </div>
    </el-form>

  </div>
</template>

<script>
import Tinymce from '@/components/Tinymce'
import Upload from '@/components/Upload/singleImage3'
import MDinput from '@/components/MDinput'
import Sticky from '@/components/Sticky'
import { save } from '@/api/agency'

const Weeks = [
  'Monday',
  'Tuesday',
  'Wednesday',
  'Thursday',
  'Friday',
  'Saturday',
  'Sunday'
]

const defaultForm = {
  address: '108 rue du vieux pont de Sévres',
  city: 'Boulogne-Billancourt',
  email: '',
  phone: '01 58 17 04 04',
  zip_code: '92100',
  is_open: true,
  schedules: [
    {
      day: 1
    },
    {
      day: 1
    },
    {
      day: 2
    },
    {
      day: 2
    },
    {
      day: 3
    },
    {
      day: 3
    },
    {
      day: 4
    },
    {
      day: 4
    },
    {
      day: 5
    },
    {
      day: 5
    },
    {
      day: 6,
      begin_hour: null,
      end_hour: null
    },
    {
      day: 6
    },
    {
      day: 7,
      begin_hour: null,
      end_hour: null
    },
    {
      day: 7
    }
  ]
}

export default {
  name: 'AgencyConfigForm',
  components: { Tinymce, MDinput, Upload, Sticky },
  props: {
    isEdit: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      agencyConfigForm: Object.assign({}, defaultForm),
      Weeks: Object.assign({}, Weeks),
      loading: false,
      userListOptions: [],
      tempRoute: {}
    }
  },
  computed: {
    lang() {
      return this.$store.getters.language
    }
  },
  created() {
    this.agencyConfigForm = Object.assign({}, defaultForm)
    this.$store.dispatch('getAgency').then(() => {
      this.agencyConfigForm = Object.assign(this.agencyConfigForm, this.$store.getters.agency)
    }).catch((e) => {
      console.log('can not fetch agnecy info.')
    })
  },
  methods: {
    fetchData() {
      this.$store.dispatch('getAgency').then(() => {
      }).catch(() => {
        console.log('can not fetch agnecy info.')
      })
    },
    submitForm() {
      // should change
      this.$refs.agencyConfigForm.validate(valid => {
        if (valid) {
          this.loading = true
          save(this.agencyConfigForm).then(() => {
            this.fetchData()
            this.$notify({
              title: 'Success',
              message: 'Saved',
              type: 'success',
              duration: 2000
            })
          })
          this.loading = false
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  @import "~@/styles/mixin.scss";

  .background-white {
    background: #fff;
  }

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
