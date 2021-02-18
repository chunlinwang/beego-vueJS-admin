<template>
  <div class="app-container">
    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" :label="$t('order.id')" >
        <template slot-scope="scope">
          <span>{{ scope.row._id }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" :label="$t('order.orderNumber')">
        <template slot-scope="scope">
          <span>{{ scope.row.order_number }}</span>
        </template>
      </el-table-column>

      <!--<el-table-column align="center" :label="$t('order.order_items')">-->
        <!--<template slot-scope="scope">-->
          <!--<span>{{ scope.row.order_items }}</span>-->
        <!--</template>-->
      <!--</el-table-column>-->

      <el-table-column align="center" :label="$t('order.shippingMode')">
        <template slot-scope="scope">
          <span>{{ scope.row.shipping_mode }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" :label="$t('order.totalPromo')">
        <template slot-scope="scope">
          <span>{{ scope.row.total_promo }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" :label="$t('order.total')">
        <template slot-scope="scope">
          <span>{{ scope.row.total }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" :label="$t('order.createAt')">
        <template slot-scope="scope">
          <span>{{ scope.row.create_at | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('order.actions')" align="center" >
        <template slot-scope="scope">
          <el-row style="margin-bottom: 5px">
            <el-button type="primary" size="small" icon="el-icon-success" round>{{$t('order.normal')}}</el-button>
          </el-row>

          <el-row style="margin-bottom: 5px">
            <el-button type="warning" size="small" icon="el-icon-warning" round>{{$t('order.busy')}}</el-button><br/>
          </el-row>

          <el-row style="margin-bottom: 5px">
            <el-button type="danger" size="small" icon="el-icon-warning" round>{{$t('order.extraBusy')}}</el-button>
          </el-row>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />

  </div>
</template>

<script>
import { fetchList } from '@/api/order'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

export default {
  name: 'OrderList',
  components: { Pagination },
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'info',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 20
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        console.log(response.data.items)
        this.list = response.data.items
        this.total = response.data.total
        this.listLoading = false
      })
    },
    handleSizeChange(val) {
      this.listQuery.limit = val
      this.getList()
    },
    handleCurrentChange(val) {
      this.listQuery.page = val
      this.getList()
    }
  }
}
</script>

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
